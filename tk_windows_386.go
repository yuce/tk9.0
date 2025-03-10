// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "tcl90.dll"
	tkBin  = "tcl9tk90.dll"
)

//go:embed embed/windows/386/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtk9.0.1.zip": "b70d986331519aa11a440dd7acbb758e6a670f8a41213bc0e8942c9eb44e0ec2",
	"libtommath.dll": "7ff97843cde97215fcf4f087d61044cda01286630b486398117967e577e039e3",
	"tcl90.dll": "9523c7a050c82adfa9d0c2e5c32ebbc803b1e449d94231e919bba3aeecb8280b",
	"tcl9dde14.dll": "842b8d64fcd03a6b712480b6a67a7b2ec2e36bed95c0de78e0ce538b2652553d",
	"tcl9registry13.dll": "27038f4b9f8aecc4f9e8295affec24a4b7ff3497c92adfc854bb25935d2aa41f",
	"tcl9tk90.dll": "40d30eb25c585e4058082e7d004468ab37ee41f5d84f4e05b6d951d520215866",
	"tcldde14.dll": "57184fd17b5824b38835ac5e89885d500503af145552e0912c11c0d9041e0382",
	"tclregistry13.dll": "8adb1dc356218393acb4f1c3e7e7dc3900971f2c2d4081df99a766c12fc06e36",
	"zlib1.dll": "60f637680d84a0717cbee4cbf219b6215ca1f21fe0b32c8de2819c328c72ef15",
}
