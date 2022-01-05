package utils

import (
	"strconv"
	"strings"
	"time"
)

func parseUnit(unit string) (time.Duration, bool) {
	unit = strings.ToLower(unit)
	if unit == "minutes" || unit == "minute" {
		return time.Minute, true
	}

	if unit == "hour" || unit == "hours" {
		return time.Hour, true
	}

	if unit == "day" || unit == "days" {
		return time.Hour * 24, true
	}

	return time.Nanosecond, false
}

func CalculateExpiresTime(original time.Time, duration string) (time.Time, error) {
	data := strings.Split(duration, " ")
	d, err := strconv.Atoi(data[0])

	if err != nil {
		return time.Time{}, err
	}

	mul, valid := parseUnit(data[1])
	if valid {
		return original.Add(mul * time.Duration(d)), nil
	}

	return time.Time{}, &time.ParseError{}
}

func IsShortenValid(expires time.Time) bool {
	return expires.Before(time.Now())
}
