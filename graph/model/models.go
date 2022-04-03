package model

type Employee struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Gender        Gender `json:"gender"`
	Email         string `json:"email"`
	LatestLoginAt string `json:"latestLoginAt"`
	//  扶養家族の人数
	DependenstNum int `json:"dependenstNum"`
	//  管理職かどうか
	IsManager  bool        `json:"isManager"`
	DepartmentID string `json:"department"`
	CompanyID    string    `json:"company"`
}

func (Employee) IsNode() {}
