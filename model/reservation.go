package model

import (

	"gorm.io/gorm"
)

type ReservationClient struct {
    DB *gorm.DB
}

type Reservation struct {
    gorm.Model
    Name  string `json:"name"`
	Date string `json:"date"`
}

type ReservationRepository interface {
    SaveReservation(reservation *Reservation) error
    ReservationFirst(query string, args ...interface{}) (*Reservation, error)
    DeleteReservation(reservation *Reservation) error
}