package schedule

import "time"

const oneDay = time.Hour * 24
const oneWeek = oneDay * 7

// This package only deals in dates, not times
// so the convention is to always truncate the hour, minute, second, and nano

type Schedule struct {
	date time.Time
}

type Deadlines struct {
	ThisMonth []Deadline
	NextMonth []Deadline
}

type Deadline struct {
	Date        time.Time
	Description string
}

func New(date time.Time) Schedule {
	return Schedule{
		date: time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location()),
	}
}

// Deadlines returns Deadlines, which contains information about dates related to
// this months and next month's jclubs.
// These may be in the past, so it's the responsibility of the caller to decide which are relevant

func (s Schedule) Deadlines() Deadlines {
	firstJclubThisMonth := s.firstThursdayThisMonth()
	secondJclubThisMonth := firstJclubThisMonth.Add(oneWeek * 2)

	firstJclubNextMonth := s.firstThursdayNextMonth()
	return Deadlines{
		ThisMonth: []Deadline{
			{
				Date:        firstJclubThisMonth.Add(oneDay * 11 * -1),
				Description: "Ask for the blurb for the first jclub",
			},
			{
				Date:        firstJclubThisMonth.Add(oneDay * 6 * -1),
				Description: "Send gentle reminder",
			},
			{
				Date:        firstJclubThisMonth.Add(oneDay * 3 * -1),
				Description: "Send email for first club",
			},
			{
				Date:        secondJclubThisMonth.Add(oneDay * 11 * -1),
				Description: "Ask for the blurb for the second jclub",
			},
			{
				Date:        secondJclubThisMonth.Add(oneDay * 6 * -1),
				Description: "Send gentle reminder",
			},
			{
				Date:        secondJclubThisMonth.Add(oneDay * 3 * -1),
				Description: "Send email for second jclub",
			},
		},
		NextMonth: []Deadline{
			{
				Date:        firstJclubNextMonth.Add(oneDay * 18 * -1),
				Description: "Begin asking for next month's jclubs",
			},
			{
				Date:        firstJclubNextMonth.Add(oneDay * 11 * -1),
				Description: "Finalize next month's jclubs",
			},
		},
	}
}

func (s Schedule) firstThursdayThisMonth() time.Time {
	return firstThursdayAfter(s.firstOfTheMonth())
}

func (s Schedule) firstThursdayNextMonth() time.Time {
	return firstThursdayAfter(s.firstOfNextMonth())
}

func firstThursdayAfter(date time.Time) time.Time {
	ret := date
	for ret.Weekday() != time.Thursday {
		ret = ret.AddDate(0, 0, 1)
	}
	return ret
}

func (s Schedule) firstOfTheMonth() time.Time {
	return time.Date(
		s.date.Year(),
		s.date.Month(),
		1,       // day
		0, 0, 0, // hour, min, sec
		0, // nanosecond
		s.date.Location(),
	)
}

func (s Schedule) firstOfNextMonth() time.Time {
	return time.Date(
		s.date.Year(),
		s.date.Month()+1, // time handles the rollover
		1,                // day
		0, 0, 0,          // hour, min, sec
		0, // nanosecond
		s.date.Location(),
	)
}
