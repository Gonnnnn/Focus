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
