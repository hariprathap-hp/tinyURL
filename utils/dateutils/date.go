package dateutils

import "time"

const (
	apiDateLayout = "2006-03-25T07:50:00Z"
)

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetExpiry() time.Time {
	return time.Now().UTC().AddDate(2, 0, 0)
}
