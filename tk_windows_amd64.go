// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "tcl90.dll"
	tkBin  = "tcl9tk90.dll"
)

//go:embed embed/windows/amd64/lib.zip
var libZip []byte

var shasig = map[string]string{
	"libtk9.0.1.zip": "942297fe73c74085c431e315cfb8e4ef0f4a8c2ef7ab11840268da91c8b31494",
	"libtommath.dll": "2d760fefb452665b6af8c8d9d29f3a8378f10fc0847cdd9938ea0cb5edf1d573",
	"tcl90.dll": "704e550abfaa9c9659bf6f16d2ca76783fb68cbafa49a83613a148641cf77605",
	"tcl9dde14.dll": "d696384cce2adf5d80e6f969b1dc262ec3fb58a81d7b614bfb39cbba2e424a21",
	"tcl9registry13.dll": "34bfa90eced8921108df3d58c1d7f878c68e416283873446e27262c6c7b01a1b",
	"tcl9tk90.dll": "ed2f91f231f840f51e2db5f666a3366455f022503a7eeb32a11413acafcec8b2",
	"tcldde14.dll": "b072fd5f98f47d6fa9d1c8e8992ffb8c7f364dc92c4d413bf47fc38aa5a6e753",
	"tclregistry13.dll": "0129eb4d73dd44e4dfe8db208e165246b986d7eeaf4f8bc5a18b3110175d770e",
	"zlib1.dll": "04117778e255ed158cf6a4a1e51aa40f49124d9035208218fbfebbe565cf254d",
}
