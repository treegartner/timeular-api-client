package timeular

// ReadActivities returns all Activities in an Activity array
func (a *API) ReadActivities() (*Activities, error) {
	dst := new(Activities)
	url := a.BuildURL("/activities")
	err := a.Get(url, dst)
	return dst, err
}

// GetActivityByName returns the activity by its name or nil
func (a *API) GetActivityByName(name string) (Activity, error) {
	activites, err := a.ReadActivities()
	if err != nil {
		return Activity{}, err
	}
	for _, a := range activites.Activities {
		if a.Name == name {
			return a, nil
		}
	}

	return Activity{}, nil
}

// GetActivityByID returns the activity by its ID or nil
func (a *API) GetActivityByID(id string) (Activity, error) {
	activites, err := a.ReadActivities()
	if err != nil {
		return Activity{}, err
	}
	for _, a := range activites.Activities {
		if a.ID == id {
			return a, nil
		}
	}

	return Activity{}, nil
}
