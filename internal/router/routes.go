package router

import (
	"database/sql"
	"log/slog"
	"net/http"

	_ "github.com/arshamroshannejad/squidshop-backend/api"
	"github.com/arshamroshannejad/squidshop-backend/config"
	"github.com/arshamroshannejad/squidshop-backend/internal/handler"
	"github.com/arshamroshannejad/squidshop-backend/internal/middleware"
	"github.com/arshamroshannejad/squidshop-backend/internal/repository"
	"github.com/arshamroshannejad/squidshop-backend/internal/service"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"
	swagger "github.com/swaggo/http-swagger"
)

func SetupRoutes(db *sql.DB, redisDB *redis.Client, logger *slog.Logger, cfg *config.Config) http.Handler {
	mux := http.NewServeMux()
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories, redisDB, logger, cfg)
	handlers := handler.NewHandler(services)
	mux.Handle(
		"POST /api/v1/auth",
		middleware.RateLimiter(0.008333, 1)(http.HandlerFunc(handlers.Auth().AuthUserHandler)),
	)
	mux.HandleFunc(
		"POST /api/v1/auth/verify",
		handlers.Auth().VerifyAuthUserHandler,
	)
	mux.Handle(
		"GET /api/v1/user/profile",
		middleware.RequireAuth(cfg)(http.HandlerFunc(handlers.User().UserProfileHandler)),
	)
	mux.Handle("/docs/", swagger.Handler(
		swagger.URL("doc.json"),
		swagger.DeepLinking(true),
		swagger.DocExpansion("none"),
		swagger.DomID("swagger-ui"),
	))
	return cors.Default().Handler(
		middleware.Logger(middleware.Timeout(mux)),
	)
}
