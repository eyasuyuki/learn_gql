package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/eyasuyuki/learn_gql/service/company"

	"github.com/eyasuyuki/learn_gql/graph/generated"
	"github.com/eyasuyuki/learn_gql/graph/model"
)

func (r *companyResolver) Departments(ctx context.Context, obj *model.Company) (*model.DepartmentPagination, error) {
	return company.Departments(r.DB, obj)
}

func (r *companyResolver) Employees(ctx context.Context, obj *model.Company) (*model.EmployeePagination, error) {
	return company.Employees(r.DB, obj)
}

// Company returns generated.CompanyResolver implementation.
func (r *Resolver) Company() generated.CompanyResolver { return &companyResolver{r} }

type companyResolver struct{ *Resolver }
