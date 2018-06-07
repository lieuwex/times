// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// http://golang.org/src/os/stat_windows.go

package times

import (
	"os"
	"syscall"
	"time"
)

// HasChangeTime and HasBirthTime are true if and only if
// the target OS supports them.
const (
	HasChangeTime = false
	HasBirthTime  = true
)

func getTimespec(fi os.FileInfo) (t Timespec) {
	stat := fi.Sys().(*syscall.Win32FileAttributeData)

	t.modTime = timespecToTime(stat.Mtimespec)

	t.hasAccessTime = true
	t.accessTime = time.Unix(0, stat.LastAccessTime.Nanoseconds())

	t.hasBirthTime = true
	t.birthTime = time.Unix(0, stat.CreationTime.Nanoseconds())

	return t
}
