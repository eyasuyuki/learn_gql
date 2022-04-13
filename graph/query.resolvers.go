package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/eyasuyuki/learn_gql/database"
	"github.com/eyasuyuki/learn_gql/graph/generated"
	"github.com/eyasuyuki/learn_gql/graph/model"
)

func (r *queryResolver) Company(ctx context.Context, id string) (*model.Company, error) {
	idInt, err := database.IdFromBase64(database.COMPANY_PREFIX, id)
	if err != nil {
		return nil, err
	}
	company := &database.Company{}
	if err := r.DB.Find(company, idInt).Error; err != nil {
		return nil, err
	}
	return model.NewCompany(company), nil
}

func (r *queryResolver) Companies(ctx context.Context, limit int, offset *int) (*model.CompanyPagination, error) {
	var companies []database.Company
	var total int64
	r.DB.Model(&database.Company{}).Count(&total)
	var err error
	if offset != nil {
		err = r.DB.Offset(*offset).Limit(limit).Find(&companies).Error
	} else {
		err = r.DB.Limit(limit).Find(&companies).Error
	}
	if err != nil {
		return nil, err
	}
	pageInfo := &model.PaginationInfo{}
	var nodes []*model.Company
	result := &model.CompanyPagination{PageInfo: pageInfo, Nodes: nodes}
	for _, company := range companies {
		nodes = append(nodes, model.NewCompany(&company))
	}
	pageInfo.Page = 1
	if offset != nil {
		pageInfo.Page = ((int(total) - *offset) / limit) + 1
	}
	pageInfo.Count = len(nodes)
	return result, nil
}

func (r *queryResolver) Department(ctx context.Context, id string) (*model.Department, error) {
	idInt, err := database.IdFromBase64(database.DEPARTMENT_PREFIX, id)
	if err != nil {
		return nil, err
	}
	department := &database.Department{}
	if err := r.DB.Select(department, idInt).Error; err != nil {
		return nil, err
	}
	return model.NewDepartment(department), nil
}

func (r *queryResolver) Departments(ctx context.Context, limit int, offset *int) (*model.DepartmentPagination, error) {
	var departments []database.Department
	var total int64
	r.DB.Model(&database.Department{}).Count(&total)
	var err error
	if offset != nil {
		err = r.DB.Offset(*offset).Limit(limit).Find(&departments).Error
	} else {
		err = r.DB.Limit(limit).Find(&departments).Error
	}
	if err != nil {
		return nil, err
	}
	pageInfo := &model.PaginationInfo{}
	var nodes []*model.Department
	result := &model.DepartmentPagination{PageInfo: pageInfo, Nodes: nodes}
	for _, Department := range departments {
		nodes = append(nodes, model.NewDepartment(&Department))
	}
	pageInfo.Page = 1
	if offset != nil {
		pageInfo.Page = ((int(total) - *offset) / limit) + 1
	}
	pageInfo.Count = len(nodes)
	return result, nil
}

func (r *queryResolver) Employee(ctx context.Context, id string) (*model.Employee, error) {
	idInt, err := database.IdFromBase64(database.EMPLOYEE_PREFIX, id)
	if err != nil {
		return nil, err
	}
	Employee := &database.Employee{}
	if err := r.DB.Select(Employee, idInt).Error; err != nil {
		return nil, err
	}
	return model.NewEmployee(Employee), nil
}

func (r *queryResolver) Employees(ctx context.Context, limit int, offset *int, email *string, gender *model.Gender, isManager *bool, hasDepartment *bool) (*model.EmployeePagination, error) {
	var employees []database.Employee
	var total int64
	r.DB.Model(&database.Employee{}).Count(&total)
	var err error
	if offset != nil {
		err = r.DB.Offset(*offset).Limit(limit).Find(&employees).Error
	} else {
		err = r.DB.Limit(limit).Find(&employees).Error
	}
	if err != nil {
		return nil, err
	}
	pageInfo := &model.PaginationInfo{}
	var nodes []*model.Employee
	result := &model.EmployeePagination{PageInfo: pageInfo, Nodes: nodes}
	for _, Employee := range employees {
		nodes = append(nodes, model.NewEmployee(&Employee))
	}
	pageInfo.Page = 1
	if offset != nil {
		pageInfo.Page = ((int(total) - *offset) / limit) + 1
	}
	pageInfo.Count = len(nodes)
	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
