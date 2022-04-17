package employee

import (
	"github.com/eyasuyuki/learn_gql/database"
	"github.com/eyasuyuki/learn_gql/graph/model"
	"gorm.io/gorm"
)

func Department(db *gorm.DB, obj *model.Employee) (*model.Department, error) {
	idInt, err := database.IdFromBase64(database.DEPARTMENT_PREFIX, obj.DepartmentID)
	if err != nil {
		return nil, err
	}
	var department database.Department
	if err = db.Find(&department, idInt).Error; err != nil {
		return nil, err
	}
	return model.NewDepartment(&department), nil
}


func Company(db *gorm.DB, obj *model.Employee) (*model.Company, error) {
	idInt, err := database.IdFromBase64(database.COMPANY_PREFIX, obj.CompanyID)
	if err != nil {
		return nil, err
	}
	var company database.Company
	if err = db.Find(&company, idInt).Error; err != nil {
		return nil, err
	}
	return model.NewCompany(&company), nil
}