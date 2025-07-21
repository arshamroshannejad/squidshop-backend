package router

import (
	"github.com/arshamroshannejad/nuke"
	_ "github.com/arshamroshannejad/squidshop-backend/api"
	swagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

func SetupRoutes() http.Handler {
	r := nuke.NewRouter()
	r.Use(nuke.RecoverMiddleware)
	r.Use(nuke.TimeoutMiddleware(time.Second * 10))
	r.Use(nuke.HeartbeatMiddleware("/ping"))
	r.Handle("/docs/", swagger.Handler(
		swagger.URL("doc.json"),
		swagger.DeepLinking(true),
		swagger.DocExpansion("none"),
		swagger.DomID("swagger-ui"),
	))
	r.HandleFunc("GET /root", func(w http.ResponseWriter, r *http.Request) {
		nuke.WriteJSON(w, http.StatusOK, nuke.M{"message": "Hello World"})
	})
	return nuke.CorsMiddleware(nil)(r)
}
