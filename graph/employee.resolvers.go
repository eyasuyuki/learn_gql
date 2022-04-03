package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/eyasuyuki/learn_gql/graph/generated"
	"github.com/eyasuyuki/learn_gql/graph/model"
)

func (r *employeeResolver) Department(ctx context.Context, obj *model.Employee) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *employeeResolver) Company(ctx context.Context, obj *model.Employee) (*model.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

// Employee returns generated.EmployeeResolver implementation.
func (r *Resolver) Employee() generated.EmployeeResolver { return &employeeResolver{r} }

type employeeResolver struct{ *Resolver }
