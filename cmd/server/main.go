package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arshamroshannejad/squidshop-backend/config"
	"github.com/arshamroshannejad/squidshop-backend/internal/database"
	"github.com/arshamroshannejad/squidshop-backend/internal/router"
)

//	@title						squidshop-backend
//	@version					0.1.0
//	@host						api.squidshop.ir
//	@description				Api for managing product and everything related to E-commerce shop
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				Arsham Roshannejad
//	@contact.url				arshamroshannejad.ir
//	@contact.email				arshamdev2001@gmail.com
//	@license.name				MIT
//	@license.url				https://www.mit.edu/~amini/LICENSE.md
//	@BasePath					/api/v1
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
func main() {
	cfg, err := config.New()
	if err != nil {
		panic("failed to load config variables: " + err.Error())
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	db, err := database.OpenDB(cfg)
	if err != nil {
		logger.Error("failed to connect postgres", "error:", err)
		return
	}
	defer db.Close()
	logger.Info(
		"connected to postgres",
		"host", cfg.Postgres.Host,
		"port", cfg.Postgres.Port,
	)
	redisDB, err := database.OpenRedis(cfg)
	if err != nil {
		logger.Error("failed to connect redis", "error:", err)
	}
	defer redisDB.Close()
	logger.Info(
		"connected to redis",
		"host", cfg.Redis.Host,
		"port", cfg.Redis.Port,
	)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.App.Port),
		Handler:      router.SetupRoutes(db, redisDB, logger, cfg),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("panic occurred in running server", "error", r)
			}
		}()
		logger.Info("server is starting...", "port", cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("failed to start server", "error", err)
		}
	}()
	<-quit
	logger.Info("server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("failed to shutdown server", "error", err)
	} else {
		logger.Info("server is shutdown successfully")
	}
}
