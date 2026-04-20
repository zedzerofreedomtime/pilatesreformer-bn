package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/config"
	httpapi "github.com/zedzerofreedomtime/pilatesreformer/api/internal/http"
	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/http/handlers"
	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/repository"
	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/service"
)

type App struct {
	cfg    config.Config
	router interface{ Run(addr ...string) error }
}

func New(ctx context.Context) (*App, error) {
	cfg := config.Load()

	dbPool, err := pgxpool.New(ctx, cfg.PostgresDSN)
	if err != nil {
		return nil, err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	repo := repository.New(dbPool)
	authService := service.NewAuthService(repo, redisClient, cfg.SessionTTL)
	catalogService := service.NewCatalogService(repo, redisClient, cfg.CacheTTL)
	bookingService := service.NewBookingService(repo)
	adminService := service.NewAdminService(repo, catalogService)
	handler := handlers.New(repo, authService, catalogService, bookingService, adminService)
	router := httpapi.NewRouter(cfg, handler, authService)

	return &App{
		cfg:    cfg,
		router: router,
	}, nil
}

func (a *App) Run() error {
	return a.router.Run(fmt.Sprintf(":%s", a.cfg.HTTPPort))
}
