package model

import (
	"github.com/eyasuyuki/learn_gql/database"
	"testing"
	"time"
)

func TestNewCompany(t *testing.T) {
	company := database.Company{ID: 1, CompanyName: "HMV", Representative: "rep1", PhoneNumber: "03-5784-1390"}
	companyModel := NewCompany(&company)
	if companyModel.ID == "" {
		t.Errorf("id invadid")
	}
	if companyModel.CompanyName != company.CompanyName {
		t.Errorf("companyName invalid")
	}
	if companyModel.Representative != company.Representative {
		t.Errorf("representative invalid")
	}
	if companyModel.PhoneNumber != company.PhoneNumber {
		t.Errorf("phoneNumber invalid")
	}
	company2 := database.NewCompanyUpdate(companyModel.ID, companyModel.CompanyName, companyModel.Representative, companyModel.PhoneNumber)
	if company.ID != company2.ID {
		t.Errorf("id2 invalid")
	}
	if company.CompanyName != company2.CompanyName {
		t.Errorf("companyName2 invalid")
	}
	if company.Representative != company2.Representative {
		t.Errorf("representative2 invalid")
	}
	if company.PhoneNumber != company2.PhoneNumber {
		t.Errorf("phoneNumber2 invalid")
	}
}

func TestNewDepartment(t *testing.T) {
	department := database.Department{ID: 2, DepartmentName: "Shibuya", Email: "email@example.com", CompanyID: 1}
	departmentModel := NewDepartment(&department)
	if departmentModel.ID == "" {
		t.Errorf("id invalid")
	}
	if departmentModel.DepartmentName != department.DepartmentName {
		t.Errorf("departmentName invalid")
	}
	if departmentModel.Email != department.Email {
		t.Errorf("email invalid")
	}
	department2 := database.NewDepartmentUpdate(departmentModel.ID, departmentModel.DepartmentName, departmentModel.Email)
	if department.ID != department2.ID {
		t.Errorf("id2 invalid")
	}
	if department.DepartmentName != department2.DepartmentName {
		t.Errorf("departmentName2 invalid")
	}
	if department.Email != department2.Email {
		t.Errorf("email2 invalid")
	}
}


func TestNewEmployee(t *testing.T) {
	utc, _ := time.LoadLocation("UTC")
	employee := database.Employee{ID: 3, Name: "玉井詩織", Gender: string(GenderFemale), Email: "nobody@example.com", LatestLoginAt: time.Now().In(utc), DependentsNum: 4, IsManager: true, DepartmentID: 2, CompanyID: 1}
	employeeModel := NewEmployee(&employee)
	if employeeModel.ID == "" {
		t.Errorf("id invalid")
	}
	if employeeModel.Name != employee.Name {
		t.Errorf("name invalid")
	}
	if string(employeeModel.Gender) != employee.Gender {
		t.Errorf("gender invalid")
	}
	if employeeModel.Email != employee.Email {
		t.Errorf("email invalid")
	}
	// TODO latestLoginAt
	if employeeModel.DependentsNum != employee.DependentsNum {
		t.Error("dependentsNum invalid")
	}
	if employeeModel.IsManager != employee.IsManager {
		t.Errorf("isManager invalid")
	}
	employee2 := database.NewEmployeeUpdate(employeeModel.ID, employeeModel.Name, string(employeeModel.Gender), employeeModel.Email, employeeModel.LatestLoginAt, employeeModel.DependentsNum, employeeModel.IsManager)
	if employee.ID != employee2.ID {
		t.Errorf("id2 invalid")
	}
	if employee.Name != employee2.Name {
		t.Errorf("name2 invadid")
	}
	if employee.Gender != employee2.Gender {
		t.Errorf("gender2 invalid")
	}
	if employee.Email != employee2.Email {
		t.Errorf("email2 invalid")
	}
	if employee.LatestLoginAt != employee2.LatestLoginAt {
		t.Errorf("latestLoginAt2 invalid")
	}
	if employee.DependentsNum != employee2.DependentsNum {
		t.Errorf("dependentsNum2 invalid")
	}
	if employee.IsManager != employee2.IsManager {
		t.Errorf("isManage2 invalid")
	}

}