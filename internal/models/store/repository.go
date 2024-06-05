package store

import (
	"context"

	_accEntities "github.com/nutikuli/internProject_backend/internal/models/account/entities"
	"github.com/nutikuli/internProject_backend/internal/models/store/entities"
)

type StoreRepository interface {
	CreateStoreAccount(ctx context.Context, req entities.StoreRegisterReq) (*int64, error)
	GetStoreById(ctx context.Context, storeId int64) (*entities.Store, error)
	UpdateStoreAccountPassword(ctx context.Context, req _accEntities.UpdatePass) error
	UpdateStoreById(ctx context.Context, storeId int64, req entities.StoreUpdatedReq) error
	DeleteStoreById(ctx context.Context, storeId int64) error
	GetStoreAccounts(ctx context.Context) ([]*entities.Store, error)
}
