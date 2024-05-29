package usecase

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nutikuli/internProject_backend/internal/models/product"
	"github.com/nutikuli/internProject_backend/internal/models/product/dtos"
	_prodEntities "github.com/nutikuli/internProject_backend/internal/models/product/entities"
	"github.com/nutikuli/internProject_backend/internal/services/file"
	_fileEntities "github.com/nutikuli/internProject_backend/internal/services/file/entities"
)

type productUsecase struct {
	productRepo product.ProductRepository
	fileRepo    file.FileRepository
}

func NewProductUsecase(productRepo product.ProductRepository, fileRepo file.FileRepository) product.ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
		fileRepo:    fileRepo,
	}
}

// OnCreateProduct implements product.ProductUsecase.
func (p *productUsecase) OnCreateProduct(c *fiber.Ctx, ctx context.Context, productDatReq *_prodEntities.ProductCreateReq, fileDatReq []*_fileEntities.FileUploaderReq) (*dtos.ProductFileRes, int, error) {
	productId, err := p.productRepo.CreateProduct(ctx, *productDatReq)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	fEntity := &_fileEntities.FileEntityReq{
		EntityType: "PRODUCT",
		EntityId:   *productId,
	}

	for _, fDatReq := range fileDatReq {
		file := &_fileEntities.File{
			Type:       fDatReq.FileType,
			PathUrl:    fDatReq.FileData,
			Name:       fDatReq.FileName,
			EntityType: "PRODUCT",
			EntityId:   *productId,
		}

		_, fUrl, status, errOnCreatedFile := file.EncodeBase64toFile(c, true)
		if errOnCreatedFile != nil {
			return nil, status, errOnCreatedFile
		}

		fDatReq.FileData = *fUrl
		_, errOnInsertFile := p.fileRepo.CreateFileByEntityAndId(ctx, fDatReq, fEntity)
		if errOnInsertFile != nil {
			return nil, http.StatusInternalServerError, errOnInsertFile
		}

	}

	filesRes, errOnGetFiles := p.fileRepo.GetFilesByIdAndEntity(ctx, fEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	}

	newProduct, err := p.productRepo.GetProductById(ctx, productId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &dtos.ProductFileRes{
		Product: newProduct,
		Files:   filesRes,
	}, http.StatusCreated, nil
}

// OnGetProductById implements product.ProductUsecase.
func (p *productUsecase) OnGetProductById(ctx context.Context, productId int64) (*dtos.ProductFileRes, int, error) {
	product, err := p.productRepo.GetProductById(ctx, &productId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	fEntity := &_fileEntities.FileEntityReq{
		EntityType: "PRODUCT",
		EntityId:   productId,
	}

	files, err := p.fileRepo.GetFilesByIdAndEntity(ctx, fEntity)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &dtos.ProductFileRes{
		Product: product,
		Files:   files,
	}, http.StatusInternalServerError, nil
}

// OnGetProductsByStoreId implements product.ProductUsecase.
func (p *productUsecase) OnGetProductsByStoreId(ctx context.Context, storeId int64) ([]*dtos.ProductFileRes, int, error) {
	products, err := p.productRepo.GetProductsByStoreId(ctx, &storeId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var productsFileRes []*dtos.ProductFileRes
	for _, product := range products {
		fEntity := &_fileEntities.FileEntityReq{
			EntityType: "PRODUCT",
			EntityId:   product.Id,
		}

		files, err := p.fileRepo.GetFilesByIdAndEntity(ctx, fEntity)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		productsFileRes = append(productsFileRes, &dtos.ProductFileRes{
			Product: product,
			Files:   files,
		})
	}

	return productsFileRes, http.StatusInternalServerError, nil
}