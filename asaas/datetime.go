package asaas

import (
	"fmt"
	"strings"
	"time"
)

type Datetime time.Time

var dtLayout = "2006-01-02 15:04:05"

func (d *Datetime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == `""` {
		return nil
	}
	t, err := time.Parse(dtLayout, s)
	*d = Datetime(t)
	return err
}

func (d Datetime) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		return []byte(fmt.Sprintf(`null`)), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Format())), nil
}

func (d Datetime) Format() string {
	t := time.Time(d)
	if t.IsZero() {
		return "null"
	}
	return time.Time(d).Format(dtLayout)
}

func (d Datetime) Time() time.Time {
	return time.Time(d)
}

func NewDateTime(year int, month time.Month, day, hour, min, sec, nSec int, loc *time.Location) Date {
	return Date(time.Date(year, month, day, hour, min, sec, nSec, loc))
}

func (d Datetime) Year() int {
	year, _, _, _ := d.date()
	return year
}

func (d Datetime) Month() time.Month {
	_, month, _, _ := d.date()
	return month
}

func (d Datetime) Day() int {
	_, _, day, _ := d.date()
	return day
}

func (d Datetime) Location() *time.Location {
	return time.Time(d).Location()
}

func (d Datetime) IsZero() bool {
	return time.Time(d).IsZero()
}

func (d Datetime) date() (year int, month time.Month, day int, yDay int) {
	t := time.Time(d)
	return t.Year(), t.Month(), t.Day(), t.YearDay()
}
