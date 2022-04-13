package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/eyasuyuki/learn_gql/database"
	"github.com/eyasuyuki/learn_gql/graph/generated"
	"github.com/eyasuyuki/learn_gql/graph/model"
)

func (r *mutationResolver) CreateCompany(ctx context.Context, input model.CreateCompanyInput) (*model.Company, error) {
	company := database.NewCompany(input)
	r.DB.Create(company)
	result := model.NewCompany(company)
	return result, nil
}

func (r *mutationResolver) UpdateCompany(ctx context.Context, input model.UpdateCompanyInput) (*model.Company, error) {
	company := database.NewCompanyUpdate(input)
	r.DB.Save(company)
	result := model.NewCompany(company)
	return result, nil
}

func (r *mutationResolver) DeleteCompany(ctx context.Context, id string) (bool, error) {
	idInt, err := database.IdFromBase64(database.COMPANY_PREFIX, id)
	if err != nil {
		return false, err
	}
	r.DB.Delete(&database.Company{}, idInt)
	return true, nil
}

func (r *mutationResolver) CreateDepartment(ctx context.Context, input model.CreateDepartmentInput) (*model.Department, error) {
	department := database.NewDepartment(input)
	r.DB.Create(department)
	return model.NewDepartment(department), nil
}

func (r *mutationResolver) UpdateDepartment(ctx context.Context, input model.UpdateDepartmentInput) (*model.Department, error) {
	department := database.NewDepartmentUpdate(input)
	r.DB.Save(department)
	return model.NewDepartment(department), nil
}

func (r *mutationResolver) DeleteDepartment(ctx context.Context, id string) (bool, error) {
	idInt, err := database.IdFromBase64(database.DEPARTMENT_PREFIX, id)
	if err != nil {
		return false, err
	}
	r.DB.Delete(&database.Department{}, idInt)
	return true, nil
}

func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.CreateEmployeeInput) (*model.Employee, error) {
	employee := database.NewEmployee(input)
	r.DB.Create(employee)
	return model.NewEmployee(employee), nil
}

func (r *mutationResolver) UpdateEmployee(ctx context.Context, input model.UpdateEmployeeInput) (*model.Employee, error) {
	employee := database.NewEmployeeUpdate(input)
	r.DB.Save(employee)
	return model.NewEmployee(employee), nil
}

func (r *mutationResolver) DeleteEmployee(ctx context.Context, id string) (bool, error) {
	idInt, err := database.IdFromBase64(database.EMPLOYEE_PREFIX, id)
	if err != nil {
		return false, err
	}
	r.DB.Delete(&database.Employee{}, idInt)
	return true, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
