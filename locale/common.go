package locale

type DurationLanguage struct {
	Day         UnitConfig
	DayOther    UnitConfig
	Hour        UnitConfig
	HourOther   UnitConfig
	Minute      UnitConfig
	MinuteOther UnitConfig
	Month       UnitConfig
	MonthOther  UnitConfig
	Second      UnitConfig
	SecondOther UnitConfig
	Week        UnitConfig
	WeekOther   UnitConfig
	Year        UnitConfig
	YearOther   UnitConfig

	Future string
	Past   string
}

type UnitConfig struct {
	Absolute string
	Relative string
}
