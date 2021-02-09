package models

import (
	"time"

	"github.com/pkg/errors"
)

type Date struct {
	time.Time
}

func (d Date) MarshalYAML() (interface{}, error) {
	return d.Format("2006-01-02"), nil
}

func (d *Date) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return errors.WithStack(err)
	}

	t, err := time.Parse("2006-01-02", buf)
	if err != nil {
		return errors.WithStack(err)
	}

	*d = Date{t}
	return nil
}
