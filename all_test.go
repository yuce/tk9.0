// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"slices"
	"testing"

	_ "github.com/adrg/xdg"       // generator.go
	_ "github.com/expr-lang/expr" // examples
	_ "golang.org/x/net/html"     // generator.go
	_ "modernc.org/egg/lib"       // ltok_generator.go
	_ "modernc.org/ngrab/lib"     // generator.go
	_ "modernc.org/rec/lib"       // generator.go
)

var (
	display = os.Getenv("DISPLAY")
	target  = fmt.Sprintf("%s/%s", goos, goarch)
)

func TestMain(m *testing.M) {
	if Error != nil {
		fmt.Fprintln(os.Stderr, Error)
		os.Exit(1)
	}

	flag.Parse()
	rc := m.Run()
	Finalize()
	os.Exit(rc)
}

// Commit 18c4e94e171d4 diff
//
// -               `([^$]|\\\$)*`,           // Not TeX, incl. "\$"
// -               `\$([^$]|\\\$)*\$`,       // $TeX$ or $$TeX$$, incl. $Te\$X$
// -               `\$\$?([^$]|\\\$)*\$\$?`, // $TeX$ or $$TeX$$, incl. $Te\$X$
// +               `([^$]|\\\$)*`,         // Not TeX, incl. "\$"
// +               `\$([^$]|\\\$)*\$`,     // $TeX$, incl. $Te\$X$
// +               `\$\$([^$]|\\\$)*\$\$`, // $$TeX$$, incl. $Te\$X$
func TestTokenizer(t *testing.T) {
	for i, test := range []struct {
		s    string
		ids  []int
		toks []string
	}{
		{},
		{"a", []int{0}, []string{"a"}},
		{"\\$", []int{0}, []string{"\\$"}},
		{"\\$\\$", []int{0}, []string{"\\$\\$"}},
		{"\\$\\$\\$", []int{0}, []string{"\\$\\$\\$"}},

		{"\\$\\$\\$\\$", []int{0}, []string{"\\$\\$\\$\\$"}},
		{"a\\$", []int{0}, []string{"a\\$"}},
		{"a\\$\\$", []int{0}, []string{"a\\$\\$"}},
		{"a\\$\\$\\$", []int{0}, []string{"a\\$\\$\\$"}},
		{"a\\$\\$\\$\\$", []int{0}, []string{"a\\$\\$\\$\\$"}},

		{"$a$", []int{1}, []string{"$a$"}},
		// Not valid since 18c4e94e171d4 {"$$a$", []int{2}, []string{"$$a$"}},
		{"$$a$$", []int{2}, []string{"$$a$$"}},
		// Not valid since 18c4e94e171d4 {"$a$$", []int{2}, []string{"$a$$"}},
		{"x$a$", []int{0, 1}, []string{"x", "$a$"}},

		// Not valid since 18c4e94e171d4 {"x$$a$", []int{0, 2}, []string{"x", "$$a$"}},
		{"x$$a$$", []int{0, 2}, []string{"x", "$$a$$"}},
		// Not valid since 18c4e94e171d4 {"x$a$$", []int{0, 2}, []string{"x", "$a$$"}},
		{"x$a$y", []int{0, 1, 0}, []string{"x", "$a$", "y"}},
		// Not valid since 18c4e94e171d4 {"x$$a$y", []int{0, 2, 0}, []string{"x", "$$a$", "y"}},

		{"x$$a$$y", []int{0, 2, 0}, []string{"x", "$$a$$", "y"}},
		// Not valid since 18c4e94e171d4 {"x$a$$y", []int{0, 2, 0}, []string{"x", "$a$$", "y"}},
		// Not valid since 18c4e94e171d4 {"x\\$0$a\\$1b$$\\$y", []int{0, 2, 0}, []string{"x\\$0", "$a\\$1b$$", "\\$y"}},
	} {
		ids, toks := tokenize(test.s)
		if g, e := fmt.Sprintf("%v %q", ids, toks), fmt.Sprintf("%v %q", test.ids, test.toks); g != e {
			t.Errorf("#%3v: `%s`\ngot %s\nexp %s", i, test.s, g, e)
		}
	}
}

// Credits: https://gitlab.com/cznic/tk9.0/-/issues/51#note_2374472931
func TestParseList(t *testing.T) {
	switch target {
	case 
		"linux/386",
		"linux/arm",
		"linux/loong64",
		"linux/ppc64le":

		if display == "" {
			t.Skipf("This test is known to work only interactively on %s, use ssh -X.", target)
		}
	case "linux/s390x":
		t.Skipf("This test is known to not work on %s.", target)
	}
	Initialize()
	defer Finalize()

	tests := []struct {
		name     string
		in       string
		expected []string
	}{
		{"empty", "", []string{}},
		{"one item", "abc", []string{"abc"}},
		{"multiple items", "abc def ghi", []string{"abc", "def", "ghi"}},
		{"multiple inter item spaces", "abc   def", []string{"abc", "def"}},
		{"leading and trailing spaces", "  abc def  ", []string{"abc", "def"}},
		{"delimited item at start", "{ab c} def ghi", []string{"ab c", "def", "ghi"}},
		{"delimited item in middle", "abc {de f} ghi", []string{"abc", "de f", "ghi"}},
		{"delimited item at end", "abc def {gh i}", []string{"abc", "def", "gh i"}},
		{"all items delimited", "{abc} {def} {ghi}", []string{"abc", "def", "ghi"}},
		{"delimited with leading and trailing space", " {ab c} def {gh i}  ", []string{"ab c", "def", "gh i"}},
		{"whitespace in items", "{ab c} {de\tf} {gh\ni}", []string{"ab c", "de\tf", "gh\ni"}},
		{"braces in items", `ab\{c de\}f`, []string{`ab{c`, "de}f"}},
		{"backslash not escaping a brace", `{ab\c}`, []string{"ab\\c"}},
		{"whitespace in items", "{ab c} {de\tf} {gh\ni}", []string{"ab c", "de\tf", "gh\ni"}},
		{"braces in elements", "a{b c}d e{f} g{{h i{}j k}{l }m}}", []string{"a{b", "c}d", "e{f}", "g{{h", "i{}j", "k}{l", "}m}}"}},
		{"nested list", "{abc {def ghi}} jkl", []string{"abc {def ghi}", "jkl"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := parseList(test.in)
			if slices.Compare(list, test.expected) != 0 {
				t.Errorf("got %#v, expected %#v", list, test.expected)
			}
		})
	}
}
