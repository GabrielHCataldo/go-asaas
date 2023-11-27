package asaas

import (
	"fmt"
	"strings"
	"time"
)

type DateTime time.Time

var dtLayout = "2006-01-02 15:04:05"

func (d *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == `""` {
		return nil
	}
	t, err := time.Parse(dtLayout, s)
	*d = DateTime(t)
	return err
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		return []byte(fmt.Sprintf(`null`)), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Format())), nil
}

func (d DateTime) Format() string {
	t := time.Time(d)
	if t.IsZero() {
		return "null"
	}
	return time.Time(d).Format(dtLayout)
}

func (d DateTime) Time() time.Time {
	return time.Time(d)
}

func NewDateTime(year int, month time.Month, day, hour, min, sec, nSec int, loc *time.Location) Date {
	return Date(time.Date(year, month, day, hour, min, sec, nSec, loc))
}

func (d DateTime) Year() int {
	year, _, _, _ := d.date()
	return year
}

func (d DateTime) Month() time.Month {
	_, month, _, _ := d.date()
	return month
}

func (d DateTime) Day() int {
	_, _, day, _ := d.date()
	return day
}

func (d DateTime) Location() *time.Location {
	return time.Time(d).Location()
}

func (d DateTime) IsZero() bool {
	return time.Time(d).IsZero()
}

func (d DateTime) date() (year int, month time.Month, day int, yDay int) {
	t := time.Time(d)
	return t.Year(), t.Month(), t.Day(), t.YearDay()
}
