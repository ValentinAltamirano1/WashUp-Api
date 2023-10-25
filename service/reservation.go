package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/ValentinAltamirano1/WashUp-Api/model"
)

func ObtenerFechasDisponibles(rr model.ReservationClient, servicio string) ([]string, error) {
    hoy := time.Now()
    unAnioDespues := hoy.AddDate(1, 0, 0)
    fechasNoDisponibles := []string{}

    for fecha := hoy; fecha.Before(unAnioDespues); fecha = fecha.AddDate(0, 0, 1) {
        fechaStr := fecha.Format("2006-01-02")

        reservasParaFecha, err := rr.GetAllReservationsByServiceAndDate(servicio, fechaStr)
        if err != nil {
            return nil, err // Manejar el error adecuadamente, según tu lógica de negocio.
        }

        if len(reservasParaFecha) == 4 {
            fechasNoDisponibles = append(fechasNoDisponibles, fechaStr)
			fmt.Printf("Fecha no disponible: %s para el servicio %s", fechaStr, servicio)
        }
    }

    return fechasNoDisponibles, nil
}



// ReservationParams contiene los parámetros necesarios para crear una reserva.
type ReservationParams struct {
	Servicio  string `json:"servicio"`
	Fecha     string `json:"fecha"`
	Horario   string `json:"horario"`
	Ubicacion string `json:"ubicacion"`
}

func CreateReservation(rr model.ReservationClient, params ReservationParams) (*model.Reservation, error) {
	// Crear una instancia de model.Reservation con los datos proporcionados en params.
	reserva := &model.Reservation{
		Service:   params.Servicio,
		Date:      params.Fecha,
		Time:      params.Horario,
		Location:  params.Ubicacion,
	}

	// Insertar la reserva en la base de datos.
	if err := rr.SaveReservation(reserva); err != nil {
		return nil, errors.New("error trying to save reservation")
	}

	return reserva, nil
}


// ReservationCheckParams contiene los parámetros necesarios para verificar la disponibilidad de un horario y fecha.
type ReservationCheckParams struct {
	Fecha   string `json:"fecha"`
	Horario string `json:"horario"`
}

// CheckReservation verifica la disponibilidad de un horario y fecha específicos.
func CheckReservation(rr model.ReservationClient, params ReservationCheckParams) (bool, error) {
	// Aquí puedes implementar la lógica para verificar la disponibilidad de un horario y fecha.
	// Consulta la base de datos u otro método para verificar si el horario está ocupado.
	// Retorna true si el horario está disponible, false si está ocupado o un error en caso de problemas.
	return false, errors.New("CheckReservation function not implemented")
}
