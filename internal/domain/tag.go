package domain

import "time"

type TagReaded struct {
	Mac      string    `json:"mac,omitempty"`
	Tag      string    `json:"tag,omitempty"`
	Antenna  uint      `json:"antenna,omitempty"`
	ReaderAt time.Time `json:"reader_at,omitempty"`
	Location string    `json:"location,omitempty"`
}
