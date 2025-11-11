package routes

import (
	"HelloWord/serverConfig/internals/handlers"
	"net/http"
)

func SetupHealthRoute(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("/health", handler.HealthHandler())
}
