package focus

type Focus interface {
	CreateActivity(title string, description string, endTimestamp int64) (*Activity, error)
	Activity(id string) (*Activity, error)
	Activities(ids []string) ([]*Activity, error)
	UpdateActivity(activity *Activity) (*Activity, error)
	DeleteActivity(id string) error
}

type Activity struct {
	Id             string
	Title          string
	Description    string
	StartTimestamp int64
	EndTimestamp   int64
}
