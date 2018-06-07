// Package times provides a platform-independent way to get atime, mtime, ctime and btime for files.
package times

import (
	"os"
	"time"
)

// Get returns the Timespec for the given FileInfo
func Get(fi os.FileInfo) Timespec {
	return getTimespec(fi)
}

// Stat returns the Timespec for the given filename.
func Stat(name string) (Timespec, error) {
	if hasPlatformSpecificStat {
		if ts, err := platformSpecficStat(name); err == nil {
			return ts, nil
		}
	}

	fi, err := os.Stat(name)
	if err != nil {
		return Timespec{}, err
	}

	return getTimespec(fi), nil
}

// Timespec provides access to file times.
type Timespec struct {
	modTime time.Time

	hasAccessTime bool
	accessTime    time.Time

	hasChangeTime bool
	changeTime    time.Time

	hasBirthTime bool
	birthTime    time.Time
}

func (ts Timespec) ModTime() time.Time {
	return ts.modTime
}

func (ts Timespec) AccessTime() (time.Time, bool) {
	return ts.accessTime, ts.hasAccessTime
}

func (ts Timespec) ChangeTime() (time.Time, bool) {
	return ts.changeTime, ts.hasChangeTime
}

func (ts Timespec) BirthTime() (time.Time, bool) {
	return ts.birthTime, ts.hasBirthTime
}
