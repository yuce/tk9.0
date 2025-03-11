// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl9.0.so"
	tkBin  = "libtcl9tk9.0.so"
)

//go:embed embed/freebsd/arm64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtcl9.0.so":    "29db3ce5108115079af8136d08fbdbddd71c97ebb4ca9fc17d46eb3cef18a538",
	"libtcl9tk9.0.so": "25f8fc725b35163e6235c3d5cb8bdf687dda0f5e0760e0a89a813e48dd10da29",
	"libtk9.0.1.zip":  "984a4e2fc2632aea902b8a6ab2d7a40fe8dbc365aee9089671fa2d8aa0551b32",
}
