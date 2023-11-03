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


func ObtenerHorariosDisponibles(rc model.ReservationClient, servicio string, fecha string) ([]string, error) {
    // Horarios predeterminados.
    horariosPredeterminados := []string{
        "9:00 AM", "10:00 AM", "11:00 AM", "2:00 PM", "3:00 PM",
    }

    // Obtener reservas para el servicio y la fecha específicos.
    reservas, err := rc.GetAllReservationsByServiceAndDate(servicio, fecha)
    if err != nil {
        return nil, err
    }

    // Crear un mapa para almacenar los horarios reservados.
    horariosOcupados := make(map[string]struct{})

    // Iterar sobre las reservas y registrar los horarios reservados.
    for _, reserva := range reservas {
        horariosOcupados[reserva.Time] = struct{}{}
    }

    // Crear un slice para almacenar los horarios disponibles.
    horariosDisponibles := []string{}

    // Iterar sobre los horarios predeterminados y agregar solo los no ocupados.
    for _, horario := range horariosPredeterminados {
        // Comprobar si el horario actual no está ocupado y, si no lo está, agregarlo a horariosDisponibles.
        if _, ocupado := horariosOcupados[horario]; !ocupado {
            horariosDisponibles = append(horariosDisponibles, horario)
        }
    }

    return horariosDisponibles, nil
}






// ReservationParams contiene los parámetros necesarios para crear una reserva.
type ReservationParams struct {
	Servicio  string `json:"servicio"`
	Fecha     string `json:"fecha"`
	Horario   string `json:"horario"`
	Ubicacion string `json:"ubicacion"`
	UserEmail string `json:"user_email"`
}

func CreateReservation(rr model.ReservationClient, ur model.UserClient, params ReservationParams) (*model.Reservation, error) {
	user, err := ur.UserFirst("email = ?", params.UserEmail)
	if err != nil {
		return nil, errors.New("error trying to find user")
	}

	reserva := &model.Reservation{
		Service:   params.Servicio,
		Date:      params.Fecha,
		Time:      params.Horario,
		Location:  params.Ubicacion,
		UserID:   &user.ID,
	}

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
