package app

import (
	"context"
	"log"
	"net/http"

	"github.com/ilhamosaurus/HRIS/internal/container"
	"github.com/ilhamosaurus/HRIS/internal/routes"
	"github.com/ilhamosaurus/HRIS/pkg/db"
	"github.com/ilhamosaurus/HRIS/pkg/setting"
	"github.com/ilhamosaurus/HRIS/pkg/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	container  *container.Container
	db         *gorm.DB
	hasher     *util.Hasher
	addr       string
	httpServer *http.Server
}

func NewApp(addr string) (*App, error) {
	db, err := db.InitDB()
	if err != nil {
		return nil, err
	}

	hasher := util.NewHasher(setting.Server.Secret)
	depContainer, err := container.NewContainer(db, hasher)
	if err != nil {
		return nil, err
	}

	if err := Seed(db, hasher, depContainer.UserDAO); err != nil {
		return nil, err
	}

	e := echo.New()
	routes := routes.NewRoutes(depContainer)
	routes.SetupRoutes(e)

	httpServer := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	return &App{
		container:  depContainer,
		db:         db,
		hasher:     hasher,
		addr:       addr,
		httpServer: httpServer,
	}, nil
}

func (a *App) Start() error {
	log.Println("Server starting on localhost" + a.addr)
	if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	if a.httpServer != nil {
		if err := a.httpServer.Shutdown(ctx); err != nil {
			return err
		}
		log.Println("Server shutdown gracefully")
	}
	return nil
}
