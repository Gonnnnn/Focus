package impl

import (
	"errors"
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
	return ConvertToFocusActivity(activity, i.clock.Now().Unix()), nil
}

func (i *impl) Activity(id string) (*focus.Activity, error) {
	activity, err := i.activityRepository.Activity(id)
	if err != nil {
		return nil, err
	}
	return ConvertToFocusActivity(activity, i.clock.Now().Unix()), nil
}

func (i *impl) Activities(ids []string) ([]*focus.Activity, error) {
	activities, err := i.activityRepository.Activities(ids)
	if err != nil {
		return nil, err
	}
	focusActivities := make([]*focus.Activity, 0, len(activities))
	for _, activity := range activities {
		focusActivities = append(focusActivities, ConvertToFocusActivity(activity, i.clock.Now().Unix()))
	}
	return focusActivities, nil
}

func (i *impl) UpdateActivity(activity *focus.Activity) (*focus.Activity, error) {
	return &focus.Activity{}, errors.New("not implemented")
}

func (i *impl) DeleteActivity(id string) error {
	return i.activityRepository.DeleteActivity(id)
}

func (i *impl) CompleteActivity(id string) (*focus.Activity, error) {
	activity, err := i.activityRepository.Activity(id)
	if err != nil {
		return nil, err
	}

	now := i.clock.Now().Unix()

	if activity.Complete {
		return ConvertToFocusActivity(activity, now), nil
	}

	if activity.StartTimestamp > now {
		return nil, errors.New("expected start time is not yet passed")
	}

	if activity.EndTimestamp < now {
		return nil, errors.New("expected end time is already passed")
	}

	_, err = i.activityRepository.CompleteActivity(activity)
	if err != nil {
		return nil, err
	}

	return  ConvertToFocusActivity(activity, now), nil
}
