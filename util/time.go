package util

import "time"

func FormatTime(t time.Time) string {
	if t.IsZero() {
		return "N/A"
	}
	return t.Format("Jan 2, 2006 15:04 MST")
}

func FormatLocalTime(t time.Time) string {
	if t.IsZero() {
		return "N/A"
	}
	return FormatTime(t.Local())
}
