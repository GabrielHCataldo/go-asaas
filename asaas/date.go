package asaas

import (
	"fmt"
	"strings"
	"time"
)

type Date time.Time

var dLayout = "2006-01-02"

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == `""` {
		return nil
	}
	t, err := time.Parse(dLayout, s)
	if err == nil {
		*d = Date(t)
	}
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte(fmt.Sprintf(`null`)), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Format())), nil
}

func (d Date) Format() string {
	return d.Time().Format(dLayout)
}

func NewDate(year int, month time.Month, day int, loc *time.Location) Date {
	return Date(time.Date(year, month, day, 23, 59, 0, 0, loc))
}

func DateNow() Date {
	return Date(time.Now())
}

func (d Date) Year() int {
	year, _, _, _ := d.date()
	return year
}

func (d Date) Month() time.Month {
	_, month, _, _ := d.date()
	return month
}

func (d Date) Day() int {
	_, _, day, _ := d.date()
	return day
}

func (d Date) Location() *time.Location {
	return d.Time().Location()
}

func (d Date) IsZero() bool {
	return d.Time().IsZero()
}

func (d Date) Time() time.Time {
	return time.Time(d)
}

func (d Date) date() (year int, month time.Month, day int, yDay int) {
	t := d.Time()
	return t.Year(), t.Month(), t.Day(), t.YearDay()
}
