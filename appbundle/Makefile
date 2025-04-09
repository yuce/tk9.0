# Copyright 2025 The tk9.0-go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean edit editor

all:
	golint 2>&1
	staticcheck 2>&1

clean:
	rm -f log-* cpu.test mem.test *.out go.work*
	go clean

edit:
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile *.go & fi

editor:
	go test -vet=off -c -o /dev/null
	gofmt -l -s -w .
	go install
