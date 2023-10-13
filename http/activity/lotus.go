package activity

import (
	"bytes"
	"encoding/gob"
	"errors"

	"github.com/lotusdblabs/lotusdb/v2"
)

type lotus struct {
	db *lotusdb.DB
}

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
	return &Activity{}, errors.New("not implemented")
}

func (l *lotus) Activities(ids string) ([]*Activity, error) {
	return []*Activity{}, errors.New("not implemented")
}

func (l *lotus) UpdateActivity(activity *Activity) (*Activity, error) {
	return &Activity{}, errors.New("not implemented")
}

func (l *lotus) DeleteActivity(id string) error {
	return errors.New("not implemented")
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
