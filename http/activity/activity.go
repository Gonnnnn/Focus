package activity

type Repository interface {
	// Creates a new activity.
	CreateActivity(title string, description string, startTimestamp int64, endTimestamp int64) (*Activity, error)
	// Returns the activity with the given id.
	Activity(id string) (*Activity, error)
	// Returns the activities with the given ids in the descending order of the start timestamp. If ids is empty, it returns all the activities.
	Activities(ids []string) ([]*Activity, error)
	// Updates the given activity.
	UpdateActivity(activity *Activity) (*Activity, error)
	// Deletes the activity with the given id.
	DeleteActivity(id string) error
	// Updates the done status of the activity with the given id.
	CompleteActivity(activity *Activity) (*Activity, error)
}

type Activity struct {
	// The unique identifier for the activity. E.g. 1
	Id uint `gorm:"primaryKey;autoIncrement"`
	// The title of the activity. E.g. "My Activity"
	Title string `gorm:"column:title;not null"`
	// The description of the activity. E.g. "This is my activity. The reason why I started this is ..."
	Description string `gorm:"column:description;not null"`
	// The start timestamp of the activity. E.g. 1234567890
	StartTimestamp int64 `gorm:"column:start_timestamp;not null"`
	// The end timestamp of the activity. E.g. 1234567890
	EndTimestamp int64 `gorm:"column:end_timestamp;not null"`
	// The timestamp when the activity was created. E.g. 1234567890
	CreatedAt int64 `gorm:"column:created_at;not null"`
	// The indicator whether the activity is done or not. E.g. true
	Complete bool `gorm:"column:complete;not null"`
}
