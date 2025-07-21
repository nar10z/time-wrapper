package fuseps

import (
	"time"

	"github.com/goccy/go-json"
)

var raw struct {
	Date string `json:"__datetime"`
}

type TimeWrapper struct {
	time.Time
}

func (tw *TimeWrapper) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	t, err := time.Parse("20060102T150405", raw.Date)
	if err != nil {
		return err
	}

	tw.Time = t

	return nil
}

func (tw TimeWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"__datetime": tw.Time.Format("20060102T150405"),
	})
}
