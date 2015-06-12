package locale

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/srhnsn/go-humanduration/types"
)

var English = DurationLanguage{
	Day:         UnitConfig{"a day", "a day"},
	DayOther:    UnitConfig{"# days", "# days"},
	Hour:        UnitConfig{"an hour", "an hour"},
	HourOther:   UnitConfig{"# hours", "# hours"},
	Minute:      UnitConfig{"a minute", "a minute"},
	MinuteOther: UnitConfig{"# minutes", "# minutes"},
	Month:       UnitConfig{"a month", "a month"},
	MonthOther:  UnitConfig{"# months", "# months"},
	Second:      UnitConfig{"a second", "a second"},
	SecondOther: UnitConfig{"# seconds", "# seconds"},
	Week:        UnitConfig{"a week", "a week"},
	WeekOther:   UnitConfig{"# weeks", "# weeks"},
	Year:        UnitConfig{"a year", "a year"},
	YearOther:   UnitConfig{"# years", "# years"},

	Future: "in #",
	Past:   "# ago",
}

func EnFormatter(amount int, unit types.Unit, future bool, relative bool) string {
	var config UnitConfig

	if unit == types.Second {
		if amount == 1 {
			config = English.Second
		} else {
			config = English.SecondOther
		}
	} else if unit == types.Minute {
		if amount == 1 {
			config = English.Minute
		} else {
			config = English.MinuteOther
		}
	} else if unit == types.Hour {
		if amount == 1 {
			config = English.Hour
		} else {
			config = English.HourOther
		}
	} else if unit == types.Day {
		if amount == 1 {
			config = English.Day
		} else {
			config = English.DayOther
		}
	} else if unit == types.Week {
		if amount == 1 {
			config = English.Week
		} else {
			config = English.WeekOther
		}
	} else if unit == types.Month {
		if amount == 1 {
			config = English.Month
		} else {
			config = English.MonthOther
		}
	} else if unit == types.Year {
		if amount == 1 {
			config = English.Year
		} else {
			config = English.YearOther
		}
	} else {
		panic(fmt.Sprintf("unknown unit: %s", unit))
	}

	var result string

	if relative {
		result = config.Relative
		result = strings.Replace(result, "#", strconv.Itoa(amount), 1)

		if future {
			result = strings.Replace(English.Future, "#", result, 1)
		} else {
			result = strings.Replace(English.Past, "#", result, 1)
		}
	} else {
		result = strings.Replace(config.Absolute, "#", strconv.Itoa(amount), 1)
	}

	return result
}
