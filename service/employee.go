package service

import (
	"fmt"

	"github.com/ValentinAltamirano1/WashUp-Api/model"

	"errors"
	"time"
)

type EmployeeParams struct {
    FullName  string `json:"fullname"`
	Email string `json:"email"`
    Password string 	
	CredentialID string `json:"credential_id"`
	Mobile string `json:"mobile"`
	BirthDate string `json:"birth_date"`
	Gender string `json:"gender"`
	Department string `json:"department"`
	Adress string `json:"adress"`
}

func CreateEmployee(er model.EmployeeClient, employeeParams EmployeeParams) (*model.Employee, error) {
	collisionEmployee, err := er.EmployeeFirst("email = ?", employeeParams.Email)
	if collisionEmployee != nil {
		return nil, errors.New("employee already exists")
	}

	fmt.Println(collisionEmployee)
	fmt.Println(err)
	fmt.Println(employeeParams)

	employee := &model.Employee{
		FullName: employeeParams.FullName,
		Email: employeeParams.Email,
		Password: hashPassword(employeeParams.Password),
		CredentialID: employeeParams.CredentialID,
		Mobile: employeeParams.Mobile,
		BirthDate: employeeParams.BirthDate,
		Gender: employeeParams.Gender,
		Department: employeeParams.Department,
		Adress: employeeParams.Adress,
		AdmissionDate: time.Now(),
	}

	fmt.Println(employee)
	err = er.SaveEmployee(employee)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error trying to save employee")
	}

	return employee, nil
}

type LoginEmployeeResponse struct {
	Email string `json:"email"`
	FullName string `json:"fullname"`
	Token string `json:"token"`
}

func LoginEmployee(er model.EmployeeClient, email string, password string) (*LoginEmployeeResponse, error) {
	employee, err := er.EmployeeFirst("email = ?", email)
	if err != nil {
		return nil, errors.New("error trying to find employee")
	}

	err = ValidatePassword(employee.Password, password)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := GenerateToken(employee.Email)
	if err != nil {
		return nil, errors.New("error trying to generate token")
	}

	return &LoginEmployeeResponse{
		Email:    employee.Email,
		FullName: employee.FullName,
		Token: token,
	}, nil
}
