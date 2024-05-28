package Account

import (
	"context"

	"github.com/textures1245/go-template/internal/models/account/entities"
)

type AccountRepository interface {
	GetAccountCustomer(ctx context.Context) (*entities.Account, error)
	GetAccountStore(ctx context.Context) (*entities.Account, error)
	GetAccountAdmin(ctx context.Context) (*entities.Account, error)
}