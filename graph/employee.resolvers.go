package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/eyasuyuki/learn_gql/graph/generated"
	"github.com/eyasuyuki/learn_gql/graph/model"
	"github.com/eyasuyuki/learn_gql/service/employee"
)

func (r *employeeResolver) Department(ctx context.Context, obj *model.Employee) (*model.Department, error) {
	return employee.Department(r.DB, obj)
}

func (r *employeeResolver) Company(ctx context.Context, obj *model.Employee) (*model.Company, error) {
	return employee.Company(r.DB, obj)
}

// Employee returns generated.EmployeeResolver implementation.
func (r *Resolver) Employee() generated.EmployeeResolver { return &employeeResolver{r} }

type employeeResolver struct{ *Resolver }
