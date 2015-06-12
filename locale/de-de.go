package locale

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/srhnsn/go-humanduration/types"
)

var German = DurationLanguage{
	Day:         UnitConfig{"ein Tag", "einem Tag"},
	DayOther:    UnitConfig{"# Tage", "# Tagen"},
	Hour:        UnitConfig{"eine Stunde", "einer Stunde"},
	HourOther:   UnitConfig{"# Stunden", "# Stunden"},
	Minute:      UnitConfig{"eine Minute", "einer Minute"},
	MinuteOther: UnitConfig{"# Minuten", "# Minuten"},
	Month:       UnitConfig{"ein Monat", "einem Monat"},
	MonthOther:  UnitConfig{"# Monate", "# Monaten"},
	Second:      UnitConfig{"eine Sekunde", "einer Sekunde"},
	SecondOther: UnitConfig{"# Sekunden", "# Sekunden"},
	Week:        UnitConfig{"eine Woche", "einer Woche"},
	WeekOther:   UnitConfig{"# Wochen", "# Wochen"},
	Year:        UnitConfig{"ein Jahr", "einem Jahr"},
	YearOther:   UnitConfig{"# Jahre", "# Jahren"},

	Future: "in #",
	Past:   "vor #",
}

func DeFormatter(amount int, unit types.Unit, future bool, relative bool) string {
	var config UnitConfig

	if unit == types.Second {
		if amount == 1 {
			config = German.Second
		} else {
			config = German.SecondOther
		}
	} else if unit == types.Minute {
		if amount == 1 {
			config = German.Minute
		} else {
			config = German.MinuteOther
		}
	} else if unit == types.Hour {
		if amount == 1 {
			config = German.Hour
		} else {
			config = German.HourOther
		}
	} else if unit == types.Day {
		if amount == 1 {
			config = German.Day
		} else {
			config = German.DayOther
		}
	} else if unit == types.Week {
		if amount == 1 {
			config = German.Week
		} else {
			config = German.WeekOther
		}
	} else if unit == types.Month {
		if amount == 1 {
			config = German.Month
		} else {
			config = German.MonthOther
		}
	} else if unit == types.Year {
		if amount == 1 {
			config = German.Year
		} else {
			config = German.YearOther
		}
	} else {
		panic(fmt.Sprintf("unknown unit: %s", unit))
	}

	var result string

	if relative {
		result = config.Relative
		result = strings.Replace(result, "#", strconv.Itoa(amount), 1)

		if future {
			result = strings.Replace(German.Future, "#", result, 1)
		} else {
			result = strings.Replace(German.Past, "#", result, 1)
		}
	} else {
		result = strings.Replace(config.Absolute, "#", strconv.Itoa(amount), 1)
	}

	return result
}
