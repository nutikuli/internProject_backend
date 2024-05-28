package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nutikuli/internProject_backend/internal/models/store"
	"github.com/nutikuli/internProject_backend/internal/models/store/entities"
	"github.com/nutikuli/internProject_backend/internal/models/store/repository/repository_query"
	"github.com/nutikuli/internProject_backend/pkg/utils"
)

type storeRepo struct {
	db *sqlx.DB
}

func NewStoreRepository(db *sqlx.DB) store.StoreRepository {
	return &storeRepo{
		db,
	}
}

func (s *storeRepo) CreateStoreAccount(ctx context.Context, req entities.StoreRegisterReq) (*int64, error) {
	var storeId *int64
	args := utils.Array{
		req.Email,
		req.Password,
		req.Name,
		req.Phone,
		req.Location,
		req.Status,
		"STORE",
		req.StoreName,
		req.StoreLocation,
	}

	err := s.db.GetContext(ctx, storeId, repository_query.InsertStoreAccount, args...)
	if err != nil {
		return nil, err
	}

	return storeId, nil
}

func (s *storeRepo) GetStoreById(ctx context.Context, storeId *int64) (*entities.Store, error) {
	var store *entities.Store
	err := s.db.GetContext(ctx, store, repository_query.GetStoreAccountById, storeId)
	if err != nil {
		return nil, err
	}

	return store, nil
}
