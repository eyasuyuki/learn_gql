package mutation

import (
	"github.com/eyasuyuki/learn_gql/graph/model"
	"github.com/eyasuyuki/learn_gql/service/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	// set-up
	dsn := "learngql:learngql@tcp(127.0.0.1:3306)/learndb?charset=utf8&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(any(err))
	}
	if db == nil {
		panic(any(err))
	}
	defer func() {
		if db != nil {
			sqlDb, err := db.DB()
			if err != nil {
				panic(any(err))
			}
			sqlDb.Close()
		}
	}()

	// test all
	m.Run()
}

func TestCreateCompany(t *testing.T) {
	input := createCompanyInput()
	company, err := createCompany(input)
	if err != nil {
		panic(any(err))
	}
	if company.ID == "" {
		t.Error("id invalid")
	}
	if company.CompanyName != input.CompanyName {
		t.Error("comanyName invalid")
	}
	if company.Representative != input.Representative {
		t.Error("representative invalid")
	}
	if company.PhoneNumber != input.PhoneNumber {
		t.Error("phoneNumber invalid")
	}
}

func createCompanyInput() model.CreateCompanyInput {
	return model.CreateCompanyInput{CompanyName: "Tower Records", Representative: "嶺脇育夫", PhoneNumber: "03-5555-5555"}
}

func createCompany(input model.CreateCompanyInput) (*model.Company, error) {
	return CreateCompany(db, input)
}

func TestCreateDepartment(t *testing.T) {
	input := createDepartmentInput()
	department, err := createDepartment(input)
	if err != nil {
		panic(any(err))
	}
	if department.ID == "" {
		t.Error("invalid id")
	}
	if department.DepartmentName != input.DepartmentName {
		t.Error("invalid deparmentName")
	}
	if department.Email != input.Email {
		t.Error("email invalid")
	}
}

func createDepartmentInput() model.CreateDepartmentInput {
	return model.CreateDepartmentInput{DepartmentName: "Jazz shop", Email: "jazz@example.com"}
}

func createDepartment(input model.CreateDepartmentInput) (*model.Department, error) {
	return CreateDepartment(db, input)
}

func TestCreateEmployee(t *testing.T) {
	input := createEmployeeInput()
	employee, err := createEmployee(input)
	if err != nil {
		panic(any(err))
	}
	if employee.ID == "" {
		t.Errorf("invalid id")
	}
	if employee.Name != input.Name {
		t.Errorf("invalid name")
	}
	if employee.Gender != input.Gender {
		t.Errorf("invalid gender")
	}
	if employee.Email != input.Email {
		t.Errorf("invalid emil")
	}
	if employee.DependentsNum != input.DependentsNum {
		t.Errorf("invalid dependentsNum")
	}
	if employee.IsManager != input.IsManager {
		t.Errorf("invalid isManager")
	}
}

func createEmployeeInput() model.CreateEmployeeInput {
	return model.CreateEmployeeInput{Name: "Kaede", Gender: model.GenderFemale, Email: "kaede@example.com", DependentsNum: 1, IsManager: true}
}

func createEmployee(input model.CreateEmployeeInput) (*model.Employee, error) {
	return CreateEmployee(db, input)
}

func TestUpdateCompany(t *testing.T) {
	input := createCompanyInput()
	company, err := createCompany(input)
	if err != nil {
		panic(any(err))
	}
	id := company.ID
	updateCompany := model.UpdateCompanyInput{ID: id, CompanyName: "Toyota motors", Representative: "豊田章男", PhoneNumber: "0565-99-9999"}
	company, err = UpdateCompany(db, updateCompany)
	if company.ID != id {
		t.Errorf("id invadlid")
	}
	if company.CompanyName != updateCompany.CompanyName {
		t.Errorf("companyName invalid")
	}
	if company.Representative != updateCompany.Representative {
		t.Errorf("representative invalid")
	}
	if company.PhoneNumber != updateCompany.PhoneNumber {
		t.Errorf("phoneNumber invalid")
	}
}

