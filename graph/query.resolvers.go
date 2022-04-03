package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/eyasuyuki/learn_gql/graph/generated"
	"github.com/eyasuyuki/learn_gql/graph/model"
)

func (r *queryResolver) Company(ctx context.Context, id string) (*model.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Companies(ctx context.Context, limit int, offset *int) (*model.CompanyPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Department(ctx context.Context, id string) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Departments(ctx context.Context, limit int, offset *int) (*model.DepertmentPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Employee(ctx context.Context, id string) (*model.Employee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Employees(ctx context.Context, limit int, offset *int, email *string, gender *model.Gender, isManager *bool, hasDepartment *bool) (*model.EmployeePagination, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
