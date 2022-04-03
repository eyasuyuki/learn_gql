package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/eyasuyuki/learn_gql/dataloader"
	"github.com/eyasuyuki/learn_gql/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/eyasuyuki/learn_gql/graph"
	"github.com/eyasuyuki/learn_gql/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	companyLoader := dataloader.NewCompanyLoader(dataloader.CompanyLoaderConfig{
		MaxBatch: 100,
		Wait: 2,
		Fetch: func(keys []string) ([]*model.Company, []error) {
			companies := make([]*model.Company, len(keys))
			errors := make([]error, len(keys))

			// TODO fettch data

			return companies, errors
		},
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		CompanyLoader: companyLoader,
	}}))
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		err.Message = e.Error()
		err.Extensions = map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}
		return err
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
