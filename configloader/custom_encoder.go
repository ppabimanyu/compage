package configloader

import (
	"time"
)

type TimeFormat time.Duration

func (tf *TimeFormat) Decode(value string) error {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	*tf = TimeFormat(duration)
	return nil
}

func (tf *TimeFormat) Duration() time.Duration {
	return time.Duration(*tf)
}
