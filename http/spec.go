package focus

type Focus interface {
	// Creates a new activity. It returns the created activity.
	CreateActivity(title string, description string, startTimeStamp int64, endTimestamp int64) (*Activity, error)
	// Returns the activity with the given id.
	Activity(id string) (*Activity, error)
	// Returns the activities with the given ids. If ids is empty, it returns all the activities.
	Activities(ids []string) ([]*Activity, error)
	// Updates the given activity. It returns the updated activity.
	UpdateActivity(activity *Activity) (*Activity, error)
	// Deletes the activity with the given id.
	DeleteActivity(id string) error
}

type Activity struct {
	// The unique identifier for the activity. E.g. "abc123"
	Id string
	// The title of the activity. E.g. "My Activity"
	Title string
	// The description of the activity. E.g. "This is my activity. The reason why I started this is ..."
	Description string
	// The start timestamp of the activity. E.g. 1234567890
	StartTimestamp int64
	// The end timestamp of the activity. E.g. 1234567890
	EndTimestamp int64
}
