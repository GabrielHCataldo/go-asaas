package asaas

import (
	"encoding/json"
	"time"
)

type Date time.Time

func (a *Date) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("2006-01-02", string(b))
	if err != nil {
		return err
	}
	*a = Date(t)
	return nil
}

func (a *Date) MarshalJSON() ([]byte, error) {
	if a == nil {
		return nil, nil
	}
	return json.Marshal(time.Time(*a))
}

func (a *Date) Format(s string) string {
	if a == nil {
		return ""
	}
	t := time.Time(*a)
	return t.Format(s)
}

func (a *Date) Time() *time.Time {
	if a == nil {
		return nil
	}
	t := time.Time(*a)
	return &t
}

func NewDate(year int, month time.Month, day, hour, min, sec, nSec int, loc *time.Location) Date {
	return Date(time.Date(year, month, day, hour, min, sec, nSec, loc))
}

func Now() Date {
	return Date(time.Now())
}

func (a *Date) Year() int {
	year, _, _, _ := a.date()
	return year
}

func (a *Date) Month() time.Month {
	_, month, _, _ := a.date()
	return month
}

func (a *Date) Day() int {
	_, _, day, _ := a.date()
	return day
}

func (a *Date) Location() *time.Location {
	if a == nil {
		return nil
	}
	t := time.Time(*a)
	return t.Location()
}

func (a *Date) date() (year int, month time.Month, day int, yDay int) {
	if a == nil {
		return 0, time.January, 1, 1
	}
	t := time.Time(*a)
	return t.Year(), t.Month(), t.Day(), t.YearDay()
}
