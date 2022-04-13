package model

import (
	"encoding/base64"
	"github.com/eyasuyuki/learn_gql/database"
	"strconv"
)

type Company struct {
	ID             string                `json:"id"`
	CompanyName    string                `json:"companyName"`
	Representative string                `json:"representative"`
	PhoneNumber    string                `json:"phoneNumber"`
}

func (Company) IsNode() {}

func NewCompany(dto *database.Company) *Company {
	id := idToBase64(database.COMPANY_PREFIX, dto.ID)
	return &Company{ID: id, CompanyName: dto.CompanyName, PhoneNumber: dto.PhoneNumber}
}

type Department struct {
	ID             string              `json:"id"`
	DepartmentName string              `json:"departmentName"`
	Email          string              `json:"email"`
	CompanyID	   string			   `json:"company"`
}

func (Department) IsNode() {}

func NewDepartment(dto *database.Department) *Department {
	id := idToBase64(database.DEPARTMENT_PREFIX, dto.ID)
	return &Department{ID: id, DepartmentName: dto.DepartmentName, Email: dto.Email}
}


type Employee struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Gender        Gender `json:"gender"`
	Email         string `json:"email"`
	LatestLoginAt string `json:"latestLoginAt"`
	//  扶養家族の人数
	DependentsNum int `json:"dependentsNum"`
	//  管理職かどうか
	IsManager  bool        `json:"isManager"`
	DepartmentID string `json:"department"`
	CompanyID    string    `json:"company"`
}

func (Employee) IsNode() {}

func NewEmployee(dto *database.Employee) *Employee {
	id := idToBase64(database.EMPLOYEE_PREFIX, dto.ID)
	latestLoginAt := dto.LatestLoginAt.Format("2006-01-02 15:04:05.999999999")
	return &Employee{ID: id, Name: dto.Name, Gender: Gender(dto.Gender), Email: dto.Email, LatestLoginAt: latestLoginAt, DependentsNum: dto.DependentsNum, IsManager: dto.IsManager, DepartmentID: dto.DepartmentID, CompanyID: dto.CompanyID}
}


func idToBase64(prefix string, id int64) string {
	strId := prefix+strconv.FormatInt(id, 10);
	return base64.StdEncoding.EncodeToString([]byte(strId))
}