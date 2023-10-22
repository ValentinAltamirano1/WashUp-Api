package service

import (
	"github.com/ValentinAltamirano1/WashUp-Api/model"

	"time"
	"errors"
)

type EmployeeParams struct {
	Name	 string `json:"name"`
	CredentialID uint `json:"credential_id"`
	PhoneNum string `json:"phone_num"`
	AdmissionDate time.Time `json:"admission_date"`
	BirthDate time.Time `json:"birth_date"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateEmployee(er model.EmployeeClient, employeeParams EmployeeParams) (*model.Employee, error) {
	collisionEmployee, err := er.EmployeeFirst("email = ?", employeeParams.Email)
	if collisionEmployee != nil {
		return nil, errors.New("employee already exists")
	}

	employee := &model.Employee{
		Name:     employeeParams.Name,
		CredentialID: employeeParams.CredentialID,
		PhoneNum: employeeParams.PhoneNum,
		AdmissionDate: employeeParams.AdmissionDate,
		BirthDate: employeeParams.BirthDate,
		Email:    employeeParams.Email,
		Password: hashPassword(employeeParams.Password),
	}

	err = er.SaveEmployee(employee)
	if err != nil {
		return nil, errors.New("error trying to save employee")
	}

	return employee, nil
}

func LoginEmployee(er model.EmployeeClient, email string, password string) (*model.Employee, error) {
	employee, err := er.EmployeeFirst("email = ?", email)
	if err != nil {
		return nil, errors.New("error trying to find employee")
	}

	err = ValidatePassword(employee.Password, password)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return employee, nil
}
