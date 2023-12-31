package date_utils

import "time"

const (
	apiDbLayout   = "2006-01-02 15:04:05"
	apiDateLayout = "2006-01-02T15:04:05.000Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {

	return GetNow().Format(apiDateLayout)
}
func GetNowDbFormat() string {

	return GetNow().Format(apiDbLayout)
}
