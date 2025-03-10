// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl9.0.so"
	tkBin  = "libtcl9tk9.0.so"
)

//go:embed embed/linux/amd64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtcl9.0.so":    "f23be853acc3cd7d11383dc8a0c0916deaf97fd69424f6115e2d4b7b2a84613a",
	"libtcl9tk9.0.so": "c5da7f669dc5a1f9c885e5e30e4d207a81cbcb2428d10938bd37ecc359398cd9",
	"libtk9.0.1.zip":  "52dedac64f3d4abce8fd578e41bf0cdfc09ac807cd8a30137ea4842efbbed700",
}
