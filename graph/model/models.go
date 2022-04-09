package model

type Company struct {
	ID             string                `json:"id"`
	CompanyName    string                `json:"companyName"`
	Representative string                `json:"representative"`
	PhoneNumber    string                `json:"phoneNumber"`
}

func (Company) IsNode() {}

type Department struct {
	ID             string              `json:"id"`
	DepartmentName string              `json:"departmentName"`
	Email          string              `json:"email"`
}

func (Department) IsNode() {}


type Employee struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Gender        Gender `json:"gender"`
	Email         string `json:"email"`
	LatestLoginAt string `json:"latestLoginAt"`
	//  扶養家族の人数
	DependentsNum int `json:"dependentsNum"`
	//  管理職かどうか
	IsManager  bool        `json:"isManager"`
	DepartmentID string `json:"department"`
	CompanyID    string    `json:"company"`
}

func (Employee) IsNode() {}

