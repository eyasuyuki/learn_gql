package database

import (
	"encoding/base64"
	"github.com/eyasuyuki/learn_gql/graph/model"
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

func NewCompany(input model.CreateCompanyInput) *Company {
	return &Company{CompanyName: input.CompanyName, Representative: input.Representative, PhoneNumber: input.PhoneNumber}
}

func NewCompanyUpdate(input model.UpdateCompanyInput) *Company {
	idInt, err := IdFromBase64(COMPANY_PREFIX, input.ID)
	if err != nil {
		panic(any(err))
	}
	return &Company{ID: idInt, CompanyName: input.CompanyName, Representative: input.Representative, PhoneNumber: input.PhoneNumber}
}

type Department struct {
	ID             int64 `gorm:"primaryKey"`
	DepartmentName string
	Email          string
}

func NewDepartment(input model.CreateDepartmentInput) *Department {
	return &Department{DepartmentName: input.DepartmentName, Email: input.Email}
}

func NewDepartmentUpdate(input model.UpdateDepartmentInput) *Department {
	idInt, err := IdFromBase64(DEPARTMENT_PREFIX, input.ID)
	if err != nil {
		panic(any(err))
	}
	return &Department{ID: idInt, DepartmentName: input.DepartmentName, Email: input.Email}
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

func NewEmployee(input model.CreateEmployeeInput) *Employee {
	return &Employee{Name: input.Name, Gender: string(input.Gender), Email: input.Email, LatestLoginAt: time.Now(), DependentsNum: input.DependentsNum, IsManager: input.IsManager}
}

func NewEmployeeUpdate(input model.UpdateEmployeeInput) *Employee {
	idInt, err := IdFromBase64(DEPARTMENT_PREFIX, input.ID)
	if err != nil {
		panic(any(err))
	}
	return &Employee{ID: idInt, Name: input.Name, Gender: string(input.Gender), Email: input.Email, LatestLoginAt: time.Now(), DependentsNum: input.DependentsNum, IsManager: input.IsManager}
}

func IdFromBase64(prefix string, id string) (int64,error) {
	var out []byte
	base64.StdEncoding.Decode(out, []byte(id))
	strId := out[len(prefix):len(out)]
	return strconv.ParseInt(string(strId), 10, 64)
}
