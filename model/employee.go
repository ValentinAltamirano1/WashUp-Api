package model

import (
	"gorm.io/gorm"

	"time"
)

type EmployeeClient struct {
    DB *gorm.DB
}

type Employee struct {
    gorm.Model
    Name  string `json:"name"`
	CredentialID uint `json:"credential_id"`
	PhoneNum string `json:"phone_num"`
	AdmissionDate time.Time `json:"admission_date"`
	BirthDate time.Time `json:"birth_date"`
    Email string `json:"email" gorm:"uniqueIndex"`
    Password string 
}

type EmployeeRepository interface {
	SaveEmployee(employee *Employee) error
	EmployeeFirst(query string, args ...interface{}) (*Employee, error)
	DeleteEmployee(employee *Employee) error
}

func (e EmployeeClient) SaveEmployee(employee *Employee) error {
	return e.DB.Save(employee).Error
}

func (e EmployeeClient) EmployeeFirst(query string, args ...interface{}) (*Employee, error) {
	var employee Employee
	if err := e.DB.Where(query, args...).First(&employee).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (e *EmployeeClient) DeleteEmployee(employee *Employee) error {
	return e.DB.Delete(employee).Error
}
