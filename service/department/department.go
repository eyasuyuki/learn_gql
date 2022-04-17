package department

import (
	"github.com/eyasuyuki/learn_gql/database"
	"github.com/eyasuyuki/learn_gql/graph/model"
	"gorm.io/gorm"
)

func Company(db *gorm.DB, obj *model.Department) (*model.Company, error) {
	idInt, err := database.IdFromBase64(database.DEPARTMENT_PREFIX, obj.CompanyID)
	if err != nil {
		return nil, err
	}
	var company database.Company
	if err = db.Find(&company, idInt).Error; err != nil {
		return nil, err
	}
	return model.NewCompany(&company), nil
}

func Employees(db *gorm.DB, obj *model.Department) (*model.EmployeePagination, error) {
	idInt, err := database.IdFromBase64(database.DEPARTMENT_PREFIX, obj.CompanyID)
	if err != nil {
		return nil, err
	}
	var employees []database.Employee
	if err = db.Find(&employees).Where("department_id = ? ", idInt).Error; err != nil {
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
