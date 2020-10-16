package dateutils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

//GetNow get standard time zone time
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString Get current time
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
