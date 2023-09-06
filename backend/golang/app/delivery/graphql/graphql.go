package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/taimoor99/three-tier-golang/app/delivery/graphql/generated"
	"github.com/taimoor99/three-tier-golang/app/delivery/graphql/resolvers"
	"github.com/taimoor99/three-tier-golang/app/models"
)

func InitGraphqlRoute(r *chi.Mux, model models.UserModel) {
	// The base path that users would use is POST /graphql which is fairly
	// idiomatic.
	r.Route("/graphql", func(r chi.Router) {
		// Initialize the dataloaders as middleware into our route
		schema := generated.NewExecutableSchema(generated.Config{
			Resolvers: &resolvers.Resolver{
				UserModel: model,
			},
			Directives: generated.DirectiveRoot{},
			Complexity: generated.ComplexityRoot{},
		})

		srv := handler.NewDefaultServer(schema)
		srv.Use(extension.FixedComplexityLimit(300))
		r.Handle("/", srv)
	})

	gqlPlayground := playground.Handler("api-gateway", "/graphql")
	r.Get("/", gqlPlayground)

	return
}

