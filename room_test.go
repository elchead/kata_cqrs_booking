package room

import (
	"testing"
	"time"
)

func TestShowFreeRooms(t *testing.T) {
	rooms := []Room{Room{"Rio"}, Room{"Berlin"}}
	arrival := NewDate(2020, time.August, 13)
	departure := NewDate(2020, time.August, 16)
	availableRooms := getFreeRooms(arrival, departure)

}
