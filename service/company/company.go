package company

import (
	"github.com/eyasuyuki/learn_gql/database"
	"github.com/eyasuyuki/learn_gql/graph/model"
	"gorm.io/gorm"
)

func Departments(db *gorm.DB, obj *model.Company) (*model.DepartmentPagination, error) {
	idInt, err := database.IdFromBase64(database.COMPANY_PREFIX, obj.ID)
	if err != nil {
		return nil, err
	}
	var departments []database.Department
	if err = db.Find(&departments).Where("company_id = ? ", idInt).Error; err != nil {
		return nil, err
	}
	var nodes []*model.Department
	for _, d := range departments {
		nodes = append(nodes, model.NewDepartment(&d))
	}
	var pageInfo model.PaginationInfo
	pageInfo.Page = 1
	pageInfo.PaginationLength = 0
	pageInfo.HasPreviousPage = false
	pageInfo.HasNextPage = false
	pageInfo.Count = len(nodes)
	pageInfo.TotalCount = len(nodes)
	var result model.DepartmentPagination
	result.PageInfo = &pageInfo
	result.Nodes = nodes
	return &result, nil
}

func Employees(db *gorm.DB, obj *model.Company) (*model.EmployeePagination, error) {
	idInt, err := database.IdFromBase64(database.COMPANY_PREFIX, obj.ID)
	if err != nil {
		return nil, err
	}
	var employees []database.Employee
	if err := db.Find(&employees).Where("company_id = ?", idInt).Error; err != nil {
		return nil, err
	}
	var nodes []*model.Employee
	for _, e := range employees {
		nodes = append(nodes, model.NewEmployee(&e))
	}
	var pageInfo model.PaginationInfo
	pageInfo.Page = 1
	pageInfo.HasPreviousPage = false
	pageInfo.HasNextPage = false
	pageInfo.Count = len(nodes)
	pageInfo.TotalCount = len(nodes)
	var result model.EmployeePagination
	result.PageInfo = &pageInfo
	result.Nodes = nodes
	return &result, nil
}