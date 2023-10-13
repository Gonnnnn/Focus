package impl

import (
	"errors"
	"fmt"
	"focus"
	"focus/activity"

	"github.com/benbjohnson/clock"
)

type impl struct {
	activityRepository activity.Repository
	clock              clock.Clock
}

func New(activityRepository activity.Repository) focus.Focus {
	return &impl{activityRepository: activityRepository, clock: clock.New()}
}

func (i *impl) CreateActivity(title string, description string, startTimestamp int64, endTimestamp int64) (*focus.Activity, error) {
	now := i.clock.Now()
	id := fmt.Sprintf("%d_%d_%d", now.UnixNano(), startTimestamp, endTimestamp)
	activity, err := i.activityRepository.CreateActivity(id, title, description, startTimestamp, endTimestamp)
	if err != nil {
		return nil, err
	}
	return convertToFocusActivity(activity), nil
}

func (i *impl) Activity(id string) (*focus.Activity, error) {
	return &focus.Activity{}, errors.New("not implemented")
}

func (i *impl) Activities(ids []string) ([]*focus.Activity, error) {
	return []*focus.Activity{}, errors.New("not implemented")
}

func (i *impl) UpdateActivity(activity *focus.Activity) (*focus.Activity, error) {
	return &focus.Activity{}, errors.New("not implemented")
}

func (i *impl) DeleteActivity(id string) error {
	return errors.New("not implemented")
}

func convertToFocusActivity(activity *activity.Activity) *focus.Activity {
	return &focus.Activity{
		Id:             activity.Id,
		Title:          activity.Title,
		Description:    activity.Description,
		StartTimestamp: activity.StartTimestamp,
		EndTimestamp:   activity.EndTimestamp,
	}
}
