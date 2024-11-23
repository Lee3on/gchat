package util

import "time"

func FormatTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func ParseTime(str string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", str)
}

func UnixMilliTime(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

func UnunixMilliTime(unix int64) time.Time {
	return time.Unix(0, unix*1000000)
}
