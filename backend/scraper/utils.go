package scraper

import (
	"strings"
	"time"
)

var format = "2006-01-02"

func isToday(date time.Time) bool {
	now := time.Now().Format(format)
	d := date.Format(format)
	return strings.EqualFold(now, d)
}

func formatTime(date time.Time) string {
	return date.Format(format)
}
