package database

type Company struct {
	ID             string `gorm:"primaryKey"`
	CompanyName    string
	Representative string
	PhoneNumber    string
}

type Department struct {
	ID             string `gorm:"primaryKey"`
	DepartmentName string
	Email          string
}


type Employee struct {
	ID            string `gorm:"primaryKey"`
	Name          string
	Gender        string
	Email         string
	LatestLoginAt string
	//  扶養家族の人数
	DependentsNum int
	//  管理職かどうか
	IsManager  bool
	DepartmentID string
	CompanyID    string
}

