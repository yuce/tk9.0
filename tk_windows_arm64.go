// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "tcl90.dll"
	tkBin  = "tcl9tk90.dll"
)

//go:embed embed/windows/arm64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtk9.0.1.zip": "bf5f8fc4c620882f8683ac83435c45a6d1da94a78fc8b27dc2e98b59d4cb4f0b",
	"libtommath.dll": "907b9c7860fc07231f1e238551715e5d813283807f52dc383dae0cb47a879d29",
	"tcl90.dll": "1c6a2ae3e8bc59b7790aa753a9b66c8b38671583050865888a10e468f644822f",
	"tcl9dde14.dll": "e4bf19739577dcca27a5ac053d60b301d2c5fbd1862b6f2f588429f04a297f98",
	"tcl9registry13.dll": "c693a596e2be187226653ec78ae61a547d9e274aa3765e03bd2540351662aa41",
	"tcl9tk90.dll": "76b01f67cf1f8582b20196f46bfa0253b1b07519ff172706082468313d23182d",
	"tcldde14.dll": "a24a942954116158d13a1bc132f4d82c743281600c74c1365f4d669746d723b1",
	"tclregistry13.dll": "12a4f38fd6efd61962a3d10e3399a0668e4649accc4472b2d98f8e0e4690d69e",
	"zlib1.dll": "6f10a76dcc2c831d1f08d98c0b345afa0911bec0238fcba357b612ccc6ab5d81",
}
