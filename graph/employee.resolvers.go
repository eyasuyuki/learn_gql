package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/eyasuyuki/learn_gql/database"

	"github.com/eyasuyuki/learn_gql/graph/generated"
	"github.com/eyasuyuki/learn_gql/graph/model"
)

func (r *employeeResolver) Department(ctx context.Context, obj *model.Employee) (*model.Department, error) {
	idInt, err := database.IdFromBase64(database.DEPARTMENT_PREFIX, obj.DepartmentID)
	if err != nil {
		return nil, err
	}
	var department database.Department
	if err = r.DB.Find(&department, idInt).Error; err != nil {
		return nil, err
	}
	return model.NewDepartment(&department), nil
}

func (r *employeeResolver) Company(ctx context.Context, obj *model.Employee) (*model.Company, error) {

	idInt, err := database.IdFromBase64(database.COMPANY_PREFIX, obj.CompanyID)
	if err != nil {
		return nil, err
	}
	var company database.Company
	if err = r.DB.Find(&company, idInt).Error; err != nil {
		return nil, err
	}
	return model.NewCompany(&company), nil
}

// Employee returns generated.EmployeeResolver implementation.
func (r *Resolver) Employee() generated.EmployeeResolver { return &employeeResolver{r} }

type employeeResolver struct{ *Resolver }
