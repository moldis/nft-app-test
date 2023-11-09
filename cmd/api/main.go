package main

import (
	mw "artemb/nft/pkg/api/middleware"
	"artemb/nft/pkg/api/routes"
	"artemb/nft/pkg/config"
	db "artemb/nft/pkg/db"
	repo "artemb/nft/pkg/db/repo"
	"artemb/nft/pkg/logging"
	"artemb/nft/pkg/nft/adapter"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
)

const (
	cmdServer = "server"
)

func main() {
	app := kingpin.New("api", "API")
	configFile := app.
		Flag("config", "path to config file").
		Short('c').
		Required().
		PlaceHolder("./path/config.yaml").
		String()

	app.Command(cmdServer, "runs the API server")
	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	cfg := config.Read(*configFile)

	logger, undo := initLogger(cfg)
	defer undo()

	router, err := initRouter(cfg, logger)
	if err != nil {
		log.Fatalln(err)
	}

	switch command {
	case cmdServer:
		err = runServer(router, cfg)
		if err != nil {
			logger.Fatal("Server fatal error", zap.Error(err))
		}
	}
}

func initRouter(cfg *config.Config, logger *zap.Logger) (*chi.Mux, error) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(mw.Logger(logger))
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.StripSlashes)
	r.Use(middleware.SetHeader("Content-type", "application/json"))
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.Api.Cors.AllowedOrigins,
		AllowedMethods:   cfg.Api.Cors.AllowedMethods,
		AllowedHeaders:   cfg.Api.Cors.AllowedHeaders,
		ExposedHeaders:   cfg.Api.Cors.ExposedHeaders,
		AllowCredentials: cfg.Api.Cors.AllowCredentials,
		MaxAge:           cfg.Api.Cors.MaxAge,
	}))

	compressor := middleware.NewCompressor(5)
	r.Use(compressor.Handler)

	if err := routes.MakeRoutes(r, cfg, logger); err != nil {
		return nil, err
	}

	return r, nil
}

func initLogger(cfg *config.Config) (*zap.Logger, func()) {
	logger, err := logging.NewLogger(cfg.Logging)
	if err != nil {
		panic(fmt.Sprintf("Can't initialize logger: %s", err.Error()))
	}

	logger = logger.Named(cfg.AppName)
	undo := zap.ReplaceGlobals(logger)

	return logger, func() {
		undo()
		_ = logger.Sync()
	}
}

func runServer(r *chi.Mux, cfg *config.Config) error {
	// trying to run nft listener first
	// it will Fatal if anything happened
	gormDb, err := db.NewInMemoryDB(cfg.DBConfig)
	if err != nil {
		return err
	}
	collectionsRepo := repo.NewCollectionsRepository(gormDb)
	go func() {
		service := adapter.NewNFTEventService(&collectionsRepo, cfg)
		err := service.RunService()
		if err != nil {
			log.Println("error with ETH NFT service, quiting...")
			log.Fatalln(err)
		}
		log.Println("NFT service running.")
	}()
	log.Println(fmt.Sprintf("Starting server on port:%d", cfg.Api.Port))
	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Api.Port), r)
}
