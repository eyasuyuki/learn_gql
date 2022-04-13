package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/eyasuyuki/learn_gql/database"

	"github.com/eyasuyuki/learn_gql/graph/generated"
	"github.com/eyasuyuki/learn_gql/graph/model"
)

func (r *departmentResolver) Company(ctx context.Context, obj *model.Department) (*model.Company, error) {
	idInt, err := database.IdFromBase64(database.DEPARTMENT_PREFIX, obj.CompanyID)
	if err != nil {
		return nil, err
	}
	var company database.Company
	if err = r.DB.Find(&company, idInt).Error; err != nil {
		return nil, err
	}
	return model.NewCompany(&company), nil
}

func (r *departmentResolver) Employees(ctx context.Context, obj *model.Department) (*model.EmployeePagination, error) {
	idInt, err := database.IdFromBase64(database.DEPARTMENT_PREFIX, obj.CompanyID)
	if err != nil {
		return nil, err
	}
	var employees []database.Employee
	if err = r.DB.Find(&employees).Where("department_id = ? ", idInt).Error; err != nil {
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

// Department returns generated.DepartmentResolver implementation.
func (r *Resolver) Department() generated.DepartmentResolver { return &departmentResolver{r} }

type departmentResolver struct{ *Resolver }
