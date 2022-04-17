package mutation

import (
	"github.com/eyasuyuki/learn_gql/database"
	"github.com/eyasuyuki/learn_gql/graph/model"
	"gorm.io/gorm"
	"time"
)

// Company

func CreateCompany(db *gorm.DB, input model.CreateCompanyInput) (*model.Company, error) {
	company := &database.Company{CompanyName: input.CompanyName, Representative: input.Representative, PhoneNumber: input.PhoneNumber}
	db.Create(company)
	result := model.NewCompany(company)
	return result, nil
}

func UpdateCompany(db *gorm.DB, input model.UpdateCompanyInput) (*model.Company, error) {
	company := database.NewCompanyUpdate(input.ID, input.CompanyName, input.Representative, input.PhoneNumber)
	db.Save(company)
	result := model.NewCompany(company)
	return result, nil
}

func DeleteCompany(db *gorm.DB, id string) (bool, error) {
	idInt, err := database.IdFromBase64(database.COMPANY_PREFIX, id)
	if err != nil {
		return false, err
	}
	db.Delete(&database.Company{}, idInt)
	return true, nil
}

// Department

func CreateDepartment(db *gorm.DB, input model.CreateDepartmentInput) (*model.Department, error) {
	department := &database.Department{DepartmentName: input.DepartmentName, Email: input.Email}
	db.Create(department)
	return model.NewDepartment(department), nil
}

func UpdateDepartment(db *gorm.DB, input model.UpdateDepartmentInput) (*model.Department, error) {
	department := database.NewDepartmentUpdate(input.ID, input.DepartmentName, input.Email)
	db.Save(department)
	return model.NewDepartment(department), nil
}

func DeleteDepartment(db *gorm.DB, id string) (bool, error) {
	idInt, err := database.IdFromBase64(database.DEPARTMENT_PREFIX, id)
	if err != nil {
		return false, err
	}
	db.Delete(&database.Department{}, idInt)
	return true, nil
}

// Employee

func CreateEmployee(db *gorm.DB, input model.CreateEmployeeInput) (*model.Employee, error) {
	utc,_ := time.LoadLocation("UTC")
	employee := &database.Employee{Name: input.Name, Gender: string(input.Gender), Email: input.Email, LatestLoginAt: time.Now().In(utc), DependentsNum: input.DependentsNum, IsManager: input.IsManager}
	db.Create(employee)
	return model.NewEmployee(employee), nil
}

func UpdateEmployee(db *gorm.DB, input model.UpdateEmployeeInput) (*model.Employee, error) {
	utc, _ := time.LoadLocation("UTC")
	employee := database.NewEmployeeUpdate(input.ID, input.Name, string(input.Gender), input.Email, time.Now().In(utc).Format(database.TIMESTAMP_PATTERN), input.DependentsNum, input.IsManager)
	db.Save(employee)
	return model.NewEmployee(employee), nil
}

func DeleteEmployee(db *gorm.DB, id string) (bool, error) {
	idInt, err := database.IdFromBase64(database.EMPLOYEE_PREFIX, id)
	if err != nil {
		return false, err
	}
	db.Delete(&database.Employee{}, idInt)
	return true, nil
}

