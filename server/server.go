package server

import (
	"context"
	"net/http"

	"log"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/fx"

	"pluto/config"
)

type Server struct {
}

func NewMux(lc fx.Lifecycle, config *config.Config) *mux.Router {
	address := ":" + config.Server.Port.String()

	if config.Misc.Env == "dev" {
		address = "127.0.0.1:" + config.Server.Port.String()
	}

	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedOrigins: config.Cors.AllowedOrigins,
		AllowedHeaders: config.Cors.AllowedHeaders,
	})

	handler := c.Handler(router)

	srv := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 30 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			go func() {
				log.Printf("Pluto server start listen at %s\n", address)
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Pluto server listen error: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Pluto server shutdown...")
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatalf("Pluto server shutdown error: %s\n", err)
			}
			return nil
		},
	})
	return router
}
