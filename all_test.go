// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"syscall"
	"testing"
	"time"

	_ "github.com/adrg/xdg"       // generator.go
	_ "github.com/expr-lang/expr" // examples
	_ "golang.org/x/net/html"     // generator.go
	_ "modernc.org/egg/lib"       // ltok_generator.go
	_ "modernc.org/ngrab/lib"     // generator.go
	_ "modernc.org/rec/lib"       // generator.go
)

const (
	xvfbDisplayVar = "XVFB_DISPLAY"
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

	if display == "" && goos == "linux" {
		if s := os.Getenv(xvfbDisplayVar); s != "" {
			display = s
			os.Setenv("DISPLAY", s)
		}
	}
	flag.Parse()
	rc := m.Run()
	Finalize()
	os.Exit(rc)
}

func killXvfb() {
	out := mustSys("pgrep", "-l", "Xvfb")
	a := strings.Split(string(out), "\n")
	for _, v := range a {
		b := strings.Fields(v)
		if len(b) != 0 {
			mustSys("kill", b[0])
		}
	}
}

func mustStart(arg0 string, args ...string) {
	if err := exec.Command(arg0, args...).Start(); err != nil {
		panic(err)
	}
}

func mustSys(arg0 string, args ...string) (r []byte) {
	var err error
	if r, err = sys(arg0, args...); err != nil {
		panic(err)
	}

	return r
}

func sys(arg0 string, args ...string) (r []byte, err error) {
	return exec.Command(arg0, args...).CombinedOutput()
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

func TestExamples(t *testing.T) {
	if !isBuilder {
		t.Skip("not a builder")
	}

	const retries = 10
	switch goos {
	case "linux":
		if display == "" {
			t.Fatal("DISPLAY=")
		}
	}

	tmpDir := t.TempDir()
	m, err := filepath.Glob(filepath.Join("_examples", "*.go"))
	if err != nil {
		t.Fatal(err)
	}

	blacklist := map[string]struct{}{
		"demo.go":        {},
		"embed.go":       {},
		"events.go":      {},
		"fontmetrics.go": {},
		"ring.go":        {},
		"tex.go":         {},
	}

	graylist := map[string]struct{}{}

next:
	for _, v := range m {
		base := filepath.Base(v)
		if _, ok := blacklist[base]; ok {
			continue
		}

		t.Log(v)
		for i := 0; i < retries; i++ {
			t.Logf("\t%v: %v", i, v)
			if err = testExample(t, tmpDir, v, i); err == nil {
				continue next
			}
		}

		if _, ok := graylist[filepath.Base(v)]; !ok {
			t.Errorf("%v: %v", v, err)
		}
	}
}

func testExample(t *testing.T, tmpDir string, fn string, n int) (err error) {
	bin := filepath.Join(tmpDir, fmt.Sprintf("prog%v", n))
	if goos == "windows" {
		bin += ".exe"
	}

	if _, err := sys("go", "build", "-o", bin, fn); err != nil {
		t.Fatal(err)
	}

	token := fmt.Sprint(time.Now().UnixNano())
	os.Setenv(testHookWaitVar, token)

	var stdout, stderr bytes.Buffer
	cmd := exec.Command(bin)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	t0 := time.Now()
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}

	crashCheckDuration := 3 * time.Second
	crashCheckTimer := time.NewTimer(crashCheckDuration)

	waitChan := make(chan error, 1) // Buffered channel to prevent goroutine leak

	go func() {
		waitChan <- cmd.Wait()
	}()

	select {
	case err := <-waitChan:
		crashCheckTimer.Stop() // Stop timer if process exits before timeout.
		if err != nil {
			if _, ok := err.(*exec.ExitError); ok {
				return fmt.Errorf("process crashed within the check duration (after %s).\nstdout=%s\nstderr=%s", time.Since(t0), stdout.Bytes(), stderr.Bytes())
			} else {
				return fmt.Errorf("process exited with error: %v", err)
			}
		} else {
			return fmt.Errorf("process unexpectedly exited normally")
		}
	case <-crashCheckTimer.C:
		if runtime.GOOS == "windows" {
			if err := cmd.Process.Kill(); err != nil {
				t.Fatalf("%v: error killing process: %v", fn, err)
			}
		} else {
			if err := cmd.Process.Signal(syscall.SIGTERM); err != nil {
				if err := cmd.Process.Kill(); err != nil {
					t.Fatalf("%v: error killing process: %v", fn, err)
				}
			}
		}
		err := <-waitChan // Wait for the process to fully exit after the kill signal.
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				status := exitError.Sys().(syscall.WaitStatus)
				if !status.Signaled() {
					return fmt.Errorf("process exited with error: %v", err)
				}
			} else {
				return fmt.Errorf("process exited with error: %v", err)
			}
		}
	}

	if g, ge, e := strings.TrimSpace(string(stdout.Bytes())), strings.TrimSpace(string(stderr.Bytes())), token; g != e {
		return fmt.Errorf("expected process stdout=%s, got stdout=%s, stderr=%s", e, g, ge)
	}

	return nil
}
