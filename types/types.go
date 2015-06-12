package types

import (
	"math"
	"time"
)

type HumanDurationFormatter struct {
	Formatter FormatterFunction
}

type FormatterFunction func(amount int, unit Unit, future bool, relative bool) string

type Unit int

const (
	Day Unit = iota
	Hour
	Minute
	Month
	Second
	Week
	Year
)

func (hdf *HumanDurationFormatter) HumanizeAbsolute(duration time.Duration) string {
	return hdf.Humanize(duration, false)
}

func (hdf *HumanDurationFormatter) HumanizeRelative(duration time.Duration) string {
	return hdf.Humanize(duration, true)
}

func (hdf *HumanDurationFormatter) Humanize(duration time.Duration, relative bool) string {
	future := duration >= 0

	seconds := Round(duration.Seconds())
	secondsAbs := Abs(seconds)

	if secondsAbs < 60 {
		return hdf.Formatter(secondsAbs, Second, future, relative)
	} else if secondsAbs < 60*60 {
		return hdf.Formatter(secondsAbs/60, Minute, future, relative)
	} else if secondsAbs < 60*60*24 {
		return hdf.Formatter(secondsAbs/60/60, Hour, future, relative)
	} else if secondsAbs < 60*60*24*7 {
		return hdf.Formatter(secondsAbs/60/60/24, Day, future, relative)
	} else if secondsAbs < 60*60*24*30 {
		return hdf.Formatter(secondsAbs/60/60/24/7, Week, future, relative)
	} else if secondsAbs < 60*60*24*365 {
		return hdf.Formatter(secondsAbs/60/60/24/30, Month, future, relative)
	}

	return hdf.Formatter(secondsAbs/60/60/24/365, Year, future, relative)
}

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func Round(f float64) int {
	return int(math.Floor(f + .5))
}
