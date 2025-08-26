package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fernandobandeira/djinn/backend/internal/graph/generated"
	"github.com/fernandobandeira/djinn/backend/internal/graph/resolver"
)

// graphqlHandler creates the GraphQL handler
func (s *Server) graphqlHandler() http.Handler {
	// Create resolver with dependencies
	res := resolver.NewResolver(s.db, s.logger)
	
	// Create GraphQL server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: res,
	}))
	
	// Configure server options
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		s.logger.Error("GraphQL panic recovered", "panic", err)
		return fmt.Errorf("internal server error")
	})
	
	return srv
}

// playgroundHandler creates the GraphQL playground handler
func (s *Server) playgroundHandler() http.Handler {
	return playground.Handler("GraphQL Playground", "/api/graphql")
}