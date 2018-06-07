// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// http://golang.org/src/os/stat_plan9.go

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
	HasBirthTime  = false
)

func getTimespec(fi os.FileInfo) (t Timespec) {
	stat := fi.Sys().(*syscall.Dir)

	t.modTime = timespecToTime(stat.Mtimespec)

	t.hasAccessTime = true
	t.accessTime = time.Unix(int64(stat.Atime), 0)

	return t
}
