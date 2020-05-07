package timeular

import (
	"time"
)

// GetTimeEntries returns all timeEntries during a given period
func (a *API) GetTimeEntries(stoppedAfter time.Time, startedBefore time.Time) (*TimeEntries, error) {
	stopafter := stoppedAfter.UTC().Format(TimeFormat)
	startbefore := startedBefore.UTC().Format(TimeFormat)
	dst := new(TimeEntries)
	url := BuildURL(a.url, "/time-entries", stopafter, startbefore)
	err := a.Get(url, dst)
	return dst, err
}
