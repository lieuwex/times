// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// http://golang.org/src/os/stat_openbsd.go

package times

import (
	"os"
	"syscall"
	"time"
)

// HasChangeTime and HasBirthTime are true if and only if
// the target OS supports them.
const (
	HasChangeTime = true
	HasBirthTime  = false
)

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func getTimespec(fi os.FileInfo) (t Timespec) {
	stat := fi.Sys().(*syscall.Stat_t)

	t.modTime = timespecToTime(stat.Mtimespec)

	t.hasAccessTime = true
	t.accessTime = timespecToTime(stat.Atim)

	t.hasChangeTime = true
	t.changeTime = timespecToTime(stat.Ctim)

	return t
}
