package room

import "time"

type Date struct {
	date time.Time
}

func NewDate(year int, month time.Month, day int) Date {
	loc, _ := time.LoadLocation("UTC")
	return Date{date: time.Date(year, month, day, 0, 0, 0, 0, loc)}
}

type Room string

func GetFreeRooms(arrival Date, departure Date) []Room {
	return []Room{"Rio", "Berlin"}
}

type Booking struct {
	id        int
	name      string
	arrival   Date
	departure Date
}
type Hotel struct {
	rooms []Room
}

func (h Hotel) BookARoom(booking Booking) error {
	return nil
}

func (h Hotel) GetFreeRooms(arrival Date, departure Date) []Room {
	return []Room{"Rio"}
}
