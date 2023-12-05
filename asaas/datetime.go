package asaas

import (
	"fmt"
	"time"
)

type Datetime time.Time

var dtLayout = "2006-01-02 15:04:05"

func (d Datetime) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte(fmt.Sprintf(`null`)), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Format())), nil
}

func (d Datetime) Format() string {
	return d.Time().Format(dtLayout)
}

func (d Datetime) Time() time.Time {
	return time.Time(d)
}

func NewDatetime(year int, month time.Month, day, hour, min, sec, nSec int, loc *time.Location) Datetime {
	return Datetime(time.Date(year, month, day, hour, min, sec, nSec, loc))
}

func NewDatetimePointer(year int, month time.Month, day, hour, min, sec, nSec int, loc *time.Location) Datetime {
	v := NewDatetime(year, month, day, hour, min, sec, nSec, loc)
	return v
}

func DatetimeNow() Datetime {
	return Datetime(time.Now())
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

func (d Datetime) Hour() int {
	return d.Time().Hour()
}

func (d Datetime) Minute() int {
	return d.Time().Minute()
}

func (d Datetime) Second() int {
	return d.Time().Second()
}

func (d Datetime) Nanosecond() int {
	return d.Time().Nanosecond()
}

func (d Datetime) Location() *time.Location {
	return d.Time().Location()
}

func (d Datetime) IsZero() bool {
	return d.Time().IsZero()
}

func (d Datetime) date() (year int, month time.Month, day int, yDay int) {
	t := d.Time()
	return t.Year(), t.Month(), t.Day(), t.YearDay()
}
