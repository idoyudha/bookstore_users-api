package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC() // take localtime value
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout) // convert to given format
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout) // convert to given format
}
