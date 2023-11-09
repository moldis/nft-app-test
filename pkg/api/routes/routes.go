package routes

import (
	"artemb/nft/pkg/api/controller"
	"artemb/nft/pkg/config"
	db "artemb/nft/pkg/db"
	repo "artemb/nft/pkg/db/repo"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

const (
	baseRoute   = "/"
	collections = "/collections"
	minted      = "/minted"
)

type dependencies struct {
	logger *zap.Logger
	config *config.Config
}

func MakeRoutes(router chi.Router, cfg *config.Config, logger *zap.Logger) error {
	return makeV1Routes(router, cfg, logger)
}

func makeV1Routes(router chi.Router, cfg *config.Config, logger *zap.Logger) error {
	deps, err := makeDeps(cfg, logger)
	if err != nil {
		return err
	}

	controller, err := makeCollectionsController(deps)
	if err != nil {
		return err
	}

	router.
		Route(baseRoute, func(r chi.Router) {
			r.Route(collections, collectionsRoutes(controller))
		}).
		Route(baseRoute, func(r chi.Router) {
			r.Route(minted, mintedRoutes(controller))
		})

	return nil
}

func collectionsRoutes(ctrl *controller.CollectionsController) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get(baseRoute, ctrl.All)
	}
}

func mintedRoutes(ctrl *controller.CollectionsController) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get(baseRoute, ctrl.AllMinted)
	}
}

func makeCollectionsController(deps *dependencies) (*controller.CollectionsController, error) {
	gormDb, err := db.NewInMemoryDB(deps.config.DBConfig)
	if err != nil {
		return nil, err
	}
	collectionsRepo := repo.NewCollectionsRepository(gormDb)
	return &controller.CollectionsController{
		Logger:      deps.logger,
		Collections: &collectionsRepo,
	}, nil
}

func makeDeps(cfg *config.Config, logger *zap.Logger) (*dependencies, error) {
	return &dependencies{logger: logger, config: cfg}, nil
}
