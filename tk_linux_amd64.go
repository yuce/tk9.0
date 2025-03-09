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
	"libtcl9.0.so":    "7294529d7397998c5c1bb13b4fd8c75ea9ad96caffcace581a5056aa65f40ead",
	"libtcl9tk9.0.so": "681a7a5fdd7edd19f844df06c0287d6a5f689d93b6001c8e9143d40265366333",
	"libtk9.0.1.zip":  "542e9ab34ae1a4eed1ccfd349d0fcd5b2acf6c606ba4e61bd6097b00761ac06b",
}
