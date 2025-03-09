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
	"libtcl9.0.so":    "01e7c42546fe12c6f75a0e7803808dd1d2b93cdfe3ffaabc8130f2f807e09fd0",
	"libtcl9tk9.0.so": "847a804bde8273f0c768a1aeef94e5fff1522ce7ceb5334e013ff49f2bbe34d6",
	"libtk9.0.1.zip":  "7c0a6d17bb157241285228e5c2772b8ce7cbb0819fab939b2503e43ab374480f",
}
