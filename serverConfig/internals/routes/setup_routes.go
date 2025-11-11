package routes

import (
	"HelloWord/serverConfig/internals/handlers"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	SetupHealthRoute(mux, handler)
}
