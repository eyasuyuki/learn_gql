package query

import (
	"github.com/eyasuyuki/learn_gql/database"
	"github.com/eyasuyuki/learn_gql/graph/model"
	"gorm.io/gorm"
)

func Company(db *gorm.DB, id string) (*model.Company, error) {
	idInt, err := database.IdFromBase64(database.COMPANY_PREFIX, id)
	if err != nil {
		return nil, err
	}
	company := database.Company{}
	if err := db.Where("id = ?", idInt).Take(&company).Error; err != nil {
		return nil, err
	}
	return model.NewCompany(&company), nil
}

func Companies(db *gorm.DB, limit int, offset *int) (*model.CompanyPagination, error) {
	var companies []database.Company
	var total int64
	db.Model(&database.Company{}).Count(&total)
	var err error
	if offset != nil {
		err = db.Offset(*offset).Limit(limit).Find(&companies).Error
	} else {
		err = db.Limit(limit).Find(&companies).Error
	}
	if err != nil {
		return nil, err
	}
	pageInfo := &model.PaginationInfo{}
	var nodes []*model.Company
	result := &model.CompanyPagination{PageInfo: pageInfo, Nodes: nodes}
	for _, company := range companies {
		nodes = append(nodes, model.NewCompany(&company))
	}
	pageInfo.Page = 1
	if offset != nil {
		pageInfo.Page = ((int(total) - *offset) / limit) + 1
	}
	pageInfo.Count = len(nodes)
	return result, nil
}

func Department(db *gorm.DB, id string) (*model.Department, error) {
	idInt, err := database.IdFromBase64(database.DEPARTMENT_PREFIX, id)
	if err != nil {
		return nil, err
	}
	department := database.Department{}
	if err := db.Where("id = ?", idInt).Take(&department).Error; err != nil {
		return nil, err
	}
	return model.NewDepartment(&department), nil
}

func Departments(db *gorm.DB, limit int, offset *int) (*model.DepartmentPagination, error) {
	var departments []database.Department
	var total int64
	db.Model(&database.Department{}).Count(&total)
	var err error
	if offset != nil {
		err = db.Offset(*offset).Limit(limit).Find(&departments).Error
	} else {
		err = db.Limit(limit).Find(&departments).Error
	}
	if err != nil {
		return nil, err
	}
	pageInfo := &model.PaginationInfo{}
	var nodes []*model.Department
	result := &model.DepartmentPagination{PageInfo: pageInfo, Nodes: nodes}
	for _, Department := range departments {
		nodes = append(nodes, model.NewDepartment(&Department))
	}
	pageInfo.Page = 1
	if offset != nil {
		pageInfo.Page = ((int(total) - *offset) / limit) + 1
	}
	pageInfo.Count = len(nodes)
	return result, nil
}

func Employee(db *gorm.DB, id string) (*model.Employee, error) {
	idInt, err := database.IdFromBase64(database.EMPLOYEE_PREFIX, id)
	if err != nil {
		return nil, err
	}
	employee := database.Employee{}
	if err := db.Where("id = ?", idInt).Take(&employee).Error; err != nil {
		return nil, err
	}
	return model.NewEmployee(&employee), nil
}

func Employees(db *gorm.DB, limit int, offset *int, email *string, gender *model.Gender, isManager *bool, hasDepartment *bool) (*model.EmployeePagination, error) {
	var employees []database.Employee
	var total int64
	q := db.Limit(limit)
	if offset != nil {
		q = q.Offset(*offset)
	}
	if email != nil {
		q = q.Where("email = ?", email)
	}
	if gender != nil {
		q = q.Where("gender = ?", gender)
	}
	if isManager != nil {
		q = q.Where("is_manager = ?", isManager)
	}
	if hasDepartment != nil {
		if *hasDepartment {
			q = q.Where("department_id is not null")
		} else {
			q = q.Where("department_id is null")
		}
	}
	q.Model(&database.Employee{}).Count(&total)
	if err := q.Find(&employees).Error; err != nil {
		return nil, err
	}
	pageInfo := &model.PaginationInfo{}
	var nodes []*model.Employee
	result := &model.EmployeePagination{PageInfo: pageInfo, Nodes: nodes}
	for _, Employee := range employees {
		nodes = append(nodes, model.NewEmployee(&Employee))
	}
	pageInfo.Page = 1
	if offset != nil {
		pageInfo.Page = ((int(total) - *offset) / limit) + 1
	}
	pageInfo.Count = len(nodes)
	return result, nil
}