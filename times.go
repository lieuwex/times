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
		return nil, err
	}
	return getTimespec(fi), nil
}

// Timespec provides access to file times.
type Timespec interface {
	ModTime() time.Time
	AccessTime() (time time.Time, has bool)
	ChangeTime() (time time.Time, has bool)
	BirthTime() (time time.Time, has bool)
}

type atime struct {
	v time.Time
}

func (a atime) AccessTime() (time.Time, bool) { return a.v, true }

type ctime struct {
	v time.Time
}

func (c ctime) ChangeTime() (time.Time, bool) { return c.v, true }

type mtime struct {
	v time.Time
}

func (m mtime) ModTime() time.Time { return m.v }

type btime struct {
	v time.Time
}

func (b btime) BirthTime() (time.Time, bool) { return b.v, true }

type noctime struct{}

func (noctime) ChangeTime() (time.Time, bool) { return time.Now(), false }

type nobtime struct{}

func (nobtime) BirthTime() (time.Time, bool) { return time.Now(), false }
