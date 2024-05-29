package product

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/product/dtos"
	entities "github.com/nutikuli/internProject_backend/internal/models/product/entities"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type ProductUsecase interface {
	OnGetProductsByStoreId(ctx context.Context, storeId int64) ([]*dtos.ProductFileRes, int, error)
	OnGetProductById(ctx context.Context, productId int64) (*dtos.ProductFileRes, int, error)
	OnCreateProduct(c *fiber.Ctx, ctx context.Context, productDatReq *entities.ProductCreateReq, fileDatReq []*_fileEntities.FileUploaderReq) (*dtos.ProductFileRes, int, error)
}
