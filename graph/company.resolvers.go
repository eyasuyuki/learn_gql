package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/eyasuyuki/learn_gql/database"

	"github.com/eyasuyuki/learn_gql/graph/generated"
	"github.com/eyasuyuki/learn_gql/graph/model"
)

func (r *companyResolver) Departments(ctx context.Context, obj *model.Company) (*model.DepartmentPagination, error) {
	idInt, err := database.IdFromBase64(database.COMPANY_PREFIX, obj.ID)
	if err != nil {
		return nil, err
	}
	var departments []database.Department
	if err = r.DB.Find(&departments).Where("company_id = ? ", idInt).Error; err != nil {
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

func (r *companyResolver) Employees(ctx context.Context, obj *model.Company) (*model.EmployeePagination, error) {
	idInt, err := database.IdFromBase64(database.COMPANY_PREFIX, obj.ID)
	if err != nil {
		return nil, err
	}
	var employees []database.Employee
	if err := r.DB.Find(&employees).Where("company_id = ?", idInt).Error; err != nil {
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

// Company returns generated.CompanyResolver implementation.
func (r *Resolver) Company() generated.CompanyResolver { return &companyResolver{r} }

type companyResolver struct{ *Resolver }
