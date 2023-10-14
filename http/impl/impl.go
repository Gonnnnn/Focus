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

func New(activityRepository activity.Repository, clock clock.Clock) focus.Focus {
	return &impl{activityRepository: activityRepository, clock: clock}
}

func (i *impl) CreateActivity(title string, description string, startTimestamp int64, endTimestamp int64) (*focus.Activity, error) {
	activity, err := i.activityRepository.CreateActivity(title, description, startTimestamp, endTimestamp)
	if err != nil {
		return nil, err
	}
	return convertToFocusActivity(activity), nil
}

func (i *impl) Activity(id string) (*focus.Activity, error) {
	activity, err := i.activityRepository.Activity(id)
	if err != nil {
		return nil, err
	}
	return convertToFocusActivity(activity), nil
}

func (i *impl) Activities(ids []string) ([]*focus.Activity, error) {
	activities, err := i.activityRepository.Activities(ids)
	if err != nil {
		return nil, err
	}
	focusActivities := make([]*focus.Activity, 0, len(activities))
	for _, activity := range activities {
		focusActivities = append(focusActivities, convertToFocusActivity(activity))
	}
	return focusActivities, nil
}

func (i *impl) UpdateActivity(activity *focus.Activity) (*focus.Activity, error) {
	return &focus.Activity{}, errors.New("not implemented")
}

func (i *impl) DeleteActivity(id string) error {
	return i.activityRepository.DeleteActivity(id)
}

func convertToFocusActivity(activity *activity.Activity) *focus.Activity {
	return &focus.Activity{
		Id:            	fmt.Sprintf("%d", activity.Id),
		Title:          activity.Title,
		Description:    activity.Description,
		StartTimestamp: activity.StartTimestamp,
		EndTimestamp:   activity.EndTimestamp,
	}
}
