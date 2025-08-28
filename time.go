package time_wrapper

import (
	"bytes"
	"time"

	"github.com/goccy/go-json"
)

const customLayout = "20060102T150405"

type TimeWrapper struct {
	time.Time
}

func (tw *TimeWrapper) UnmarshalJSON(data []byte) error {
	// Обработка null
	if len(data) == 0 || BytesToString(data) == "null" {
		*tw = TimeWrapper{}
		return nil
	}

	var date string
	if err := json.Unmarshal(data, &date); err != nil {
		return err
	}

	t, err := time.Parse(customLayout, date)
	if err != nil {
		return err
	}

	tw.Time = t

	return nil
}

func (tw TimeWrapper) MarshalJSON() ([]byte, error) {
	if tw.Time.IsZero() {
		return StringToBytes("null"), nil
	}

	formatted := tw.Format(customLayout)
	return json.Marshal(formatted)
}

type DateTimeNoTZ struct{ time.Time }

var jsonNull = []byte("null")

func (t *DateTimeNoTZ) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, jsonNull) {
		*t = DateTimeNoTZ{}

		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	v, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}

	t.Time = v

	return nil
}

func (t DateTimeNoTZ) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return jsonNull, nil
	}

	return json.Marshal(t.Time.Format("2006-01-02T15:04:05"))
}
