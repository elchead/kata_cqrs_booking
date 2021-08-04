package room

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShowFreeRooms(t *testing.T) {
	rooms := []Room{"Amsterdam", "Berlin"}
	arrival := NewDate(2020, time.August, 13)
	departure := NewDate(2020, time.August, 16)
	availableRooms := GetFreeRooms(arrival, departure)
	assert.Equal(t, rooms, availableRooms)

}

func TestBookRoom(t *testing.T) {
	arrival := NewDate(2020, time.September, 8)
	departure := NewDate(2020, time.September, 11)
	booking := Booking{id: 0, name: "Berlin", arrival: arrival, departure: departure}

	rooms := []Room{"Rio", "Berlin"}
	hotel := Hotel{rooms}
	err := hotel.BookARoom(booking)
	assert.NoError(t, err)
	availableRooms := hotel.GetFreeRooms(arrival, departure)
	assert.Equal(t, []Room{"Rio"}, availableRooms)
}
