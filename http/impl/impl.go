package impl

import (
	"errors"
	"focus"
)

type impl struct{}

func New() focus.Focus {
	return &impl{}
}

func (i *impl) CreateActivity(title string, description string, endTimestamp int64) (*focus.Activity, error) {
	return &focus.Activity{}, errors.New("not implemented")
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
