package room

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShowFreeRooms(t *testing.T) {
	rooms := []Room{Room{"Rio"}, Room{"Berlin"}}
	arrival := NewDate(2020, time.August, 13)
	departure := NewDate(2020, time.August, 16)
	availableRooms := GetFreeRooms(arrival, departure)
	assert.Equal(t, rooms, availableRooms)

}

// func TestShowNoRooms(t *testing.T) {
// 	rooms := []Room{Room{"Rio"}, Room{"Berlin"}}
// 	arrival := NewDate(2020, time.August, 13)
// 	departure := NewDate(2020, time.August, 16)
// 	availableRooms := GetFreeRooms(arrival, departure)
// 	assert.Equal(t, rooms, availableRooms)

// }
