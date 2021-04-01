package date_utils

import "time"

const (
	apiDateLayout = "02-01-2006T15:04:05Z"
	apiDbLayout   = "2006-01-01 15:04:05"
)

//here we are saying take local time and display in given format
//for time format just chnage the placeholder and can display time accordigly
//for change in zone just do -digit hrs like (02-01-2006T15:04:00-07:00)
//best practice is to work on standard time zone
//when microservices working on different time zone then change need to do is now := time.Now().UTC()
//any change in date has to make will be done here, & it will be common package for entire microservice
func GetNow() time.Time {
	return time.Now()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
