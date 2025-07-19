package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/arshamroshannejad/nuke"
	"github/arshamroshannejad/squidshop-backend/config"
	"github/arshamroshannejad/squidshop-backend/internal/database"
	"github/arshamroshannejad/squidshop-backend/internal/logger"
	"github/arshamroshannejad/squidshop-backend/internal/router"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	zapLog, err := logger.New(cfg.App.Debug)
	if err != nil {
		panic("failed to create zap logger instance: " + err.Error())
	}
	defer zapLog.Sync()
	db, err := database.OpenDB(cfg)
	if err != nil {
		zapLog.Fatal("failed to connect postgres", zap.Error(err))
	}
	defer db.Close()
	zapLog.Info(
		"connected to postgres",
		zap.String("host", cfg.Postgres.Host),
		zap.Int("port", cfg.Postgres.Port),
	)
	redisDB, err := database.OpenRedis(cfg)
	if err != nil {
		zapLog.Fatal("failed to connect redis", zap.Error(err))
	}
	defer redisDB.Close()
	zapLog.Info(
		"connected to redis",
		zap.String("host", cfg.Redis.Host),
		zap.Int("port", cfg.Redis.Port),
	)
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.App.Port),
		Handler:      router.SetupRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	nuke.Background(func() {
		zapLog.Info("starting server", zap.Int("port", cfg.App.Port))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zapLog.Fatal("failed to start server", zap.Error(err))
		}
	})
	<-quit
	zapLog.Info("server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zapLog.Error("server shutdown failed", zap.Error(err))
	} else {
		zapLog.Info("server shutdown completed")
	}
}
