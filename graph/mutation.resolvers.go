package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/eyasuyuki/learn_gql/graph/generated"
	"github.com/eyasuyuki/learn_gql/graph/model"
	"github.com/eyasuyuki/learn_gql/service/mutation"
)

func (r *mutationResolver) CreateCompany(ctx context.Context, input model.CreateCompanyInput) (*model.Company, error) {
	return mutation.CreateCompany(r.DB, input)
}

func (r *mutationResolver) UpdateCompany(ctx context.Context, input model.UpdateCompanyInput) (*model.Company, error) {
	return mutation.UpdateCompany(r.DB, input)
}

func (r *mutationResolver) DeleteCompany(ctx context.Context, id string) (bool, error) {
	return mutation.DeleteCompany(r.DB, id)
}

func (r *mutationResolver) CreateDepartment(ctx context.Context, input model.CreateDepartmentInput) (*model.Department, error) {
	return mutation.CreateDepartment(r.DB, input)
}

func (r *mutationResolver) UpdateDepartment(ctx context.Context, input model.UpdateDepartmentInput) (*model.Department, error) {
	return mutation.UpdateDepartment(r.DB, input)
}

func (r *mutationResolver) DeleteDepartment(ctx context.Context, id string) (bool, error) {
	return mutation.DeleteDepartment(r.DB, id)
}

func (r *mutationResolver) SetCompanyToDepartment(ctx context.Context, id string, companyID string) (*model.Department, error) {
	return mutation.SetCompanyToDepartment(r.DB, id, companyID)
}

func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.CreateEmployeeInput) (*model.Employee, error) {
	return mutation.CreateEmployee(r.DB, input)
}

func (r *mutationResolver) UpdateEmployee(ctx context.Context, input model.UpdateEmployeeInput) (*model.Employee, error) {
	return mutation.UpdateEmployee(r.DB, input)
}

func (r *mutationResolver) DeleteEmployee(ctx context.Context, id string) (bool, error) {
	return mutation.DeleteEmployee(r.DB, id)
}

func (r *mutationResolver) SetDepartmentToEmployee(ctx context.Context, id string, departmentID string) (*model.Employee, error) {
	return mutation.SetDepartmentToEmployee(r.DB, id, departmentID)
}

func (r *mutationResolver) SetCompanyToEmployee(ctx context.Context, id string, companyID string) (*model.Employee, error) {
	return mutation.SetCompanyToEmployee(r.DB, id, companyID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
