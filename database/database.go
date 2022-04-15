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
	TIMESTAMP_PATTERN = "2006-01-02 15:04:05.999999999"
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
	CompanyID		int64
}

func NewDepartmentUpdate(id string, departmentName string, email string, companyID string) *Department {
	idInt, err := IdFromBase64(DEPARTMENT_PREFIX, id)
	if err != nil {
		panic(any(err))
	}
	companyIDInt, err := IdFromBase64(COMPANY_PREFIX, companyID)
	if err != nil {
		panic(any(err))
	}
	return &Department{ID: idInt, DepartmentName: departmentName, Email: email, CompanyID: companyIDInt}
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
	DepartmentID int64
	CompanyID    int64
}

func NewEmployeeUpdate(id string, name string, gender string, email string, lastLoginAt string, dependentsNum int, isManager bool, departmentID string, companyID string) *Employee {
	idInt, err := IdFromBase64(EMPLOYEE_PREFIX, id)
	if err != nil {
		panic(any(err))
	}
	departmentIDInt, err := IdFromBase64(DEPARTMENT_PREFIX, departmentID)
	if err != nil {
		panic(any(err))
	}
	companyIDInt, err := IdFromBase64(COMPANY_PREFIX, companyID)
	if err != nil {
		panic(any(err))
	}
	lastLoginAtTime, err := time.Parse(TIMESTAMP_PATTERN, lastLoginAt)
	if err != nil {
		panic(any(err))
	}
	return &Employee{ID: idInt, Name: name, Gender: gender, Email: email, LatestLoginAt: lastLoginAtTime, DependentsNum: dependentsNum, IsManager: isManager, DepartmentID: departmentIDInt, CompanyID: companyIDInt}
}

func IdFromBase64(prefix string, id string) (int64,error) {
	out, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		panic(any(err))
	}
	strId := out[len(prefix):len(out)]
	return strconv.ParseInt(string(strId), 10, 64)
}
