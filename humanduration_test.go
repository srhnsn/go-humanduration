package humanduration

import (
	"testing"
	"time"
)

func TestAbsoluteDurationsGerman(t *testing.T) {
	humanDuration := NewGermanFormatter()

	for duration, expectedHumanizedDuration := range absoluteDurationsGerman {
		actualHumanizedDuration := humanDuration.HumanizeAbsolute(duration)

		if actualHumanizedDuration != expectedHumanizedDuration {
			t.Errorf(`expected absolute humanized duration for "%s" was "%s", got "%s"`,
				duration, expectedHumanizedDuration, actualHumanizedDuration)
		}
	}
}

func TestRelativeDurationsGerman(t *testing.T) {
	humanDuration := NewGermanFormatter()

	for duration, expectedHumanizedDuration := range relativeDurationsGerman {
		actualHumanizedDuration := humanDuration.HumanizeRelative(duration)

		if actualHumanizedDuration != expectedHumanizedDuration {
			t.Errorf(`expected relative humanized duration for "%s" was "%s", got "%s"`,
				duration, expectedHumanizedDuration, actualHumanizedDuration)
		}
	}
}

func TestAbsoluteDurationsEnglish(t *testing.T) {
	humanDuration := NewEnglishFormatter()

	for duration, expectedHumanizedDuration := range absoluteDurationsEnglish {
		actualHumanizedDuration := humanDuration.HumanizeAbsolute(duration)

		if actualHumanizedDuration != expectedHumanizedDuration {
			t.Errorf(`expected absolute humanized duration for "%s" was "%s", got "%s"`,
				duration, expectedHumanizedDuration, actualHumanizedDuration)
		}
	}
}

func TestRelativeDurationsEnglish(t *testing.T) {
	humanDuration := NewEnglishFormatter()

	for duration, expectedHumanizedDuration := range relativeDurationsEnglish {
		actualHumanizedDuration := humanDuration.HumanizeRelative(duration)

		if actualHumanizedDuration != expectedHumanizedDuration {
			t.Errorf(`expected relative humanized duration for "%s" was "%s", got "%s"`,
				duration, expectedHumanizedDuration, actualHumanizedDuration)
		}
	}
}

var absoluteDurationsEnglish = map[time.Duration]string{
	0: "0 seconds",

	1 * time.Second:  "a second",
	2 * time.Second:  "2 seconds",
	3 * time.Second:  "3 seconds",
	59 * time.Second: "59 seconds",

	1 * time.Minute:  "a minute",
	2 * time.Minute:  "2 minutes",
	3 * time.Minute:  "3 minutes",
	59 * time.Minute: "59 minutes",

	1 * time.Hour:  "an hour",
	2 * time.Hour:  "2 hours",
	3 * time.Hour:  "3 hours",
	23 * time.Hour: "23 hours",

	24 * time.Hour: "a day",
	48 * time.Hour: "2 days",
	72 * time.Hour: "3 days",

	24 * time.Hour * 6:  "6 days",
	24 * time.Hour * 7:  "a week",
	24 * time.Hour * 8:  "a week",
	24 * time.Hour * 13: "a week",

	24 * time.Hour * 7 * 2: "2 weeks",
	24 * time.Hour * 7 * 3: "3 weeks",
	24 * time.Hour * 7 * 4: "4 weeks",

	24 * time.Hour * 7 * 5: "a month",
	24 * time.Hour * 60:    "2 months",
	24 * time.Hour * 89:    "2 months",
	24 * time.Hour * 90:    "3 months",
	24 * time.Hour * 364:   "12 months",

	24 * time.Hour * 365:     "a year",
	24 * time.Hour * 365 * 2: "2 years",
	24 * time.Hour * 365 * 3: "3 years",
}

