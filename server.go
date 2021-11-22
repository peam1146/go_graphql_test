package main

import (
	"graph_jwt/graph/generated"
	"graph_jwt/graph/resovlers"
	"os"

	// echo
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

const defaultPort = "8080"

func playgroundHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func graphqlHandler() echo.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e := echo.New()

	e.GET("/", playgroundHandler())
	e.POST("/query", graphqlHandler())

	e.Logger.Fatal(e.Start(":" + port))
}
