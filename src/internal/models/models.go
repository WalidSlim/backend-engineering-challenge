package models

import (
	"strings"
	"time"
)

type TruncatedTimeAndDuration struct {
	Timestamp Time `json:"timestamp,omitempty"`
	Duration  int  `json:"duration,omitempty"`
}

type MovingAverageResult struct {
	Timestamp           Time    `json:"date,omitempty"`
	AverageDeliveryTime float64 `json:"average_delivery_time"`
}

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	if s == "null" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse(time.DateTime, s)
	return
}

type Translation struct {
	Timestamp      Time   `json:"timestamp,omitempty"`
	TranslationId  string `json:"translation_id,omitempty"`
	SourceLanguage string `json:"source_language,omitempty"`
	TargetLanguage string `json:"target_language,omitempty"`
	ClientName     string `json:"client_name,omitempty"`
	EventName      string `json:"event_name,omitempty"`
	NrWords        int    `json:"nr_words,omitempty"`
	Duration       int    `json:"duration,omitempty"`
}

// Method to go from continuous time to discrete time and only keep duration
func (t *Translation) FormatToTimeAndDuration() TruncatedTimeAndDuration {
	truncated := t.Timestamp.Truncate(time.Minute).Add(time.Minute)
	truncatedTime := Time{Time: truncated}
	timeAndDuration := TruncatedTimeAndDuration{Timestamp: truncatedTime, Duration: t.Duration}
	return timeAndDuration
}
