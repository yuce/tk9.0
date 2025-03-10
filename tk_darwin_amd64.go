// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl9.0.dylib"
	tkBin  = "libtcl9tk9.0.dylib"
)

//go:embed embed/darwin/amd64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtcl9.0.dylib":    "b278e78fb3fd034091be7697500ab17018062cd63f5a9b9ff6249d529bc1b49f",
	"libtcl9tk9.0.dylib": "1a9c23a20a2162742dbbcbcdd647689ef2910beb55127f8281c5cd4b3b6549b5",
	"libtk9.0.1.zip":     "a975213545f57f71ca73a17ca7ab000fd6b436f352e3a7e9d9c878c59e094119",
}
