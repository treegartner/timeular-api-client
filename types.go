package timeular

// Auth holds the token that is sent after authentication
type Auth struct {
	Token string `json:"token"`
}

type Activity struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Integration string `json:"integration"`
}

type Activities struct {
	Activities []Activity `json:"activities"`
}

type ArchivedActivities struct {
	ArchivedActivities []Activity `json:"archivedActivities"`
}

type Duration struct {
	StartedAt string `json:"startedAt"`
	StoppedAt string `json:"stoppedAt"`
}

type Errors struct {
	Errors []string `json:"errors"`
}

// Message contains the error message if an error occurred
type Message struct {
	Message string `json:"message"`
}

type Note struct {
	Text string `json:"text"`
	Tags []struct {
		Indices []int  `json:"indices"`
		Key     string `json:"key"`
	} `json:"tags"`
	Mentions []struct {
		Indices []int  `json:"indices"`
		Key     string `json:"key"`
	} `json:"mentions"`
}

type TimeEntry struct {
	ID       string   `json:"id"`
	Activity Activity `json:"activity"`
	Duration Duration `json:"duration"`
	Note     Note     `json:"note"`
}

type Tracking struct {
	CurrentTracking struct {
		Activity Activity `json:"activity,omitempty"`
	} `json:"currentTracking,omitempty"`
	StartedAt string `json:"startedAt"`
	Note      Note   `json:"note"`
	Errors    Errors `json:"errors"`
}

type CreatedTimeEntry struct {
	CreatedTimeEntry TimeEntry `json:"createdTimeEntry"`
}

type TimeEntries struct {
	TimeEntries []TimeEntry `json:"timeEntries"`
}
