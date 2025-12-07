package core_engine

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ConvertDuration converts a duration string to a duration object
func ConvertDuration(durationStr string) (time.Duration, error) {
	parts := strings.Split(durationStr, " ")
	var seconds int64
	var err error
	if len(parts) == 1 {
		seconds, err = strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return 0, err
		}
	} else {
		matcher, _ := regexp.Compile(`(\d+) days?|(\d+) hours?|(\d+) minutes?|(\d+) seconds?`)
		match := matcher.FindStringSubmatch(durationStr)
		if match == nil {
			return 0, fmt.Errorf("invalid duration string: %s", durationStr)
		}
		seconds = 0
		for i, part := range match {
			switch i {
			case 1:
				seconds += int64(math.Round(float64(part) * 24 * 60 * 60))
			case 3:
				seconds += int64(math.Round(float64(part) * 60 * 60))
			case 5:
				seconds += int64(math.Round(float64(part) * 60))
			case 7:
				seconds += int64(math.Round(float64(part)))
			}
		}
	}
	return time.Duration(seconds), nil
}

// ConvertTime converts a time string to a time.Time object
func ConvertTime(timeStr string, layout string) (time.Time, error) {
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}