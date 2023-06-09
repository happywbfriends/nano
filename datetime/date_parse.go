package datetime

import (
	"errors"
	"strings"
	"time"
)

const (
	lenDt          = len("2006-01-02")
	lenDtTmNoSec   = len("2006-01-02T15:04")
	lenDtTmWithSec = len("2006-01-02T15:04:05")
)

var errInvalidTime = errors.New("invalid time format")

/*
	Fast Datetime parser according to following formats:
	1. 2006-01-02					=>	2006-01-02T00:00:00 UTC
	2. 2006-01-02T15:04				=>	2006-01-02T15:04:00 UTC
	3. 2006-01-02T15:04:05			=>	2006-01-02T15:04:05 UTC
	4. 2006-01-02T15:04:05TZ		=>	2006-01-02T15:04:05 TZ
	5. 2006-01-02T15:04:05.999		=>	2006-01-02T15:04:05.999 UTC
	6. 2006-01-02T15:04:05.999TZ	=>	2006-01-02T15:04:05.999 TZ

	Always returns time in UTC regardless of what TZ was provided.
*/
func ParseDateTime(date string) (t time.Time, e error) {
	// In the absence of a time zone indicator, time.Parse returns a time in UTC.
	strLen := len(date)
	switch strLen {
	case lenDt:
		return time.Parse("2006-01-02", date)
	case lenDtTmNoSec:
		return time.Parse("2006-01-02T15:04", date)
	case lenDtTmWithSec:
		return time.Parse("2006-01-02T15:04:05", date)
	}

	if strLen <= lenDtTmWithSec {
		return time.Time{}, errInvalidTime
	}

	if date[lenDtTmWithSec] == '.' { // cases 5, 6
		if strings.ContainsAny(date[lenDtTmWithSec+1:], "Z+-") {
			if t, e = time.Parse("2006-01-02T15:04:05.9Z07:00", date); e != nil {
				return time.Time{}, e
			} else {
				return t.UTC(), nil
			}

		} else {
			return time.Parse("2006-01-02T15:04:05.9", date)
		}
	} else { // case 4
		if t, e = time.Parse("2006-01-02T15:04:05Z07:00", date); e != nil {
			return time.Time{}, e
		} else {
			return t.UTC(), nil
		}
	}
}
