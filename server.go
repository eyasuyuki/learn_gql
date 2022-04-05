package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/eyasuyuki/learn_gql/graph"
	"github.com/eyasuyuki/learn_gql/graph/generated"
)

const defaultPort = "8080"

const defaultDsn = "learngql:learngql@tcp(127.0.0.1:3306)/learndb?charset=utf8&parseTime=True&loc=Local"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dsn := os.Getenv("DATA_SET_NAME")
	if dsn == "" {
		dsn = defaultDsn
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(any(err))
	}
	if db == nil {
		panic(any(err))
	}
	defer func() {
		if db != nil {
			sqlDb, err := db.DB()
			if err != nil {
				panic(any(err))
			}
			sqlDb.Close()
		}
	}()


	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
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
	http.Handle("/q             uery", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
