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
	mux.HandleFunc(
		"GET /api/v1/category",
		handlers.Category().GetAllCategoriesHandler,
	)
	mux.Handle(
		"POST /api/v1/category",
		middleware.RequireAuth(cfg)(
			middleware.RequireAdmin(
				http.HandlerFunc(handlers.Category().CreateCategoryHandler),
			),
		),
	)
	mux.Handle(
		"PUT /api/v1/category/{id}",
		middleware.RequireAuth(cfg)(
			middleware.RequireAdmin(
				http.HandlerFunc(handlers.Category().UpdateCategoryHandler),
			),
		),
	)
	mux.Handle(
		"DELETE /api/v1/category/{id}",
		middleware.RequireAuth(cfg)(
			middleware.RequireAdmin(
				http.HandlerFunc(handlers.Category().DeleteCategoryHandler),
			),
		),
	)
	mux.Handle(
		"GET /api/v1/category/exists",
		middleware.RequireAuth(cfg)(
			middleware.RequireAdmin(
				http.HandlerFunc(handlers.Category().ExistsCategoryHandler),
			),
		),
	)
	mux.HandleFunc(
		"GET /api/v1/product",
		handlers.Product().GetAllProductsHandler,
	)
	mux.HandleFunc(
		"GET /api/v1/product/id/{id}",
		handlers.Product().GetProductByIDHandler,
	)
	mux.HandleFunc(
		"GET /api/v1/product/slug/{slug}",
		handlers.Product().GetProductBySlugHandler,
	)
	mux.Handle(
		"POST /api/v1/product",
		middleware.RequireAuth(cfg)(
			middleware.RequireAdmin(
				http.HandlerFunc(handlers.Product().CreateProductHandler),
			),
		),
	)
	mux.Handle(
		"PUT /api/v1/product/{id}",
		middleware.RequireAuth(cfg)(
			middleware.RequireAdmin(
				http.HandlerFunc(handlers.Product().UpdateProductHandler),
			),
		),
	)
	mux.Handle(
		"DELETE /api/v1/product/{id}",
		middleware.RequireAuth(cfg)(
			middleware.RequireAdmin(
				http.HandlerFunc(handlers.Product().DeleteProductHandler),
			),
		),
	)
	mux.Handle(
		"GET /api/v1/product/exists/{slug}",
		middleware.RequireAuth(cfg)(
			middleware.RequireAdmin(
				http.HandlerFunc(handlers.Product().ExistsProductHandler),
			),
		),
	)
	mux.Handle(
		"POST /api/v1/product/rating/{id}",
		middleware.RequireAuth(cfg)(
			http.HandlerFunc(handlers.ProductRating().CreateOrUpdateProductRatingHandler),
		),
	)
	mux.Handle(
		"DELETE /api/v1/product/rating/{id}",
		middleware.RequireAuth(cfg)(
			http.HandlerFunc(handlers.ProductRating().DeleteProductRatingHandler),
		),
	)
	mux.Handle(
		"POST /api/v1/product/image/{id}",
		middleware.RequireAuth(cfg)(
			middleware.RequireAdmin(
				http.HandlerFunc(handlers.ProductImage().CreateProductImageHandler),
			),
		),
	)
	mux.Handle(
		"POST /api/v1/product/comment/{id}",
		middleware.RequireAuth(cfg)(
			http.HandlerFunc(handlers.ProductComment().CreateProductCommentHandler),
		),
	)
	mux.Handle(
		"PUT /api/v1/product/comment/{id}",
		middleware.RequireAuth(cfg)(
			http.HandlerFunc(handlers.ProductComment().UpdateProductCommentHandler),
		),
	)
	mux.Handle(
		"DELETE /api/v1/product/comment/{id}",
		middleware.RequireAuth(cfg)(
			http.HandlerFunc(handlers.ProductComment().DeleteProductCommentHandler),
		),
	)
	mux.Handle(
		"POST /api/v1/product/comment/like/{id}",
		middleware.RequireAuth(cfg)(
			http.HandlerFunc(handlers.ProductCommentLike().CreateProductCommentLikeHandler),
		),
	)
	mux.Handle(
		"PUT /api/v1/product/comment/like/{id}",
		middleware.RequireAuth(cfg)(
			http.HandlerFunc(handlers.ProductCommentLike().UpdateProductCommentLikeHandler),
		),
	)
	mux.Handle(
		"DELETE /api/v1/product/comment/like/{id}",
		middleware.RequireAuth(cfg)(
			http.HandlerFunc(handlers.ProductCommentLike().DeleteProductCommentLikeHandler),
		),
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
