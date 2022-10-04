package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/stneto1/gqlgen-example/pkg/graph"
	"github.com/stneto1/gqlgen-example/pkg/graph/generated"
	"github.com/stneto1/gqlgen-example/pkg/middlewares"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.GinContextToContextMiddleware())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	r.Run()
}

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