var relativeDurationsEnglish = map[time.Duration]string{
	-1 * time.Second:  "a second ago",
	-2 * time.Second:  "2 seconds ago",
	-3 * time.Second:  "3 seconds ago",
	-59 * time.Second: "59 seconds ago",

	-1 * time.Minute:  "a minute ago",
	-2 * time.Minute:  "2 minutes ago",
	-3 * time.Minute:  "3 minutes ago",
	-59 * time.Minute: "59 minutes ago",

	-1 * time.Hour:  "an hour ago",
	-2 * time.Hour:  "2 hours ago",
	-3 * time.Hour:  "3 hours ago",
	-23 * time.Hour: "23 hours ago",

	-24 * time.Hour: "a day ago",
	-48 * time.Hour: "2 days ago",
	-72 * time.Hour: "3 days ago",

	-24 * time.Hour * 6:  "6 days ago",
	-24 * time.Hour * 7:  "a week ago",
	-24 * time.Hour * 8:  "a week ago",
	-24 * time.Hour * 13: "a week ago",

	-24 * time.Hour * 7 * 2: "2 weeks ago",
	-24 * time.Hour * 7 * 3: "3 weeks ago",
	-24 * time.Hour * 7 * 4: "4 weeks ago",

	-24 * time.Hour * 7 * 5: "a month ago",
	-24 * time.Hour * 60:    "2 months ago",
	-24 * time.Hour * 89:    "2 months ago",
	-24 * time.Hour * 90:    "3 months ago",
	-24 * time.Hour * 364:   "12 months ago",

	-24 * time.Hour * 365:     "a year ago",
	-24 * time.Hour * 365 * 2: "2 years ago",
	-24 * time.Hour * 365 * 3: "3 years ago",

	0: "in 0 seconds",

	1 * time.Second:  "in a second",
	2 * time.Second:  "in 2 seconds",
	3 * time.Second:  "in 3 seconds",
	59 * time.Second: "in 59 seconds",

	1 * time.Minute:  "in a minute",
	2 * time.Minute:  "in 2 minutes",
	3 * time.Minute:  "in 3 minutes",
	59 * time.Minute: "in 59 minutes",

	1 * time.Hour:  "in an hour",
	2 * time.Hour:  "in 2 hours",
	3 * time.Hour:  "in 3 hours",
	23 * time.Hour: "in 23 hours",

	24 * time.Hour: "in a day",
	48 * time.Hour: "in 2 days",
	72 * time.Hour: "in 3 days",

	24 * time.Hour * 6:  "in 6 days",
	24 * time.Hour * 7:  "in a week",
	24 * time.Hour * 8:  "in a week",
	24 * time.Hour * 13: "in a week",

	24 * time.Hour * 7 * 2: "in 2 weeks",
	24 * time.Hour * 7 * 3: "in 3 weeks",
	24 * time.Hour * 7 * 4: "in 4 weeks",

	24 * time.Hour * 7 * 5: "in a month",
	24 * time.Hour * 60:    "in 2 months",
	24 * time.Hour * 89:    "in 2 months",
	24 * time.Hour * 90:    "in 3 months",
	24 * time.Hour * 364:   "in 12 months",

	24 * time.Hour * 365:     "in a year",
	24 * time.Hour * 365 * 2: "in 2 years",
	24 * time.Hour * 365 * 3: "in 3 years",
}

var absoluteDurationsGerman = map[time.Duration]string{
	0: "0 Sekunden",

	1 * time.Second:  "eine Sekunde",
	2 * time.Second:  "2 Sekunden",
	3 * time.Second:  "3 Sekunden",
	59 * time.Second: "59 Sekunden",

	1 * time.Minute:  "eine Minute",
	2 * time.Minute:  "2 Minuten",
	3 * time.Minute:  "3 Minuten",
	59 * time.Minute: "59 Minuten",

	1 * time.Hour:  "eine Stunde",
	2 * time.Hour:  "2 Stunden",
	3 * time.Hour:  "3 Stunden",
	23 * time.Hour: "23 Stunden",

	24 * time.Hour: "ein Tag",
	48 * time.Hour: "2 Tage",
	72 * time.Hour: "3 Tage",

	24 * time.Hour * 6:  "6 Tage",
	24 * time.Hour * 7:  "eine Woche",
	24 * time.Hour * 8:  "eine Woche",
	24 * time.Hour * 13: "eine Woche",

	24 * time.Hour * 7 * 2: "2 Wochen",
	24 * time.Hour * 7 * 3: "3 Wochen",
	24 * time.Hour * 7 * 4: "4 Wochen",

	24 * time.Hour * 7 * 5: "ein Monat",
	24 * time.Hour * 60:    "2 Monate",
	24 * time.Hour * 89:    "2 Monate",
	24 * time.Hour * 90:    "3 Monate",
	24 * time.Hour * 364:   "12 Monate",

	24 * time.Hour * 365:     "ein Jahr",
	24 * time.Hour * 365 * 2: "2 Jahre",
	24 * time.Hour * 365 * 3: "3 Jahre",
}

