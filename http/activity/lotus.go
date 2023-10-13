package activity

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"strings"

	"github.com/lotusdblabs/lotusdb/v2"
	"github.com/uber-go/multierr"
	"go.uber.org/multierr"
)

type lotus struct {
	db *lotusdb.DB
}

var KEY_DELIMETER = ","

func NewLotus() *lotus {
	return &lotus{db: nil}
}

func (l *lotus) CreateActivity(id string, title string, description string, startTimestamp int64, endTimestamp int64) (*Activity, error) {
	activity := &Activity{
		Id:             id,
		Title:          title,
		Description:    description,
		StartTimestamp: startTimestamp,
		EndTimestamp:   endTimestamp,
	}

	buffer := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(activity)
	if err != nil {
		return nil, err
	}

	err = l.db.Put([]byte(id), buffer.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	err = l.createKey(id)
	if err != nil {
		return nil, fmt.Errorf("failed to add the key[%s]: %w", id, err)
	}
	return activity, nil
}

func (l *lotus) Activity(id string) (*Activity, error) {
	data, err := l.db.Get([]byte(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get activity[%s]: %w", id, err)
	}

	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	activity := &Activity{}
	err = decoder.Decode(activity)
	if err != nil {
		return nil, fmt.Errorf("failed to decode activity[%s]: %w", id, err)
	}
	return activity, nil
}

func (l *lotus) Activities(ids []string) ([]*Activity, error) {
	if len(ids) == 0 {
		existingIds, err := l.keys()
		if err != nil {
			return nil, fmt.Errorf("failed to get keys: %w", err)
		}
		ids = existingIds
	}

	activities := make([]*Activity, 0, len(ids))
	var errs error
	for _, id := range ids {
		activity, err := l.Activity(id)
		if err != nil {
			errs = multierr.Append(errs, err)
			continue
		}
		activities = append(activities, activity)
	}
	if errs != nil {
		return activities, fmt.Errorf("failed to get activities: %w", errs)
	}
	return activities, nil
}

func (l *lotus) UpdateActivity(activity *Activity) (*Activity, error) {
	return &Activity{}, errors.New("not implemented")
}

func (l *lotus) DeleteActivity(id string) error {
	return errors.New("not implemented")
}

func (l *lotus) keys() ([]string, error) {
	data, err := l.db.Get([]byte("KEYS"))
	if err != nil {
		return nil, fmt.Errorf("failed to get keys: %w", err)
	}
	return strings.Split(string(data), KEY_DELIMETER), nil
}

func (l *lotus) createKey(key string) error {
	data, err := l.db.Get([]byte("KEYS"))
	if err != nil {
		return fmt.Errorf("failed to get keys: %w", err)
	}

	keys := string(data)
	err = l.db.Put([]byte("KEYS"), []byte(fmt.Sprintf("%s%s%s", keys, KEY_DELIMETER, key)), nil)
	if err != nil {
		return fmt.Errorf("failed to put keys: %w", err)
	}
	return nil
}
