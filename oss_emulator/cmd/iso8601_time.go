package emulator

import (
	"strings"
	"time"
)

const (
	iso8601TimeFormat     = "2006-01-02T15:04:05.000Z"
	iso8601TimeFormatLong = "2006-01-02T15:04:05.000000Z"
)

func ISO8601Format(t time.Time) string {
	value := t.Format(iso8601TimeFormat)
	if len(value) < len(iso8601TimeFormat) {
		value = t.Format(iso8601TimeFormat[:len(iso8601TimeFormat)-1])
		return value + strings.Repeat("0", (len(iso8601TimeFormat)-1)-len(value)) + "Z"
	}
	return value
}

func ISO8601Parse(iso8601 string) (t time.Time, err error) {
	for _, layout := range []string{
		iso8601TimeFormat,
		iso8601TimeFormatLong,
		time.RFC3339,
	} {
		t, err = time.Parse(layout, iso8601)
		if err == nil {
			return t, nil
		}
	}

	return t, err
}
