package model

import (

	"gorm.io/gorm"
)

type ReservationClient struct {
    DB *gorm.DB
}

type Reservation struct {
    gorm.Model
	Service   string `json:"service"`
    Date      string `json:"date" gorm:"type:date"`
	Time   string `json:"time"`
	Location  string `json:"location"`
    EmployeeID uint
    UserID uint
}

type ReservationRepository interface {
    SaveReservation(reservation *Reservation) error
    ReservationFirst(query string, args ...interface{}) (*Reservation, error)
    DeleteReservation(reservation *Reservation) error
}

func (r ReservationClient) SaveReservation(reservation *Reservation) error {
	return r.DB.Save(reservation).Error
}

func (r ReservationClient) ReservationFirst(query string, args ...interface{}) (*Reservation, error) {
	var reservation Reservation
	if err := r.DB.Where(query, args...).First(&reservation).Error; err != nil {
		return nil, err
	}
	return &reservation, nil
}


func (r *ReservationClient) DeleteReservation(reservation *Reservation) error {
	return r.DB.Delete(reservation).Error
}

// GetAllReservationsByService obtiene todas las reservas para un servicio específico.
func (rc *ReservationClient) GetAllReservationsByService(servicio string) ([]Reservation, error) {
    var reservas []Reservation

    // Realiza una consulta en la base de datos para obtener todas las reservas para el servicio dado.
    if err := rc.DB.Where("service = ?", servicio).Find(&reservas).Error; err != nil {
        return nil, err
    }

    return reservas, nil
}

// GetAllReservationsByServiceAndDate recupera todas las reservas para un servicio y fecha específicos.
func (rr ReservationClient) GetAllReservationsByServiceAndDate(servicio string, fecha string) ([]Reservation, error) {
    var reservas []Reservation
    if err := rr.DB.Where("service = ? AND date = ?", servicio, fecha).Find(&reservas).Error; err != nil {
        return nil, err
    }
    return reservas, nil
}
