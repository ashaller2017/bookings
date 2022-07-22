package dbrepo

import (
	"time"

	"github.com/ashaller2017/bookings/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

//inserts reservation into DB
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {

	return 1, nil
}

//insert a room restriction into DB
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	return nil
}

//returns true if availability exist for roomID and false if not
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {

	return false, nil

}

//search availability for all rooms and returns slice of available rooms if any for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room

	return rooms, nil
}

//gets room by ID
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	return room, nil
}
