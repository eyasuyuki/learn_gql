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
	idInt, err := database.CompanyIDFromBase64(id)
	if err != nil {
		return false, err
	}
	if err = db.Where("id = ?", idInt).Delete(&database.Company{}).Error; err != nil {
		return false, err
	}
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
	idInt, err := database.DepartmentIDFromBase64(id)
	if err != nil {
		return false, err
	}
	if err = db.Delete(&database.Department{}, idInt).Error; err != nil {
		return false, err
	}
	return true, nil
}

func SetCompanyToDepartment(db *gorm.DB, id string, companyID string) (*model.Department, error) {
	idInt, err := database.DepartmentIDFromBase64(id)
	if err != nil {
		return nil, err
	}
	companyIdInt, err := database.CompanyIDFromBase64(companyID)
	if err != nil {
		return nil, err
	}
	var department database.Department
	if err = db.Where("id = ?", idInt).Find(&department).Error; err != nil {
		return nil, err
	}
	department.CompanyID = companyIdInt
	if err = db.Save(&department).Error; err != nil {
		return nil, err
	}
	return model.NewDepartment(&department), nil
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
	idInt, err := database.EmployeeIDFromBase64(id)
	if err != nil {
		return false, err
	}
	if err = db.Delete(&database.Employee{}, idInt).Error; err != nil {
		return false, err
	}
	return true, nil
}

func SetDepartmentToEmployee(db *gorm.DB, id string, departmentId string) (*model.Employee, error) {
	idInt, err := database.EmployeeIDFromBase64(id)
	if err != nil {
		return nil, err
	}
	departmentIdInt, err := database.DepartmentIDFromBase64(departmentId)
	if err != nil {
		return nil, err
	}
	var employee database.Employee
	if err = db.Where("id = ?", idInt).Find(&employee).Error; err != nil {
		return nil, err
	}
	employee.DepartmentID = departmentIdInt
	if err = db.Save(&employee).Error; err != nil {
		return nil, err
	}
	return model.NewEmployee(&employee), nil
}

func SetCompanyToEmployee(db *gorm.DB, id string, companyID string) (*model.Employee, error) {
	idInt, err := database.EmployeeIDFromBase64(id)
	if err != nil {
		return nil, err
	}
	companyIdInt, err := database.CompanyIDFromBase64(companyID)
	if err != nil {
		return nil, err
	}
	var employee database.Employee
	if err = db.Where("id = ?", idInt).Find(&employee).Error; err != nil {
		return nil, err
	}
	employee.CompanyID = companyIdInt
	if err = db.Save(&employee).Error; err != nil {
		return nil, err
	}
	return model.NewEmployee(&employee), nil
}

