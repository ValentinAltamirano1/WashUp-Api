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
    Date      string `json:"date"`
	Time   string `json:"time"`
	Location  string `json:"location"`
    TotalPrice float64 `json:"total_price"`
    EmployeeID *uint
    UserID *uint
    State string `json:"state"`
}

type ReservationRepository interface {
    SaveReservation(reservation *Reservation) error
    ReservationFirst(query string, args ...interface{}) (*Reservation, error)
    DeleteReservation(reservation *Reservation) error
    UpdateReservation(reservation *Reservation, newReservation *Reservation) error
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

func (rc *ReservationClient) GetAllReservationsByService(servicio string) ([]Reservation, error) {
    var reservas []Reservation

    if err := rc.DB.Where("service = ?", servicio).Find(&reservas).Error; err != nil {
        return nil, err
    }

    return reservas, nil
}

func (rr ReservationClient) GetAllReservationsByServiceAndDate(servicio string, fecha string) ([]Reservation, error) {
    var reservas []Reservation
    if err := rr.DB.Where("service = ? AND date = ?", servicio, fecha).Find(&reservas).Error; err != nil {
        return nil, err
    }
    return reservas, nil
}

func (rr ReservationClient) GetAllReservationsWithoutEmployee() ([]Reservation, error) {
    var reservas []Reservation
    if err := rr.DB.Where("employee_id IS NULL").Find(&reservas).Error; err != nil {
        return nil, err
    }
    return reservas, nil
}

func (rr ReservationClient) UpdateReservation(reservation *Reservation, newReservation *Reservation) error {
    return rr.DB.Model(reservation).Updates(newReservation).Error
}

func (rr ReservationClient) GetAllReservationsByEmployee(employeeID uint) ([]Reservation, error) {
    var reservas []Reservation
    if err := rr.DB.Where("employee_id = ? AND state IS NULL", employeeID).Find(&reservas).Error; err != nil {
        return nil, err
    }
    return reservas, nil
}

func (rr ReservationClient) GetAllReservationsDoneByEmployee(employeeID uint) ([]Reservation, error) {
    var reservas []Reservation
    if err := rr.DB.Where("employee_id = ? AND state = ?", employeeID, "done").Find(&reservas).Error; err != nil {
        return nil, err
    }
    return reservas, nil
}

func (rc *ReservationClient) GetAllReservationsByUserID(userID string) ([]Reservation, error) {
    var misReservas []Reservation

    // Realiza una consulta en la base de datos para obtener todas las reservas para el servicio dado.
    if err := rc.DB.Where("userID = ?", userID).Find(&misReservas).Error; err != nil {
        return nil, err
    }

    return misReservas, nil
}
