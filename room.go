package room

import (
	"errors"
	"time"
)

type Date struct {
	time time.Time
}

func (d Date) Before(other Date) bool {
	return d.time.Before(other.time)
}

func (d Date) After(other Date) bool {
	return d.time.After(other.time)
}

func (d Date) Equal(other Date) bool {
	return d.time.Equal(other.time)
}

func NewDate(year int, month time.Month, day int) Date {
	loc, _ := time.LoadLocation("UTC")
	return Date{time: time.Date(year, month, day, 0, 0, 0, 0, loc)}
}

type Room string

func GetFreeRooms(arrival Date, departure Date) []Room {
	return []Room{"Rio", "Berlin"}
}

type Booking struct {
	id        int
	room      Room
	arrival   Date
	departure Date
}
type Hotel struct {
	rooms    []Room
	bookings []Booking
}

func NewHotel(rooms []Room) Hotel {
	return Hotel{rooms: rooms}
}

func (h *Hotel) BookARoom(booking Booking) error {
	availableRooms := h.GetFreeRooms(booking.arrival, booking.departure)
	for _, room := range availableRooms {
		if room == booking.room {
			h.bookings = append(h.bookings, booking)
			return nil
		}
	}
	return errors.New("Room not found")
}

func (h Hotel) getBookedRooms(arrival Date, departure Date) []Room {
	bookedRooms := make([]Room, 0)
	for _, booking := range h.bookings {
		if (booking.arrival.Before(arrival) || booking.arrival.Equal(arrival)) && booking.departure.After(arrival) {
			bookedRooms = append(bookedRooms, booking.room)
		}
	}
	return bookedRooms
}

func (h Hotel) GetFreeRooms(arrival Date, departure Date) []Room {
	// check if existing booking affects availability
	// if yes, remove room
	bookedRooms := h.getBookedRooms(arrival, departure)
	availableRooms := make([]Room, 0)
	for _, room := range h.rooms {
		// check if inside of bookedRooms
		booked := false
		for _, bookedRoom := range bookedRooms {
			if room == bookedRoom {
				booked = true
				break
			}
		}
		if !booked {
			availableRooms = append(availableRooms, room)
		}
	}
	return availableRooms
}
