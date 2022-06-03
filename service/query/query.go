package query

import (
	"github.com/eyasuyuki/learn_gql/database"
	"github.com/eyasuyuki/learn_gql/graph/model"
	"gorm.io/gorm"
)

func Company(db *gorm.DB, id string) (*model.Company, error) {
	idInt, err := database.CompanyIDFromBase64(id)
	if err != nil {
		return nil, err
	}
	companies := []database.Company{}
	if err := db.Where("id = ?", idInt).Limit(1).Find(&companies).Error; err != nil {
		return nil, err
	}
	if len(companies) < 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return model.NewCompany(&companies[0]), nil
}

func Companies(db *gorm.DB, limit int, offset *int) (*model.CompanyPagination, error) {
	// totalCount
	var totalCount int64
	db.Model(&database.Company{}).Count(&totalCount)

	// totalPage
	totalPage := (totalCount + int64(limit)) / int64(limit)

	// page
	ofs := 0
	if offset != nil {
		ofs = *offset
	}
	page := (ofs / limit) + 1

	// hasNextPage
	hasNextPage := int64(page) < totalPage
	// hasPreviousPage
	hasPreviousPage := page > 1

	var companies []database.Company
	var err error
	d := db
	if offset != nil {
		d = d.Offset(*offset)
	}
	if err = d.Limit(limit).Find(&companies).Error; err != nil {
		return nil, err
	}

	pageInfo := &model.PaginationInfo{Page: page, PaginationLength: limit, HasNextPage: hasNextPage, HasPreviousPage: hasPreviousPage, Count: len(companies), TotalCount: int(totalCount)}
	result := &model.CompanyPagination{PageInfo: pageInfo, Nodes: []*model.Company{}}
	for _, company := range companies {
		result.Nodes = append(result.Nodes, model.NewCompany(&company))
	}
	return result, nil
}

func Department(db *gorm.DB, id string) (*model.Department, error) {
	idInt, err := database.DepartmentIDFromBase64(id)
	if err != nil {
		return nil, err
	}
	departments := []database.Department{}
	if err := db.Where("id = ?", idInt).Limit(1).Find(&departments).Error; err != nil {
		return nil, err
	}
	if len(departments) < 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return model.NewDepartment(&departments[0]), nil
}

func Departments(db *gorm.DB, limit int, offset *int) (*model.DepartmentPagination, error) {
	// totalCount
	var totalCount int64
	db.Model(&database.Department{}).Count(&totalCount)

	// totalPage
	totalPage := (totalCount / int64(limit)) / int64(limit)

	// page
	ofs := 0
	if offset != nil {
		ofs = *offset
	}
	page := (ofs / limit) + 1

	// hasNextPage
	hasNextPage := int64(page) < totalPage
	// hasPreviousPage
	hasPreviousPage := page > 1

	var departments []database.Department
	var err error
	if offset != nil {
		err = db.Offset(*offset).Limit(limit).Find(&departments).Error
	} else {
		err = db.Limit(limit).Find(&departments).Error
	}
	if err != nil {
		return nil, err
	}
	pageInfo := &model.PaginationInfo{Page: page, PaginationLength: limit, HasNextPage: hasNextPage, HasPreviousPage: hasPreviousPage, Count: len(departments), TotalCount: int(totalCount)}
	result := &model.DepartmentPagination{PageInfo: pageInfo, Nodes: []*model.Department{}}
	for _, Department := range departments {
		result.Nodes = append(result.Nodes, model.NewDepartment(&Department))
	}
	return result, nil
}

func Employee(db *gorm.DB, id string) (*model.Employee, error) {
	idInt, err := database.EmployeeIDFromBase64(id)
	if err != nil {
		return nil, err
	}
	employees := []database.Employee{}
	if err := db.Where("id = ?", idInt).Limit(1).Find(&employees).Error; err != nil {
		return nil, err
	}
	if len(employees) < 1 {
		return nil, gorm.ErrRecordNotFound
	}
	return model.NewEmployee(&employees[0]), nil
}

func Employees(db *gorm.DB, limit int, offset *int, email *string, gender *model.Gender, isManager *bool, hasDepartment *bool) (*model.EmployeePagination, error) {
	// totalCount
	var totalCount int64
	db.Model(&database.Employee{}).Count(&totalCount)

	// totalPage
	totalPage := (totalCount + int64(limit)) / int64(limit)

	// page
	ofs := 0
	if offset != nil {
		ofs = *offset
	}
	page := (ofs / limit) + 1

	// hasNextPage
	hasNextPage := int64(page) < totalPage
	// hasPreviousPage
	hasPreviousPage := page > 1

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

	pageInfo := &model.PaginationInfo{Page: page, PaginationLength: limit, HasNextPage: hasNextPage, HasPreviousPage: hasPreviousPage, Count: len(employees), TotalCount: int(totalCount)}
	result := &model.EmployeePagination{PageInfo: pageInfo, Nodes: []*model.Employee{}}
	for _, Employee := range employees {
		result.Nodes = append(result.Nodes, model.NewEmployee(&Employee))
	}
	return result, nil
}
