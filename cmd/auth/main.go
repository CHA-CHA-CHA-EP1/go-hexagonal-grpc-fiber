package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/CHA-CHA-CHA-EP1/go-hexagonal-grpc-fiber/configs"
	"github.com/CHA-CHA-CHA-EP1/go-hexagonal-grpc-fiber/internal/auth/controller"
	"github.com/CHA-CHA-CHA-EP1/go-hexagonal-grpc-fiber/internal/auth/service"
	"github.com/CHA-CHA-CHA-EP1/go-hexagonal-grpc-fiber/internal/auth/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    cfgs := configs.GetConfig()

    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()

    conn, err := mongo.Connect(ctx, options.Client().ApplyURI(cfgs.Database.Uri))
    if err != nil {
        log.Fatal(err)
        return
    }


    err = conn.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
        return
    }

    db := conn.Database(cfgs.Database.Name)

    defer func() {
        if err = conn.Disconnect(ctx); err != nil {
            log.Fatal(err)
        }
    }()



    userRepository := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepository)
    userController := controller.NewUserController(userService)

    fmt.Println("[AUTH] Connected to MongoDB!")

    // Start the server
    app := fiber.New()

    userGroup := app.Group("/api/users")
    userGroup.Get("/:id", userController.GetUserById)

    authGroup := app.Group("/api/auth")
    authGroup.Post("/register", userController.Register)

    app.Listen(":3000")
}
