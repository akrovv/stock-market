package main

import (
	"context"
	"fmt"

	redisdb "github.com/akrovv/client/internal/adapter/redisDB"
	"github.com/akrovv/client/internal/handler/rest"
	restmiddleware "github.com/akrovv/client/internal/handler/rest/rest_middleware"
	"github.com/akrovv/client/internal/service"
	"github.com/akrovv/client/pkg/hasher"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func main() {
	e := echo.New()

	ctxRedis := context.Background()
	redisDB := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})

	hasher := hasher.NewHasher([]byte("idk"))

	status := redisDB.Ping(ctxRedis)

	if status.Err() != nil {
		fmt.Println(status.Err())
		return
	}

	sessionStorage := redisdb.NewSessionStorage(ctxRedis, redisDB, hasher)

	sessionService := service.NewSessionService(sessionStorage)
	userService := service.NewUserService()

	userHander := rest.NewUserHandler(userService, sessionService)
	// rootHandler := rest.NewRootHandler()

	e.Use(restmiddleware.AuthMiddleware(sessionService))

	e.Static("/", "front/build/")
	e.File("/register", "front/build/")
	e.File("/user-profile", "front/build/")

	e.POST("/api/user/register", userHander.Register)
	e.POST("/api/user/login", userHander.Login)
	e.GET("/api/user/profile", userHander.Profile)
	e.GET("/api/user/logout", userHander.Logout)
	// e.GET("/login", userHander.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
