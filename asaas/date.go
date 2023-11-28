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
	*d = Date(t)
	return err
}

func (d Date) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		return []byte(fmt.Sprintf(`null`)), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Format())), nil
}

func (d Date) Format() string {
	t := time.Time(d)
	if t.IsZero() {
		return "null"
	}
	return time.Time(d).Format(dLayout)
}

func (d Date) Time() time.Time {
	return time.Time(d)
}

func NewDate(year int, month time.Month, day int, loc *time.Location) Date {
	return Date(time.Date(year, month, day, 23, 59, 0, 0, loc))
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
	return time.Time(d).Location()
}

func (d Date) IsZero() bool {
	return time.Time(d).IsZero()
}

func (d Date) date() (year int, month time.Month, day int, yDay int) {
	t := time.Time(d)
	return t.Year(), t.Month(), t.Day(), t.YearDay()
}
