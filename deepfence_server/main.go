package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"

	"github.com/casbin/casbin/v2"
	"github.com/deepfence/ThreatMapper/deepfence_server/common"
	"github.com/deepfence/ThreatMapper/deepfence_server/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/swaggest/openapi-go/openapi3"
)

var (
	verbosity = flag.String("verbose", "info", "log level")
)

type Config struct {
	RedisEndpoint      string
	HttpListenEndpoint string
}

func main() {
	flag.Parse()

	config, err := initialize()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	log.Info().Msg("starting deepfence-server")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	router.SetupRoutes(r)

	httpServer := http.Server{Addr: config.HttpListenEndpoint, Handler: r}

	idleConnectionsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.Error().Msgf("http server shutdown error: %v", err)
		}
		close(idleConnectionsClosed)
	}()

	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Error().Msgf("http server ListenAndServe error: %v", err)
		return
	}

	<-idleConnectionsClosed

	log.Info().Msg("deepfence-server stopped")
}

func initialize() (Config, error) {
	// logger
	// Default log level
	switch *verbosity {
	case zerolog.LevelTraceValue:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case zerolog.LevelDebugValue:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case zerolog.LevelInfoValue:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case zerolog.LevelWarnValue:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case zerolog.LevelErrorValue:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case zerolog.LevelFatalValue:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	//redisEndpoint, has := os.LookupEnv("REDIS_ENDPOINT")
	//if !has {
	//	return Config{}, errors.New("REDIS_ENDPOINT undefined")
	//}

	httpListenEndpoint := os.Getenv("HTTP_LISTEN_ENDPOINT")
	if httpListenEndpoint == "" {
		httpListenEndpoint = "8080"
	}

	// JWT
	common.TokenAuth = jwtauth.New("HS256", uuid.New(), nil)

	var err error
	// authorization
	common.CasbinEnforcer, err = casbin.NewEnforcer("authorization/casbin_model.conf", "authorization/casbin_policy.csv")
	if err != nil {
		return Config{}, err
	}

	// OpenAPI generation
	description := "Deepfence ThreatMapper API Documentation"
	common.OpenAPI = &openapi3.Reflector{
		Spec: &openapi3.Spec{
			Openapi: "3.0.3",
			Info: openapi3.Info{
				Title:          "Deepfence ThreatMapper",
				Description:    &description,
				TermsOfService: nil,
				Contact:        nil,
				License:        nil,
				Version:        "2.0.0",
			},
			ExternalDocs:  nil,
			Servers:       nil,
			Security:      nil,
			Tags:          nil,
			Paths:         openapi3.Paths{},
			Components:    nil,
			MapOfAnything: nil,
		},
	}
	//schema, err := common.OpenAPI.Spec.MarshalYAML()
	return Config{
		RedisEndpoint:      "",
		HttpListenEndpoint: ":" + httpListenEndpoint,
	}, nil
}
