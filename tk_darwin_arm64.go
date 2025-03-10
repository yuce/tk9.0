// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl9.0.dylib"
	tkBin  = "libtcl9tk9.0.dylib"
)

//go:embed embed/darwin/arm64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtcl9.0.dylib":    "643ebd83a980e599552b4ddfa540daeefa0caf6a240ce8b1234ba8598e5df1d0",
	"libtcl9tk9.0.dylib": "2969f7993c5d35a2f139062d1ff578884b867c523046a76bbd333a10e73f2c87",
	"libtk9.0.1.zip":     "5e8b7b3bc7e37881fad6aacfa868e767baf5890be0bebe21094587bf81938c09",
}
