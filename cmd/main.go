package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	middleKit "github.com/laironacosta/kit-go/middleware/echo"
	pgKit "github.com/laironacosta/kit-go/postgresql"
	repository "github.com/laironacosta/ms-echo-go-layout/infrastructure/adapters/repositories"
	"github.com/laironacosta/ms-echo-go-layout/infrastructure/framework/controllers"
	"github.com/laironacosta/ms-echo-go-layout/infrastructure/framework/middleware"
	router "github.com/laironacosta/ms-echo-go-layout/infrastructure/framework/routers"
	"github.com/laironacosta/ms-echo-go-layout/internal/app/users"
	"github.com/laironacosta/ms-echo-go-layout/pkg/i18n"
	"github.com/laironacosta/ms-echo-go-layout/script/migrations"
	"github.com/pkg/errors"
)

// cfg is the struct type that contains fields that stores the necessary configuration
// gathered from the environment.
var cfg struct {
	DBUser string `envconfig:"DB_USER" default:"postgres"`
	DBPass string `envconfig:"DB_PASS" default:"postgres"`
	DBName string `envconfig:"DB_NAME" default:"postgres"`
	DBHost string `envconfig:"DB_HOST" default:"localhost"`
	DBPort int    `envconfig:"DB_PORT" default:"5432"`
	Locale string `envconfig:"LOCALE"  default:"es"`
}

func main() {
	echo := echo.New()
	echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency_human={latency_human}\n",
	}))
	echo.Use(middleware.Recover())

	if err := envconfig.Process("LIST", &cfg); err != nil {
		err = errors.Wrap(err, "parse environment variables")
		return
	}

	db := pgKit.NewPgDB(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.DBHost, cfg.DBPort),
		User:     cfg.DBUser,
		Password: cfg.DBPass,
		Database: cfg.DBName,
	})
	migrations.Init(db)

	_ = i18n.NewI18n(cfg.Locale)

	userRepo := repository.NewUserRepository(db)
	userService := users.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	errorHanlderMiddle := middleKit.NewErrorHandlerMiddleware()
	i18nMiddle := middlewares.NewI18nMiddleware()

	r := router.NewRouter(echo, userController, errorHanlderMiddle, i18nMiddle)
	r.Init()

	echo.Start(":8080") // listen and serve on 0.0.0.0:8080
}
