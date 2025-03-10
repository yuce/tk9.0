// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl9.0.so"
	tkBin  = "libtcl9tk9.0.so"
)

//go:embed embed/linux/arm64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtcl9.0.so":    "da5f4b8a5d540e653cdf64c474b69ce6e4b4105f8191d8e73080cd8a2bf38866",
	"libtcl9tk9.0.so": "f8e827d135283dff08c24ea7d579791c62c93a864d6b89a3dc4fda81ce873646",
	"libtk9.0.1.zip":  "9704dbb3220e947ac020367521415d79f1609506800d933d290b2397df113518",
}
