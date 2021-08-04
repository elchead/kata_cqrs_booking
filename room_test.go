package booking

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShowFreeRooms(t *testing.T) {
	rooms := []Room{"Rio", "Berlin"}
	arrival := NewDate(2020, time.August, 13)
	departure := NewDate(2020, time.August, 16)
	hotel := NewHotel(rooms)
	availableRooms := hotel.GetFreeRooms(arrival, departure)
	assert.Equal(t, rooms, availableRooms)
}

func TestBookRoom(t *testing.T) {
	arrival := NewDate(2020, time.September, 8)
	departure := NewDate(2020, time.September, 11)
	booking := Booking{id: 0, room: "Berlin", arrival: arrival, departure: departure}

	rooms := []Room{"Rio", "Berlin"}
	hotel := NewHotel(rooms)
	err := hotel.BookARoom(booking)
	assert.NoError(t, err)
	assert.Equal(t, []Booking{booking}, hotel.bookings)
}

func TestNoDoubleBooking(t *testing.T) {
	arrival := NewDate(2020, time.September, 8)
	departure := NewDate(2020, time.September, 11)
	booking := Booking{id: 0, room: "Berlin", arrival: arrival, departure: departure}

	rooms := []Room{"Rio", "Berlin"}
	hotel := NewHotel(rooms)
	err := hotel.BookARoom(booking)
	assert.NoError(t, err)
	err = hotel.BookARoom(booking)
	assert.Error(t, err)
}

func TestFailBookingOverlappingDays(t *testing.T) {
	arrival := NewDate(2020, time.September, 1)
	departure := NewDate(2020, time.September, 31)
	booking := Booking{id: 0, room: "Berlin", arrival: arrival, departure: departure}

	rooms := []Room{"Rio", "Berlin"}
	hotel := NewHotel(rooms)
	err := hotel.BookARoom(booking)
	assert.NoError(t, err)
	assert.Equal(t, []Room{"Rio"}, hotel.GetFreeRooms(arrival.AddDate(0, 0, -5), departure.AddDate(0, 0, -5)))
	booking2 := Booking{id: 1, room: "Berlin", arrival: arrival.AddDate(0, 0, -5), departure: departure.AddDate(0, 0, -5)}
	err = hotel.BookARoom(booking2)
	assert.Error(t, err)
}

func TestFailBookingInvalidRoom(t *testing.T) {
	arrival := NewDate(2020, time.September, 8)
	departure := NewDate(2020, time.September, 11)
	booking := Booking{id: 0, room: "Amsterdam", arrival: arrival, departure: departure}
	rooms := []Room{"Rio", "Berlin"}
	hotel := NewHotel(rooms)
	err := hotel.BookARoom(booking)
	assert.Error(t, err)
}

func TestBookAndShow(t *testing.T) {
	arrival := NewDate(2020, time.September, 8)
	departure := NewDate(2020, time.September, 11)
	booking := Booking{id: 0, room: "Berlin", arrival: arrival, departure: departure}

	rooms := []Room{"Rio", "Berlin"}
	hotel := NewHotel(rooms)
	err := hotel.BookARoom(booking)
	assert.NoError(t, err)
	availableRooms := hotel.GetFreeRooms(arrival, departure)
	assert.Equal(t, []Room{"Rio"}, availableRooms)
}

func TestUnaffectedBookingAndFreeRooms(t *testing.T) {
	booking := Booking{id: 0, room: "Berlin", arrival: NewDate(2020, time.January, 1), departure: NewDate(2020, time.May, 13)}

	rooms := []Room{"Rio", "Berlin"}
	arrival := NewDate(2020, time.August, 13)
	departure := NewDate(2020, time.August, 16)
	hotel := NewHotel(rooms)
	assert.NoError(t, hotel.BookARoom(booking))
	availableRooms := hotel.GetFreeRooms(arrival, departure)
	assert.Equal(t, rooms, availableRooms)
}

func TestGetBookedRoom(t *testing.T) {
	arrival := NewDate(2020, time.September, 8)
	departure := NewDate(2020, time.September, 11)
	booking := Booking{id: 0, room: "Berlin", arrival: arrival, departure: departure}

	rooms := []Room{"Rio", "Berlin"}
	hotel := NewHotel(rooms)
	err := hotel.BookARoom(booking)
	assert.NoError(t, err)
	assert.Equal(t, []Room{"Berlin"}, hotel.getBookedRooms(arrival, departure))

}
