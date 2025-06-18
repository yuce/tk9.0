// Copyright 2024 The tk9.0-go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tk9_0 // import "modernc.org/tk9.0"

import _ "embed"

const (
	tclBin = "libtcl90.so.1.0"
	tkBin  = "libtcl9tk90.so.1.0"
)

var moreDLLs = []*dllInfo{
	{"libtcl9tkimg201.so", "Tkimg_Init"},
	{"libtcl9jpegtcl960.so", "Jpegtcl_Init"},
	{"libtcl9tkimgjpeg201.so", "Tkimgjpeg_Init"},
	{"libtcl9tkimgbmp201.so", "Tkimgbmp_Init"},
	{"libtcl9tkimgico201.so", "Tkimgico_Init"},
	{"libtcl9tkimgpcx201.so", "Tkimgpcx_Init"},
	{"libtcl9tkimgxpm201.so", "Tkimgxpm_Init"},
	{"libtcl9zlibtcl131.so", "Zlibtcl_Init"},
	{"libtcl9pngtcl1644so", "Pngtcl_Init"},
	{"libtcl9tkimgpng201.so", "Tkimgpng_Init"},
	{"libtcl9tkimgppm201.so", "Tkimgppm_Init"},
	{"libtcl9tkimgtga201.so", "Tkimgtga_Init"},
	{"libtcl9tifftcl470.so", "Tifftcl_Init"},
	{"libtcl9tkimgtiff201.so", "Tkimgtiff_Init"},
	{"libtcl9tkimgxbm201.so", "Tkimgxbm_Init"},
}

//go:embed embed/openbsd/amd64/lib.zip
var libZip []byte

// Keep last for internal/shasig.go to update.
var shasig = map[string]string{
	"libtcl90.so.1.0":          "b7990b901c8dd609edaab5e90980bd1a5ef41d8bb7088109f759f7377d721d20",
	"libtcl9jpegtcl960.so":     "e652c9ab4c265766fbcc01e0274cbc1f6ac269800ac5c91e4980e132b33654d3",
	"libtcl9pngtcl1644.so":     "3f649554f49f7ec9df7eedccbb10d4fdb62919775489d051d979e9e3ead65302",
	"libtcl9tifftcl470.so":     "b5b8d77d61a1bf068b14c9685190403b0752cff5e98a1ca61f71317deabb53fa",
	"libtcl9tk90.so.1.0":       "f398047e3cff5a650128aa16ed0b58b03d15c31dbde3040f93bdf0b7b056522d",
	"libtcl9tkimg201.so":       "a23d041e5ff55ffe05c0143aee641f61b0430490b21a1a3cfe199608bce9f481",
	"libtcl9tkimgbmp201.so":    "b887e1aa163ccb77a2ab7977cb7df59403e27095bbde4f655cf138fd9397a90a",
	"libtcl9tkimgdted201.so":   "437af1d919d0c43d5a7e9d4dbc0464de195cee4742c0b13e8608f9ec39ae9a16",
	"libtcl9tkimgflir201.so":   "1211f660a8ca088e835528a9a2ff2b0515106c96d4a9af0def54ec2d76117725",
	"libtcl9tkimggif201.so":    "551519482d774f36613509c2eda78ec7a75622ebe20e9b0b4cc57facff61faaa",
	"libtcl9tkimgico201.so":    "df70087dab4ada85d2f9d6e4dcc148f6331033ab37915aadc349734b9d840167",
	"libtcl9tkimgjpeg201.so":   "b2e92e2f077fe40fe77fe1bb9556e0c521b545ff8f5aec71747ec84b82c35484",
	"libtcl9tkimgpcx201.so":    "6dc0042e2a3053480995920f99b9e1e9c84bafa82431f20a45f1bfbfc646b77f",
	"libtcl9tkimgpixmap201.so": "f1cca252d1c8d08648eb2c8e747f94de31329d63a2b0b4b29184074a353d4cea",
	"libtcl9tkimgpng201.so":    "96b7098220d2d85c17ebd8032678b88d054d17486cfddc8a6e9cc2b6bd9f9bf0",
	"libtcl9tkimgppm201.so":    "5e3bd01f863472dd2d0e76ddf2aa331c15e5630ada80be3623b75688a78682fd",
	"libtcl9tkimgps201.so":     "3c74ddd3f3c414936b17233d78cbc453fcd1341b5ee93453d5da6ae415c2ca13",
	"libtcl9tkimgraw201.so":    "13ef194b485915f5eec8054bd7f2c4947e02866d2c551843c91314bdd4f3ad1f",
	"libtcl9tkimgsgi201.so":    "3b0ad5f68beeab3123b53ca0058f7676f690148b5b9cb83aa418171ea5374151",
	"libtcl9tkimgsun201.so":    "8015b1e862d5607e288486e32af2cf2f5bb0036e18029a7e194ad9044a8b0805",
	"libtcl9tkimgtga201.so":    "ab87b3dec3500e875dc243b09399686808787c969e4281208d071d14a273d12b",
	"libtcl9tkimgtiff201.so":   "d23e0510c881740a5843cc3fb9ab2501451a89a2b7a4b12afc19415e1902793e",
	"libtcl9tkimgwindow201.so": "2a42e121107f17d2d667bfb9c16e49626ea1bf603d3bb915fd139e7ab2eb27a0",
	"libtcl9tkimgxbm201.so":    "fd2447646d4bb8dd5951615c7cf092679de9bf99fa711efbccdb8c8e39921190",
	"libtcl9tkimgxpm201.so":    "689be5106397a478d10c76569d865fb6a3b27e5e67a352136a73716c4f40f446",
	"libtcl9zlibtcl131.so":     "ae9695311027c61b33934dad2f39e7d0d41c8d9a96c629be76c8cf59c45fe508",
	"libtk9.0.1.zip":           "3a82ad0a534b61e9de9968315e2c55b320b9551c35608095a16015584a71738b",
}
