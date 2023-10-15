package focus

type Focus interface {
	// Creates a new activity. It returns the created activity.
	CreateActivity(title string, description string, startTimeStampMilli int64, endTimestampMilli int64) (*Activity, error)
	// Returns the activity with the given id.
	Activity(id string) (*Activity, error)
	// Returns the activities with the given ids. If ids is empty, it returns all the activities.
	Activities(ids []string) ([]*Activity, error)
	// Updates the given activity. It returns the updated activity.
	UpdateActivity(activity *Activity) (*Activity, error)
	// Deletes the activity with the given id.
	DeleteActivity(id string) error
	// Updates the complete status of the activity with the given id. It returns the updated activity.
	CompleteActivity(id string) (*Activity, error)
}

type Activity struct {
	// The unique identifier for the activity. E.g. "abc123"
	Id string
	// The title of the activity. E.g. "My Activity"
	Title string
	// The description of the activity. E.g. "This is my activity. The reason why I started this is ..."
	Description string
	// The start timestamp of the activity in milliseconds. E.g. 1671401696789
	StartTimestampMilli int64
	// The end timestamp of the activity in milliseconds. E.g. 1671401696789
	EndTimestampMilli int64
	// The timestamp when the activity was created in milliseconds. E.g. 1671401696789
	CreatedAt int64
	// The status of the activity. E.g. "IN_PROGRESS"
	Status Status
}

type Status string

const (
	Unknown Status = "UNKNOWN"
	// The activity is not started if the current timestamp is less than the start timestamp.
	NotStarted Status = "NOT_STARTED"
	// The activity is in progress if the current timestamp is greater than the start timestamp.
	InProgress Status = "IN_PROGRESS"
	// The activity is complete if the user completes the activity by themselves.
	Complete Status = "COMPLETE"
	// The activity is expired if the current timestamp is greater than the end timestamp
	// while the activity is not completed yet.
	Expired Status = "EXPIRED"
)