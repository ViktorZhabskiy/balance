package v1

import (
	"balance/internal/api/v1/handlers"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func NewV1Routes(handlers handlers.Handlers) *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(404), 404)
	})

	r.Route("/users", func(r chi.Router) {
		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/balance", handlers.BalanceHandler.GetUserBalance())
		})
	})

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		logrus.Fatal("Logging err: %s\n", err.Error())
	}

	return r
}
