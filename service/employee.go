package service

import (
	"github.com/ValentinAltamirano1/WashUp-Api/model"

	"time"
	"errors"
)

type EmployeeParams struct {
    FullName  string `json:"fullname"`
	CredentialID uint `json:"credential_id"`
	Email string `json:"email"`
    Password string 	
	Mobile string `json:"mobile"`
	BirthDate time.Time `json:"birth_date"`
	Gender string `json:"gender"`
}

func CreateEmployee(er model.EmployeeClient, employeeParams EmployeeParams) (*model.Employee, error) {
	collisionEmployee, err := er.EmployeeFirst("email = ?", employeeParams.Email)
	if err != nil {
		return nil, errors.New("error trying to find employee")
	}
	if collisionEmployee != nil {
		return nil, errors.New("employee already exists")
	}

	employee := &model.Employee{
		FullName:     employeeParams.FullName,
		CredentialID: employeeParams.CredentialID,
		Email:    employeeParams.Email,
		Password: hashPassword(employeeParams.Password),
		Mobile: employeeParams.Mobile,
		BirthDate: employeeParams.BirthDate,
		Gender: employeeParams.Gender,
		AdmissionDate: time.Now(),
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
