package appTime

import (
	"strings"
	"time"
)

type client struct{}

func New() *client {
	return &client{}
}

func (t client) NowTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}

func (t client) NowDateTime() string {
	return time.Now().Format(time.RFC3339)
}

func (t client) AddToTimestamp(years int, months int, days int) int64 {
	today := time.Now()
	return today.AddDate(years, months, days).UnixNano() / int64(time.Second)
}

func (t client) GetEpochFromDate(date string) (int64, error) {
	rfc3339t := date + "T00:00:00Z"
	d, err := time.Parse(time.RFC3339, rfc3339t)
	if err != nil {
		return 0, err
	}
	
	return d.Unix(), err
}

func (t client) GetEpochFromDateAndTime(datetime string) (int64, error) {
	rfc3339t := strings.Replace(datetime, " ", "T", 1)
	d, err := time.Parse(time.RFC3339, rfc3339t)
	if err != nil {
		return 0, err
	}
	
	return d.Unix(), err
}

type Port interface {
	NowTimestamp() int64
	AddToTimestamp(years int, months int, days int) int64
	GetEpochFromDate(date string) (int64, error)
	GetEpochFromDateAndTime(datetime string) (int64, error)
	NowDateTime() string
}
