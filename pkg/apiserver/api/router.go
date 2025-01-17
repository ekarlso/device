package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nais/device/pkg/apiserver/auth"
	"github.com/nais/device/pkg/apiserver/database"
	"github.com/nais/device/pkg/apiserver/jita"
	"github.com/nais/device/pkg/apiserver/middleware"
)

type Config struct {
	DB            database.APIServer
	Jita          jita.Client
	APIKeys       map[string]string
	Authenticator auth.Authenticator
}

func New(cfg Config) chi.Router {
	api := &api{db: cfg.DB, jita: cfg.Jita}
	authenticator := cfg.Authenticator

	latencyHistBuckets := []float64{.001, .005, .01, .025, .05, .1, .5, 1, 3, 5}
	prometheusMiddleware := middleware.PrometheusMiddleware("apiserver", latencyHistBuckets...)
	prometheusMiddleware.Initialize("/devices", http.MethodGet, http.StatusOK)

	r := chi.NewRouter()

	r.Use(prometheusMiddleware.Handler())

	r.Group(func(r chi.Router) {
		r.Use(authenticator.Validator())
		r.Get("/deviceconfig", api.deviceConfig)
	})

	r.Get("/login", authenticator.LoginHTTP)
	r.Get("/authurl", authenticator.AuthURL)

	return r
}
