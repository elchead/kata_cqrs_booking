package booking

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BookingTests struct {
	suite.Suite
	rooms     []Room
	arrival   Date
	departure Date
	hotel     Hotel
}

func (suite *BookingTests) SetupTest() {
	suite.rooms = []Room{"Rio", "Berlin"}
	suite.arrival = NewDate(2020, time.August, 1)
	suite.departure = NewDate(2020, time.August, 31)
	suite.hotel = NewHotel(suite.rooms)
}

func (suite *BookingTests) TestShowFreeRooms(t *testing.T) {
	availableRooms := suite.hotel.GetFreeRooms(suite.arrival, suite.departure)
	assert.Equal(t, suite.rooms, availableRooms)
}

func (suite *BookingTests) TestBookRoom(t *testing.T) {
	booking := Booking{id: 0, room: "Berlin", arrival: suite.arrival, departure: suite.departure}
	err := suite.hotel.BookARoom(booking)
	assert.NoError(t, err)
	assert.Equal(t, []Booking{booking}, suite.hotel.bookings)
}

func (suite *BookingTests) TestNoDoubleBooking(t *testing.T) {
	booking := Booking{id: 0, room: "Berlin", arrival: suite.arrival, departure: suite.departure}
	err := suite.hotel.BookARoom(booking)
	assert.NoError(t, err)
	err = suite.hotel.BookARoom(booking)
	assert.Error(t, err)
}

func (suite *BookingTests) TestFailBookingOverlappingDays(t *testing.T) {
	booking := Booking{id: 0, room: "Berlin", arrival: suite.arrival, departure: suite.departure}
	err := suite.hotel.BookARoom(booking)
	assert.NoError(t, err)

	arrival2 := suite.arrival.AddDate(0, 0, -5)
	departure2 := suite.departure.AddDate(0, 0, -5)
	assert.Equal(t, []Room{"Rio"}, suite.hotel.GetFreeRooms(arrival2, departure2))

	booking2 := Booking{id: 1, room: "Berlin", arrival: arrival2, departure: departure2}
	err = suite.hotel.BookARoom(booking2)
	assert.Error(t, err)
}

func (setup *BookingTests) TestFailBookingInvalidRoom(t *testing.T) {
	booking := Booking{id: 0, room: "Amsterdam", arrival: setup.arrival, departure: setup.departure}
	err := setup.hotel.BookARoom(booking)
	assert.Error(t, err)
}

func (suite *BookingTests) TestBookAndShowFreeRooms(t *testing.T) {
	booking := Booking{id: 0, room: "Berlin", arrival: suite.arrival, departure: suite.departure}

	err := suite.hotel.BookARoom(booking)
	assert.NoError(t, err)
	availableRooms := suite.hotel.GetFreeRooms(suite.arrival, suite.departure)
	assert.Equal(t, []Room{"Rio"}, availableRooms)
}

func (suite *BookingTests) TestUnaffectedBookingAndFreeRooms(t *testing.T) {
	booking := Booking{id: 0, room: "Berlin", arrival: NewDate(2020, time.January, 1), departure: NewDate(2020, time.May, 13)}

	arrival := NewDate(2020, time.August, 13)
	departure := NewDate(2020, time.August, 16)
	assert.NoError(t, suite.hotel.BookARoom(booking))
	availableRooms := suite.hotel.GetFreeRooms(arrival, departure)
	assert.Equal(t, suite.rooms, availableRooms)
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

func TestIsDateBetween(t *testing.T) {
	start := NewDate(2020, time.January, 2)
	end := NewDate(2020, time.January, 25)
	isBetween := NewDate(2020, time.January, 15).IsBetween(start, end)
	assert.Equal(t, isBetween, true)

	isNotBetween := NewDate(2020, time.January, 1).IsBetween(start, end)
	assert.Equal(t, isNotBetween, false)
	isBetween = NewDate(2020, time.January, 25).IsBetween(start, end)
	assert.Equal(t, isBetween, true)
}