func TestUpdateDepartment(t *testing.T) {
	input := createDepartmentInput()
	department, err := createDepartment(input)
	if err != nil {
		panic(any(err))
	}
	id := department.ID
	updateDepartment := model.UpdateDepartmentInput{ID: id, DepartmentName: "中央研究所", Email: "labo@example.com"}
	department, err = UpdateDepartment(db, updateDepartment)
	if department.ID != updateDepartment.ID {
		t.Errorf("id invalid")
	}
	if department.DepartmentName != updateDepartment.DepartmentName {
		t.Errorf("departmentName invalid")
	}
	if department.Email != updateDepartment.Email {
		t.Errorf("email invalid")
	}
}

func TestUpdateEmployee(t *testing.T) {
	input := createEmployeeInput()
	employee, err := createEmployee(input)
	if err != nil {
		panic(any(err))
	}
	id := employee.ID
	updateEmployee := model.UpdateEmployeeInput{ID: id, Name: "賀来賢人", Gender: model.GenderMale, Email: "kaku@example.com", DependentsNum: 2, IsManager: false}
	employee, err = UpdateEmployee(db, updateEmployee)
	if err != nil {
		panic(any(err))
	}
	if employee.ID != updateEmployee.ID {
		t.Errorf("id invalid")
	}
	if employee.Gender != updateEmployee.Gender {
		t.Errorf("gender invalid")
	}
	if employee.Name != updateEmployee.Name {
		t.Errorf("name invalid")
	}
	if employee.DependentsNum != updateEmployee.DependentsNum {
		t.Errorf("dependentsNum invalid")
	}
	if employee.IsManager != updateEmployee.IsManager {
		t.Errorf("isManager invalid")
	}
}

func TestDeleteCompany(t *testing.T) {
	companyPagination, err := query.Companies(db, 10000, nil)
	if err != nil {
		panic(any(err))
	}
	for _, node := range companyPagination.Nodes {
		company, err := query.Company(db, node.ID)
		if err != nil {
			panic(any(err))
		}
		if company == nil || company.ID != node.ID {
			t.Errorf("query companies failed")
		}
		result, err := DeleteCompany(db, node.ID)
		if err != nil {
			panic(any(err))
		}
		if !result {
			t.Errorf("delete company failed")
		}
	}
	companyPagination, err = query.Companies(db, 10000, nil)
	if err != nil {
		panic(any(err))
	}
	if companyPagination.PageInfo.Count > 0 {
		t.Errorf("delete company failed")
	}
	if len(companyPagination.Nodes) > 0 {
		t.Errorf("delete company failed")
	}
}

func TestDeleteDepartment(t *testing.T) {
	departmentPagination, err := query.Departments(db, 10000, nil)
	if err != nil {
		panic(any(err))
	}
	for _, node := range departmentPagination.Nodes {
		department, err := query.Department(db, node.ID)
		if err != nil {
			panic(any(err))
		}
		if department == nil || department.ID != node.ID {
			t.Error("query departments failed")
		}
		result, err := DeleteDepartment(db, department.ID)
		if err != nil {
			panic(any(err))
		}
		if !result {
			t.Errorf("delete department failed")
		}
	}
	departmentPagination, err = query.Departments(db, 10000, nil)
	if err != nil {
		panic(any(err))
	}
	if departmentPagination.PageInfo.Count > 0 {
		t.Errorf("delete departments failed")
	}
	if len(departmentPagination.Nodes) > 0 {
		t.Errorf("delete departments failed")
	}

}

func TestDeleteEmployee(t *testing.T) {
	employeePagination, err := query.Employees(db, 10000, nil, nil, nil, nil, nil)
	if err != nil {
		panic(any(err))
	}
	for _, node := range employeePagination.Nodes {
		employee, err := query.Employee(db, node.ID)
		if err != nil {
			panic(any(err))
		}
		if employee == nil || employee.ID != node.ID {
			t.Error("query employees failed")
		}
		result, err := DeleteEmployee(db, employee.ID)
		if err != nil {
			panic(any(err))
		}
		if !result {
			t.Errorf("delete employee failed")
		}
	}
	employeePagination, err = query.Employees(db, 10000, nil, nil, nil, nil, nil)
	if err != nil {
		panic(any(err))
	}
	if employeePagination.PageInfo.Count > 0 {
		t.Error("delete employee failed")
	}
	if len(employeePagination.Nodes) > 0 {
		t.Errorf("delete employee failed")
	}
}