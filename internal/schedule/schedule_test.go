package schedule

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	jan1  = time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)  // 1st Wednesday
	jan2  = time.Date(2025, time.January, 2, 0, 0, 0, 0, time.UTC)  // 1st Thursday
	jan16 = time.Date(2025, time.January, 16, 0, 0, 0, 0, time.UTC) // 3rd Thursday
	jan30 = time.Date(2025, time.January, 30, 0, 0, 0, 0, time.UTC) // 5th Thursday

	jan26 = time.Date(2025, time.January, 26, 0, 0, 0, 0, time.UTC) // 4th Sunday

	feb1 = time.Date(2025, time.February, 1, 0, 0, 0, 0, time.UTC) // 1st Sunday
	feb6 = time.Date(2025, time.February, 6, 0, 0, 0, 0, time.UTC) // 1st Thursday

	feb9 = time.Date(2025, time.February, 9, 0, 0, 0, 0, time.UTC) // 2nd Thursday

	feb16 = time.Date(2025, time.February, 16, 0, 0, 0, 0, time.UTC) // 3rd Sunday
	feb23 = time.Date(2025, time.February, 23, 0, 0, 0, 0, time.UTC) // 4th Sunday
)

func TestScheduleHelpers(t *testing.T) {
	cases := []struct {
		description                      string
		schedule                         Schedule
		expectedFirstOfTheMonth          time.Time
		expectedFirstThursdayOfTheMonth  time.Time
		expectedFirstOfNextMonth         time.Time
		expectedFirstThursdayOfNextMonth time.Time
	}{
		{
			description: "On the first of the month",
			schedule: Schedule{
				date: jan1,
			},
			expectedFirstOfTheMonth:          jan1,
			expectedFirstThursdayOfTheMonth:  jan2,
			expectedFirstOfNextMonth:         feb1,
			expectedFirstThursdayOfNextMonth: feb6,
		},
		{
			description: "On the third thursday of the month",
			schedule: Schedule{
				date: jan16,
			},
			expectedFirstOfTheMonth:          jan1,
			expectedFirstThursdayOfTheMonth:  jan2,
			expectedFirstOfNextMonth:         feb1,
			expectedFirstThursdayOfNextMonth: feb6,
		},
		{
			description: "On the last thursday of the month",
			schedule: Schedule{
				date: jan30,
			},
			expectedFirstOfTheMonth:          jan1,
			expectedFirstThursdayOfTheMonth:  jan2,
			expectedFirstOfNextMonth:         feb1,
			expectedFirstThursdayOfNextMonth: feb6,
		},
	}
	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			actualFirstOfTheMonth := tc.schedule.firstOfTheMonth()
			assert.Equal(t, tc.expectedFirstOfTheMonth, actualFirstOfTheMonth)

			actualFirstOfNextMonth := tc.schedule.firstOfNextMonth()
			assert.Equal(t, tc.expectedFirstOfNextMonth, actualFirstOfNextMonth)

			actualFirstThursdayOfTheMonth := tc.schedule.firstThursdayThisMonth()
			assert.Equal(t, tc.expectedFirstThursdayOfTheMonth, actualFirstThursdayOfTheMonth)

			actualFirstThursdayOfNextMonth := tc.schedule.firstThursdayNextMonth()
			assert.Equal(t, tc.expectedFirstThursdayOfNextMonth, actualFirstThursdayOfNextMonth)
		})
	}
}

func TestScheduleDeadlines(t *testing.T) {
	cases := []struct {
		description       string
		schedule          Schedule
		expectedDeadlines Deadlines
	}{
		{
			description: "In February",
			schedule: Schedule{
				date: feb1,
			},
			expectedDeadlines: Deadlines{
				ThisMonth: []Deadline{
					{
						Date:        jan26,
						Description: "Ask for the blurb for the first jclub",
					},
					{
						Date:        feb9,
						Description: "Ask for the blurb for the second jclub",
					},
				},
				NextMonth: []Deadline{
					{
						Date:        feb16,
						Description: "Begin asking for next month's jclubs",
					},
					{
						Date:        feb23,
						Description: "Finalize next month's jclubs",
					},
				},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			actualDeadlines := tc.schedule.Deadlines()
			assert.Equal(t, tc.expectedDeadlines, actualDeadlines)
		})
	}
}
