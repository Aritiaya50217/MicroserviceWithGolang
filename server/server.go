package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Aritiaya50217/MicroserviceWithGolang/config"
	middlewareHandler "github.com/Aritiaya50217/MicroserviceWithGolang/modules/middleware/middlewareHandler"
	middlewareRepository "github.com/Aritiaya50217/MicroserviceWithGolang/modules/middleware/middlewareRepository"
	middlewareUsecase "github.com/Aritiaya50217/MicroserviceWithGolang/modules/middleware/middlewareUsecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	server struct {
		app        *echo.Echo
		db         *mongo.Client
		cfg        *config.Config
		middleware middlewareHandler.MiddlewareHandlerService
	}
)

func newMiddleware(cfg *config.Config) middlewareHandler.MiddlewareHandlerService {
	repo := middlewareRepository.NewMiddlewareReposiroty()
	usecase := middlewareUsecase.NewMiddlewareUsecase(repo)

	return middlewareHandler.NewMiddlewareHandler(cfg, usecase)
}

func (s *server) gracarefulShutdown(pctx context.Context, quit <-chan os.Signal) {
	log.Printf("start service : %s ", s.cfg.App.Name)
	<-quit
	log.Printf("shutting down service : %s ", s.cfg.App.Name)

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	if err := s.app.Shutdown(ctx); err != nil {
		log.Fatal("Error : ", err.Error())
	}
}

func (s *server) httpListening() {
	if err := s.app.Start(s.cfg.App.Url); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error : %v ", err)
	}
}

// start server
func Start(pctx context.Context, cfg *config.Config, db *mongo.Client) {
	s := &server{
		app:        echo.New(),
		db:         db,
		cfg:        cfg,
		middleware: newMiddleware(cfg),
	}
	// basic middlewre
	// request timeout
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		Timeout:      30 * time.Second,
		ErrorMessage: "Error Request time out",
	}))

	// cors ( set การเข้าถึง api)
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
	}))

	// body limit
	s.app.Use(middleware.BodyLimit("10M"))

	// custom middleware
	switch cfg.App.Name {
	case "auth":
		s.authService()
	case "player":
		s.playerService()
	case "item":
		s.itemService()
	case "inventory":
		s.inventoryService()
	case "payment":
		s.paymentService()
	}

	// graceful shutdown (การสั่งให้ server หยุดทำงาน)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	s.app.Use(middleware.Logger())

	go s.gracarefulShutdown(pctx, quit)

	// listening
	s.httpListening()
}
