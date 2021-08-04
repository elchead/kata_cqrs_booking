package room

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShowFreeRooms(t *testing.T) {
	rooms := []Room{"Rio", "Berlin"}
	arrival := NewDate(2020, time.August, 13)
	departure := NewDate(2020, time.August, 16)
	availableRooms := GetFreeRooms(arrival, departure)
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
	// 	availableRooms := hotel.GetFreeRooms(arrival, departure)
	// 	assert.Equal(t, []Room{"Rio"}, availableRooms)
}

func TestFailBookRoom(t *testing.T) {
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
