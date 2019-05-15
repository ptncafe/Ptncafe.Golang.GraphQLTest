package main

import (
    "github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"Ptncafe.Golang.GraphQLTest/generated"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
    // NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
    // Setting up Gin  
    r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
    r.Run()
}  