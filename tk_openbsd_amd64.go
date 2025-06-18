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
	"lib.zip": "bb574bde782040eedf97731952969ebf5b7c159521694c8050dfaeeaec8498f8",
}
