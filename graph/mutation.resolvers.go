package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/eyasuyuki/learn_gql/graph/generated"
	"github.com/eyasuyuki/learn_gql/graph/model"
)

func (r *mutationResolver) CreateCompany(ctx context.Context, input model.CreateCompanyInput) (*model.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCompany(ctx context.Context, input model.UpdateCompanyInput) (*model.Company, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCompany(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateDepartment(ctx context.Context, input model.CreateDepartmentInput) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDepartment(ctx context.Context, input model.UpdateDepartmentInput) (*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDepartment(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.CreateEmployeeInput) (*model.Employee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEmployee(ctx context.Context, input model.UpdateEmployeeInput) (*model.Employee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEmployee(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
