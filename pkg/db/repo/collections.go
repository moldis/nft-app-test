package repository

import (
	"artemb/nft/pkg/db/model"
	"context"

	"gorm.io/gorm"
)

type CollectionsRepository struct {
	db *gorm.DB
}

func (r *CollectionsRepository) AllCollections(ctx context.Context) ([]model.Collections, error) {
	var beans []model.Collections

	err := r.db.WithContext(ctx).Model(&model.Collections{}).Find(&beans).Error
	if err != nil {
		return nil, err
	}

	return beans, nil
}

func (r *CollectionsRepository) AddCollection(ctx context.Context, address, name, symbol string) error {
	return r.db.WithContext(ctx).Create(&model.Collections{Address: address, Name: name, Symbol: symbol}).Error
}

func (r *CollectionsRepository) AllMinted(ctx context.Context) ([]model.MintedCollections, error) {
	var beans []model.MintedCollections

	err := r.db.WithContext(ctx).Model(&model.MintedCollections{}).Find(&beans).Error
	if err != nil {
		return nil, err
	}

	return beans, nil
}

func (r *CollectionsRepository) AddMinted(ctx context.Context, collection, recipient, tokenId, tokenUri string) error {
	return r.db.WithContext(ctx).Save(&model.MintedCollections{Collection: collection, Recipient: recipient, TokenID: tokenId, TokenURL: tokenUri}).Error
}

func NewCollectionsRepository(db *gorm.DB) CollectionsRepository {
	return CollectionsRepository{db: db}
}
