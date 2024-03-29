package main

import (
	"backend/internal/config"
	controllerhttp "backend/internal/controller/http"
	AuthRepo "backend/internal/repository/auth"
	PasteRepo "backend/internal/repository/paste"
	UserRepo "backend/internal/repository/user"
	AuthService "backend/internal/services/auth"
	PasteService "backend/internal/services/paste"
	"backend/internal/services/pylinter"
	UserService "backend/internal/services/user"
	"backend/pkg/middleware"
	"context"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func init() {
	if err := godotenv.Load("backend/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()

	clientOptions := options.Client().ApplyURI(cfg.MongoURL)
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err := mongoClient.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	db := mongoClient.Database(cfg.Database)

	userCollection := db.Collection(cfg.UserCollection)
	pasteCollection := db.Collection(cfg.PasteCollection)

	userRepo := UserRepo.NewUserRepository(userCollection)
	authRepo := AuthRepo.NewUserRepository(userCollection)
	pasteRepo := PasteRepo.NewPasteRepository(pasteCollection)
	pylint := pylinter.PylintLinter{}

	userService := UserService.NewService(UserService.Deps{UserRepo: userRepo})
	pasteService := PasteService.NewService(PasteService.Deps{PasteRepo: pasteRepo})
	linterService := pylinter.NewClient(&pylint)
	authService := AuthService.NewService(AuthService.Deps{AuthRepo: authRepo})

	ctr := controllerhttp.NewController(controllerhttp.UserService{UserManagement: userService},
		controllerhttp.PasteService{PasteManagement: pasteService},
		controllerhttp.LinterService{Linter: linterService},
		controllerhttp.AuthService{AuthManagement: authService})

	router := ctr.NewRouter()

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	handler := c.Handler(router)

	err = http.ListenAndServe(cfg.Port, middleware.PanicRecovery(middleware.LogRequest(handler)))
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
