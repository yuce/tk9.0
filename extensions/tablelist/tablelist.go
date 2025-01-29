// Copyright 2025 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tablelist provides the [tklib tablelist].
//
// To make the extension available in an application:
//
//	import "modernc.org/tk9.0/extension/tablelist"
//
// See the [modernc.org/tk9.0.Extension] documentation for information about
// initalizing extensions at runtime.
//
// [tklib tablelist]: https://github.com/tcltk/tklib/tree/master/modules/tablelist
package tablelist // import "modernc.org/tk9.0/extensions/tablelist"

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	. "modernc.org/tk9.0"
)

const (
	// Update when tablelist.zip file changed.
	version = "v0.1.0"
)

var (
	_ Extension = (*extension)(nil)

	//go:embed embed/tablelist.zip
	zip []byte

	ctx         ExtensionContext
	initialized bool

	// Version reports the version of the tablelist package. Valid after successful
	// initialization.
	Version string
)

func init() {
	RegisterExtension("tablelist", newExtension())
}

func setup(context ExtensionContext) (err error) {
	defer func() {
		initialized = true
	}()

	if initialized {
		return nil
	}

	ctx = context
	root, err := os.UserCacheDir()
	if err != nil {
		return err
	}

	const tablelist = "tablelist.zip"
	dir, err := mkzip(filepath.Join(root, "modernc.org", "tk9.0.0", "extensions", "tablelist", version), tablelist, zip)
	if err != nil {
		return err
	}

	var wd string
	if wd, err = os.Getwd(); err != nil {
		return err
	}

	defer func() {
		err = errors.Join(err, os.Chdir(wd))
	}()

	if err = os.Chdir(dir); err != nil {
		return
	}

	mount := "/extensions/tablelist"
	if Version, err = ctx.Eval(fmt.Sprintf(`
lappend auto_path [zipfs mount %s %s]
package require tablelist_tile
`, tablelist, mount)); err != nil {
		return err
	}

	return nil
}

func mkzip(dir, base string, zip []byte) (r string, err error) {
	if _, err = os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			return "", err
		}

		if err = os.MkdirAll(dir, 0o770); err != nil {
			return "", err
		}
	}

	path := filepath.Join(dir, base)
	if _, err = os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			return "", err
		}

		return dir, os.WriteFile(path, zip, 0o660)
	}

	b, err := os.ReadFile(path)
	if err == nil {
		if bytes.Equal(b, zip) {
			return dir, nil
		}
	}

	os.Remove(path)
	if err := os.WriteFile(path, zip, 0o660); err == nil {
		return dir, nil
	}

	dir, err = os.MkdirTemp("", "tablelist-extension")
	if err != nil {
		return "", err
	}

	path = filepath.Join(dir, base)
	return dir, os.WriteFile(path, zip, 0o660)
}

type extension struct{}

func newExtension() *extension {
	return &extension{}
}

func (e *extension) Initialize(context ExtensionContext) error {
	return setup(context)
}

// TablelistWidget represents a tklib tablelist.
//
// The extensive documentation is available [here].
//
// [here]: https://www.nemethi.de/tablelist/index.html
type TablelistWidget struct {
	*Window
}

// Tablelist0 returns a newly created tablelist using raw Tcl string 'args'.
// Example:
//
//	Tablelist0(`.t -columns {0 "First Column" 0 "Another column"}`)
//
// This is a bootstrap function enabling to use the megawidget without first
// implementing the overwhelming amount of its options and methods in Go.
//
// Use with caution.
func Tablelist0(args string) (r *TablelistWidget) {
	return &TablelistWidget{Window: ctx.RegisterWindow(ctx.EvalErr(fmt.Sprintf("tablelist::tablelist %s", args)))}
}

// Do executes the raw Tcl string 'args'. Example:
//
//	t.Do(`insert end [list "first row" "another value"]`)
//
// This is a bootstrap function enabling to use the megawidget without first
// implementing the overwhelming amount of its options and methods in Go.
//
// Use with caution.
func (t *TablelistWidget) Do(args string) string {
	return ctx.EvalErr(fmt.Sprintf("%s %s", t, args))
}
