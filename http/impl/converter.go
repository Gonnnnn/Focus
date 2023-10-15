package impl

import (
	"fmt"
	"focus"
	"focus/activity"
)


func ConvertToFocusActivity(activity *activity.Activity, now int64) *focus.Activity {
	return &focus.Activity{
		Id:            	fmt.Sprintf("%d", activity.Id),
		Title:          activity.Title,
		Description:    activity.Description,
		StartTimestamp: activity.StartTimestamp,
		EndTimestamp:   activity.EndTimestamp,
		CreatedAt:      activity.CreatedAt,
		Status:         activityStatus(activity, now),
	}
}

func activityStatus(activity *activity.Activity, now int64) focus.Status {
	if activity.Complete {
		return focus.Complete
	}

	if activity.StartTimestamp > now {
		return focus.NotStarted
	}

	if activity.EndTimestamp < now {
		return focus.Expired
	}

	return focus.InProgress
}