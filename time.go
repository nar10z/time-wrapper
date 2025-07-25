package time_wrapper

import (
	"time"
)

const customLayout = "20060102T150405"

type TimeWrapper struct {
	time.Time
}

func (tw *TimeWrapper) UnmarshalJSON(data []byte) error {
	date := BytesToString(data)

	t, err := time.Parse(customLayout, date)
	if err != nil {
		return err
	}

	tw.Time = t

	return nil
}

func (tw TimeWrapper) MarshalJSON() ([]byte, error) {
	return StringToBytes(tw.Time.Format(customLayout)), nil
}
