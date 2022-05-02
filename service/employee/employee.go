package employee

import (
	"github.com/eyasuyuki/learn_gql/database"
	"github.com/eyasuyuki/learn_gql/graph/model"
	"gorm.io/gorm"
)

func Department(db *gorm.DB, obj *model.Employee) (*model.Department, error) {
	departmentId, err := database.DepartmentIDFromBase64(obj.DepartmentID)
	if err != nil {
		return nil, err
	}
	var department database.Department
	if err = db.Where("id = ?", departmentId).Find(&department).Error; err != nil {
		return nil, err
	}
	return model.NewDepartment(&department), nil
}


func Company(db *gorm.DB, obj *model.Employee) (*model.Company, error) {
	companyId, err := database.CompanyIDFromBase64(obj.CompanyID)
	if err != nil {
		return nil, err
	}
	var company database.Company
	if err = db.Where("id = ?", companyId).Find(&company).Error; err != nil {
		return nil, err
	}
	return model.NewCompany(&company), nil
}