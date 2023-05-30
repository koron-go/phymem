// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build generate
// +build generate

package psapi

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go psapi_windows.go
