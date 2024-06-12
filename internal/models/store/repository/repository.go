package repository

import (
	"context"

	"github.com/gofiber/fiber/v2/log"

	"github.com/jmoiron/sqlx"
	_accEntities "github.com/nutikuli/internProject_backend/internal/models/account/entities"
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

	res, err := s.db.ExecContext(ctx, repository_query.InsertStoreAccount, args...)
	if err != nil {
		return nil, err
	}

	storeId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &storeId, nil
}

func (s *storeRepo) GetStoreById(ctx context.Context, storeId int64) (*entities.Store, error) {
	store := &entities.Store{}
	log.Debug("store id", storeId)
	err := s.db.GetContext(ctx, store, repository_query.GetStoreAccountById, storeId)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *storeRepo) UpdateStoreAccountPassword(ctx context.Context, req _accEntities.UpdatePass) error {
	args := utils.Array{
		req.Password,
		req.Id,
		req.Role,
	}

	_, err := s.db.ExecContext(ctx, repository_query.UpdateStoreAccountPassword, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *storeRepo) UpdateStoreById(ctx context.Context, storeId int64, req entities.StoreUpdatedReq) error {
	args := utils.Array{
		req.Email,
		req.Name,
		req.Phone,
		req.Location,
		req.Status,
		req.StoreName,
		req.StoreLocation,
		storeId,
		"STORE",
	}

	_, err := s.db.ExecContext(ctx, repository_query.UpdateStoreById, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s *storeRepo) DeleteStoreById(ctx context.Context, storeId int64) error {
	_, err := s.db.ExecContext(ctx, repository_query.DeleteStoreById, storeId)
	if err != nil {
		return err
	}

	return nil
}

// GetStoreAccounts implements store.StoreRepository.
func (s *storeRepo) GetStoreAccounts(ctx context.Context) ([]*entities.Store, error) {
	var stores = make([]*entities.Store, 0)

	err := s.db.SelectContext(ctx, &stores, repository_query.GetStoreAccounts)
	if err != nil {
		return nil, err
	}

	return stores, nil
}
