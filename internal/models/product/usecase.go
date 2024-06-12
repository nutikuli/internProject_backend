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
	OnGetProductsByOrderId(ctx context.Context, orderId int64) ([]*dtos.ProductWithOrderProductFileRes, int, error)
	OnDeleteProductById(ctx context.Context, productId int64) (int, error)
	OnUpdateProductById(c *fiber.Ctx, ctx context.Context, productId int64, productDatReq *entities.ProductUpdateReq, fileDatReq []*_fileEntities.FileUploaderReq) (*dtos.ProductFileRes, int, error)
	OnGetAllProducts(ctx context.Context) ([]*dtos.ProductFileRes, int, error)
}
