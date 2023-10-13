package activity

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/lotusdblabs/lotusdb/v2"
	"go.uber.org/multierr"
)

type lotus struct {
	db *lotusdb.DB
}

var KEY_OF_KEYS = []byte("KEYS")
var KEY_DELIMETER = ","

// NewLotus creates a new Lotus repository.
// Lotus is a key-value store. It doesn't support scan operation. So, we need to maintain the keys.
// We use a key named "KEYS" to store all the keys. The value of "KEYS" is a string which is a list of keys separated by ",".
// If "KEYS" doesn't exist, it has to be created before using the repository.
func NewLotus(db *lotusdb.DB) (Repository, error) {
	_, err := db.Get([]byte(KEY_OF_KEYS))
	if err != nil {
		if err.Error() != "key not found in database" {
			return nil, fmt.Errorf("failed to get KEYS: %w", err)
		}
		err = db.Put([]byte(KEY_OF_KEYS), []byte("DUMMY_KEY"), nil)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize KEYS: %w", err)
		}
	}
	return &lotus{db: db}, nil
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

	sort.Slice(activities, func(i, j int) bool {
		return activities[i].StartTimestamp > activities[j].StartTimestamp
	})
	return activities, nil
}

func (l *lotus) UpdateActivity(activity *Activity) (*Activity, error) {
	return &Activity{}, errors.New("not implemented")
}

func (l *lotus) DeleteActivity(id string) error {
	return errors.New("not implemented")
}

func (l *lotus) keys() ([]string, error) {
	data, err := l.db.Get([]byte(KEY_OF_KEYS))
	if err != nil {
		return nil, fmt.Errorf("failed to get keys: %w", err)
	}
	keys := strings.Split(string(data), KEY_DELIMETER)
	// Remove the first element which is a dummy key.
	return keys[1:], nil
}

func (l *lotus) createKey(key string) error {
	data, err := l.db.Get([]byte(KEY_OF_KEYS))
	if err != nil {
		return fmt.Errorf("failed to get keys: %w", err)
	}

	keys := string(data)
	err = l.db.Put([]byte(KEY_OF_KEYS), []byte(fmt.Sprintf("%s%s%s", keys, KEY_DELIMETER, key)), nil)
	if err != nil {
		return fmt.Errorf("failed to put keys: %w", err)
	}
	return nil
}
