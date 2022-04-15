package database

import (
	"encoding/base64"
	"strconv"
	"time"
)

const (
	COMPANY_PREFIX = "Company:"
	DEPARTMENT_PREFIX = "Department:"
	EMPLOYEE_PREFIX = "Employee:"
)

type Company struct {
	ID             int64 `gorm:"primaryKey"`
	CompanyName    string
	Representative string
	PhoneNumber    string
}

func NewCompanyUpdate(id string, companyName string, representative string, phoneNumber string) *Company {
	idInt, err := IdFromBase64(COMPANY_PREFIX, id)
	if err != nil {
		panic(any(err))
	}
	return &Company{ID: idInt, CompanyName: companyName, Representative: representative, PhoneNumber: phoneNumber}
}

type Department struct {
	ID             int64 `gorm:"primaryKey"`
	DepartmentName string
	Email          string
}

func NewDepartmentUpdate(id string, departmentName string, email string) *Department {
	idInt, err := IdFromBase64(DEPARTMENT_PREFIX, id)
	if err != nil {
		panic(any(err))
	}
	return &Department{ID: idInt, DepartmentName: departmentName, Email: email}
}

type Employee struct {
	ID            int64 `gorm:"primaryKey"`
	Name          string
	Gender        string
	Email         string
	LatestLoginAt time.Time
	//  扶養家族の人数
	DependentsNum int
	//  管理職かどうか
	IsManager  bool
	DepartmentID string
	CompanyID    string
}

func NewEmployeeUpdate(id string, name string, gender string, email string, lastLoginAt time.Time, dependentsNum int, isManager bool) *Employee {
	idInt, err := IdFromBase64(DEPARTMENT_PREFIX, id)
	if err != nil {
		panic(any(err))
	}
	return &Employee{ID: idInt, Name: name, Gender: gender, Email: email, LatestLoginAt: lastLoginAt, DependentsNum: dependentsNum, IsManager: isManager}
}

func IdFromBase64(prefix string, id string) (int64,error) {
	var out []byte
	base64.StdEncoding.Decode(out, []byte(id))
	strId := out[len(prefix):len(out)]
	return strconv.ParseInt(string(strId), 10, 64)
}
