// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl9.0.so"
	tkBin  = "libtcl9tk9.0.so"
)

//go:embed embed/freebsd/amd64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtcl9.0.so":    "2ed139db82c91e57346e58eab14b47c7975a50734365fac47ac6e2785d342e26",
	"libtcl9tk9.0.so": "768b7ee6597403c6fb2ea5b38ba261dd387a2c1c54b9f75b50acacf3bc25f9cc",
	"libtk9.0.1.zip":  "1a13f7fbf49f7ba597b76592084c86563d097ccaa5930ac185c2c02c40ec81c3",
}
