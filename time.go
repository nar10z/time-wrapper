package time_wrapper

import (
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
