package impl

import (
	"fmt"
	"focus"
	"focus/activity"
)


func ConvertToFocusActivity(activity *activity.Activity, nowMilli int64) *focus.Activity {
	return &focus.Activity{
		Id:            	fmt.Sprintf("%d", activity.Id),
		Title:          activity.Title,
		Description:    activity.Description,
		StartTimestampMilli: activity.StartTimestampMilli,
		EndTimestampMilli:   activity.EndTimestampMilli,
		CreatedAt:      activity.CreatedAt,
		Status:         activityStatus(activity, nowMilli),
	}
}

func activityStatus(activity *activity.Activity, nowMilli int64) focus.Status {
	if activity.Complete {
		return focus.Complete
	}

	fmt.Printf("activity.StartTimestampMilli: %d\n", activity.StartTimestampMilli)
	fmt.Printf("nowMilli: %d\n", nowMilli)
	if activity.StartTimestampMilli > nowMilli {
		return focus.NotStarted
	}

	if activity.EndTimestampMilli < nowMilli {
		return focus.Expired
	}

	return focus.InProgress
}