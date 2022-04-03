package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/eyasuyuki/learn_gql/graph/generated"
	"github.com/eyasuyuki/learn_gql/graph/model"
)

func (r *departmentResolver) Company(ctx context.Context, obj *model.Department) (*model.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *departmentResolver) Employees(ctx context.Context, obj *model.Department) (*model.EmployeePagination, error) {
	panic(fmt.Errorf("not implemented"))
}

// Department returns generated.DepartmentResolver implementation.
func (r *Resolver) Department() generated.DepartmentResolver { return &departmentResolver{r} }

type departmentResolver struct{ *Resolver }
