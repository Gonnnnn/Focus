package activity

import (
	"github.com/benbjohnson/clock"
	"gorm.io/gorm"
)

// Define the SQLite implementation of the Repository interface.
type sqlite struct {
    db *gorm.DB
    clock clock.Clock
}

func NewSQLite(db *gorm.DB, clock clock.Clock) Repository {
    return &sqlite{db: db, clock: clock}
}

func (repo *sqlite) CreateActivity(title string, description string, startTimestampMilli int64, endTimestampMilli int64) (*Activity, error) {
    activity := &Activity{
        Title:          title,
        Description:    description,
        StartTimestampMilli: startTimestampMilli,
        EndTimestampMilli:   endTimestampMilli,
        CreatedAt:      repo.clock.Now().UnixMilli(),
        Complete:       false,
    }

    if err := repo.db.Create(activity).Error; err != nil {
        return nil, err
    }

    return activity, nil
}

func (repo *sqlite) Activity(id string) (*Activity, error) {
    var activity Activity
    if err := repo.db.Where("id = ?", id).First(&activity).Error; err != nil {
        return nil, err
    }

    return &activity, nil
}

func (repo *sqlite) Activities(ids []string) ([]*Activity, error) {
    var activities []*Activity
    query := repo.db.Order("start_timestamp DESC")

    if len(ids) > 0 {
        query = query.Where("id IN (?)", ids)
    }

    if err := query.Find(&activities).Error; err != nil {
        return nil, err
    }

    return activities, nil
}

func (repo *sqlite) UpdateActivity(activity *Activity) (*Activity, error) {
    if err := repo.db.Save(activity).Error; err != nil {
        return nil, err
    }

    return activity, nil
}

func (repo *sqlite) DeleteActivity(id string) error {
    if err := repo.db.Where("id = ?", id).Delete(&Activity{}).Error; err != nil {
        return err
    }

    return nil
}

func (repo *sqlite) CompleteActivity(activity *Activity) (*Activity, error) {
    activity.Complete = true
    if err := repo.db.Save(activity).Error; err != nil {
        return nil, err
    }

    return activity, nil
}