package chronograph

import (
	"fmt"
	"time"
)

const (
	Layout string = "2006-01-02T15:04:05.000Z"
)

type Day struct {
	Name   string
	Number int
	T      time.Weekday
}

type Month struct {
	Name   string
	Number int
	T      time.Month
}

type Span struct {
	Days    int
	Hours   int
	Minutes int
	Months  int
	Seconds int
	Years   int
}

type Time struct {
	Day    *Day
	Days   int
	Epoch  int64
	Month  *Month
	Second int
	T      time.Time
	Year   int
	Zone   *Zone
}

type Zone struct {
	Name   string
	Offset int
}

func NewDay(time time.Time) *Day {
	t := time.Weekday()
	return &Day{
		Name:   t.String(),
		Number: time.Day(),
		T:      t}
}

func NewMonth(time time.Time) *Month {
	t := time.Month()
	return &Month{
		Name:   t.String(),
		Number: int(t),
		T:      t}
}

func NewSpan(a, b time.Time) *Span {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	var (
		year1, month1, day1     = a.Date()
		year2, month2, day2     = b.Date()
		hour1, minute1, second1 = a.Clock()
		hour2, minute2, second2 = b.Clock()
	)
	var (
		year   = int(year2 - year1)
		month  = int(month2 - month1)
		day    = int(day2 - day1)
		hour   = int(hour2 - hour1)
		minute = int(minute2 - minute1)
		second = int(second2 - second1)
	)
	if second < 0 {
		second = (second + 60)
		minute = (minute - 1)
	}
	if minute < 0 {
		minute = (minute + 60)
		hour = (hour - 1)
	}
	if hour < 0 {
		hour = (hour + 24)
		day = (day - 1)
	}
	if day < 0 {
		day = ((day + 32) - time.Date(year1, month1, 32, 0, 0, 0, 0, time.UTC).Day())
		month = (month - 1)
	}
	if month < 0 {
		month = (month + 12)
		year = (year - 1)
	}
	return &Span{
		Days:    day,
		Hours:   hour,
		Minutes: minute,
		Months:  month,
		Seconds: second,
		Years:   year}
}

func NewSpanFromISO(a, b string) *Span {
	x, err := time.Parse(Layout, a)
	if err != nil {
		panic(err)
	}
	y, err := time.Parse(Layout, b)
	if err != nil {
		panic(err)
	}
	return NewSpan(x, y)
}

func NewTime(time time.Time) *Time {
	return &Time{
		Day:    NewDay(time),
		Days:   time.YearDay(),
		Epoch:  time.Unix(),
		Month:  NewMonth(time),
		Second: time.Second(),
		T:      time,
		Year:   time.Year(),
		Zone:   NewZone(time)}
}

func NewTimeFromISO(ISO string) *Time {
	time, err := time.Parse(Layout, ISO)
	if err != nil {
		fmt.Println(err)
	}
	return NewTime(time)
}

func NewZone(time time.Time) *Zone {
	name, offset := time.Zone()
	return &Zone{
		Name:   name,
		Offset: offset}
}
