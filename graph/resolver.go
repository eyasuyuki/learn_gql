package graph

import "github.com/eyasuyuki/learn_gql/dataloader"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	CompanyLoader *dataloader.CompanyLoader
}