var relativeDurationsGerman = map[time.Duration]string{
	-1 * time.Second:  "vor einer Sekunde",
	-2 * time.Second:  "vor 2 Sekunden",
	-3 * time.Second:  "vor 3 Sekunden",
	-59 * time.Second: "vor 59 Sekunden",

	-1 * time.Minute:  "vor einer Minute",
	-2 * time.Minute:  "vor 2 Minuten",
	-3 * time.Minute:  "vor 3 Minuten",
	-59 * time.Minute: "vor 59 Minuten",

	-1 * time.Hour:  "vor einer Stunde",
	-2 * time.Hour:  "vor 2 Stunden",
	-3 * time.Hour:  "vor 3 Stunden",
	-23 * time.Hour: "vor 23 Stunden",

	-24 * time.Hour: "vor einem Tag",
	-48 * time.Hour: "vor 2 Tagen",
	-72 * time.Hour: "vor 3 Tagen",

	-24 * time.Hour * 6:  "vor 6 Tagen",
	-24 * time.Hour * 7:  "vor einer Woche",
	-24 * time.Hour * 8:  "vor einer Woche",
	-24 * time.Hour * 13: "vor einer Woche",

	-24 * time.Hour * 7 * 2: "vor 2 Wochen",
	-24 * time.Hour * 7 * 3: "vor 3 Wochen",
	-24 * time.Hour * 7 * 4: "vor 4 Wochen",

	-24 * time.Hour * 7 * 5: "vor einem Monat",
	-24 * time.Hour * 60:    "vor 2 Monaten",
	-24 * time.Hour * 89:    "vor 2 Monaten",
	-24 * time.Hour * 90:    "vor 3 Monaten",
	-24 * time.Hour * 364:   "vor 12 Monaten",

	-24 * time.Hour * 365:     "vor einem Jahr",
	-24 * time.Hour * 365 * 2: "vor 2 Jahren",
	-24 * time.Hour * 365 * 3: "vor 3 Jahren",

	0: "in 0 Sekunden",

	1 * time.Second:  "in einer Sekunde",
	2 * time.Second:  "in 2 Sekunden",
	3 * time.Second:  "in 3 Sekunden",
	59 * time.Second: "in 59 Sekunden",

	1 * time.Minute:  "in einer Minute",
	2 * time.Minute:  "in 2 Minuten",
	3 * time.Minute:  "in 3 Minuten",
	59 * time.Minute: "in 59 Minuten",

	1 * time.Hour:  "in einer Stunde",
	2 * time.Hour:  "in 2 Stunden",
	3 * time.Hour:  "in 3 Stunden",
	23 * time.Hour: "in 23 Stunden",

	24 * time.Hour: "in einem Tag",
	48 * time.Hour: "in 2 Tagen",
	72 * time.Hour: "in 3 Tagen",

	24 * time.Hour * 6:  "in 6 Tagen",
	24 * time.Hour * 7:  "in einer Woche",
	24 * time.Hour * 8:  "in einer Woche",
	24 * time.Hour * 13: "in einer Woche",

	24 * time.Hour * 7 * 2: "in 2 Wochen",
	24 * time.Hour * 7 * 3: "in 3 Wochen",
	24 * time.Hour * 7 * 4: "in 4 Wochen",

	24 * time.Hour * 7 * 5: "in einem Monat",
	24 * time.Hour * 60:    "in 2 Monaten",
	24 * time.Hour * 89:    "in 2 Monaten",
	24 * time.Hour * 90:    "in 3 Monaten",
	24 * time.Hour * 364:   "in 12 Monaten",

	24 * time.Hour * 365:     "in einem Jahr",
	24 * time.Hour * 365 * 2: "in 2 Jahren",
	24 * time.Hour * 365 * 3: "in 3 Jahren",
}
