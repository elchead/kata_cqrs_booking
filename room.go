package room

import "time"

type Date struct {
	date time.Time
}

func NewDate(year int, month time.Month, day int) Date {
	loc, _ := time.LoadLocation("UTC")
	return Date{date: time.Date(year, month, day, 0, 0, 0, 0, loc)}
}

type Room struct {
	name string
}

func GetFreeRooms(arrival Date, departure Date) []Room {
	return []Room{Room{"Rio"}, Room{"Berlin"}}
}
