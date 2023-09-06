package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/taimoor99/three-tier-golang/app/delivery/graphql"
	"github.com/taimoor99/three-tier-golang/app/delivery/rest"
	"github.com/taimoor99/three-tier-golang/app/models"
	"github.com/taimoor99/three-tier-golang/app/repositories/mongo"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	goc, err := mongo.GetSession(ctx, os.Getenv("DATABASE_NAME"))
	if err != nil {
		fmt.Println("[mongo] mongodb connection error")
		panic(err)
	}
	defer goc.Client().Disconnect(ctx)

	r := chi.NewRouter()
	// Basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // AllowOriginFunc:  func(r *rest.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	repo := mongo.NewUserRepository(goc)
	model := models.NewUserModel(repo)
	ctrl := rest.NewUser(model)

	rest.NewRouter(r, ctrl)
	graphql.InitGraphqlRoute(r, model)

	fmt.Println("server started on PORT " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
