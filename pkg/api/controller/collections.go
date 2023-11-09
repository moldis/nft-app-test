package controller

import (
	"artemb/nft/pkg/api/response"
	"artemb/nft/pkg/db/model"
	repo "artemb/nft/pkg/db/repo"
	"go.uber.org/zap"
	"net/http"
)

type CollectionsController struct {
	Logger      *zap.Logger
	Collections *repo.CollectionsRepository
}

type CollectionsResponse struct {
	Collections []model.Collections `json:"collections"`
}

type MintedResponse struct {
	Minted []model.MintedCollections `json:"minted"`
}

func (c *CollectionsController) All(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := c.Collections.AllCollections(ctx)
	if err != nil {
		response.WriteJSONResponse(w, r, http.StatusBadRequest, response.ErrorResponse{Error: "unable to get data"})
		return
	}
	response.WriteJSONResponse(w, r, http.StatusOK, CollectionsResponse{Collections: result})
}

func (c *CollectionsController) AllMinted(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := c.Collections.AllMinted(ctx)
	if err != nil {
		response.WriteJSONResponse(w, r, http.StatusBadRequest, response.ErrorResponse{Error: "unable to get data"})
		return
	}
	response.WriteJSONResponse(w, r, http.StatusOK, MintedResponse{Minted: result})
}
