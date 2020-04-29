package timeular

import (
	"fmt"
	"time"
)

// CurrentTracking returns the activity by its name or nil
func (a *API) CurrentTracking() (*Tracking, error) {
	dst := new(Tracking)
	url := a.BuildURL("/tracking")
	err := a.Get(url, dst)
	return dst, err
}

// StartTracking starts the activity given by Name
func (a *API) StartTracking(activityName string) (*Tracking, error) {
	tracking, err := a.CurrentTracking()
	if err != nil {
		return nil, err
	}
	// check if tracking is active
	if (tracking.CurrentTracking.Activity != Activity{}) {
		if tracking.CurrentTracking.Activity.Name == activityName {
			fmt.Printf("DEBUG: Wanted activity is already running. %v", tracking)
			return tracking, nil
		}
		_, err = a.StopTracking()
		if err != nil {
			return nil, err
		}
	}

	// Start Tracking request
	act, err := a.GetActivityByName(activityName)
	if err != nil {
		return nil, err
	}
	var startTime = struct {
		StartedAt string `json:"startedAt"`
	}{}
	startTime.StartedAt = time.Now().UTC().Format("2006-01-02T15:04:05.000")
	dst := new(CreatedTimeEntry)
	url := a.BuildURL("/tracking/" + act.ID + "/start")
	err = a.Post(url, startTime, &dst)
	if err != nil {
		return nil, err
	}

	return tracking, nil
}

// StopTracking stops any current tracking activity
func (a *API) StopTracking() (*CreatedTimeEntry, error) {
	tracking, err := a.CurrentTracking()
	if err != nil {
		return nil, err
	}
	if (tracking.CurrentTracking.Activity == Activity{}) {
		return nil, err
	}

	// Stop Tracking request
	var stopTime = struct {
		StoppedAt string `json:"stoppedAt"`
	}{}
	stopTime.StoppedAt = time.Now().UTC().Format("2006-01-02T15:04:05.000")
	dst := new(CreatedTimeEntry)
	url := a.BuildURL("/tracking/" + tracking.CurrentTracking.Activity.ID + "/stop")
	err = a.Post(url, stopTime, &dst)
	if err != nil {
		return nil, err
	}

	return dst, nil
}

// ToggleTracking starts or stops tracking of an activity
func (a *API) ToggleTracking(activityName string) error {
	tracking, err := a.CurrentTracking()
	if err != nil {
		return err
	}
	if (tracking.CurrentTracking.Activity == Activity{}) {
		_, err = a.StartTracking(activityName)
		if err != nil {
			return err
		}
		return nil
	}

	stoppedActivityName := tracking.CurrentTracking.Activity.Name
	_, err = a.StopTracking()
	if err != nil {
		return err
	}

	if stoppedActivityName != activityName {
		_, err = a.StartTracking(activityName)
		if err != nil {
			return err
		}
	}

	return nil
}
