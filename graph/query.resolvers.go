package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/eyasuyuki/learn_gql/graph/generated"
	"github.com/eyasuyuki/learn_gql/graph/model"
	"github.com/eyasuyuki/learn_gql/service/query"
)

func (r *queryResolver) Company(ctx context.Context, id string) (*model.Company, error) {
	return query.Company(r.DB, id)
}

func (r *queryResolver) Companies(ctx context.Context, limit int, offset *int) (*model.CompanyPagination, error) {
	return query.Companies(r.DB, limit, offset)
}

func (r *queryResolver) Department(ctx context.Context, id string) (*model.Department, error) {
	return query.Department(r.DB, id)
}

func (r *queryResolver) Departments(ctx context.Context, limit int, offset *int) (*model.DepartmentPagination, error) {
	return query.Departments(r.DB, limit, offset)
}

func (r *queryResolver) Employee(ctx context.Context, id string) (*model.Employee, error) {
	return query.Employee(r.DB, id)
}

func (r *queryResolver) Employees(ctx context.Context, limit int, offset *int, email *string, gender *model.Gender, isManager *bool, hasDepartment *bool) (*model.EmployeePagination, error) {
	return query.Employees(r.DB, limit, offset, email, gender, isManager, hasDepartment)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
