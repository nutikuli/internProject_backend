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

func (p *productUsecase) OnGetAllProducts(ctx context.Context) ([]*dtos.ProductFileRes, int, error) {
	products, err := p.productRepo.GetAllProducts(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var productFileRes []*dtos.ProductFileRes

	for _, product := range products {
		fEntity := &_fileEntities.FileEntityReq{
			EntityType: "PRODUCT",
			EntityId:   product.Id, // Assuming product has an ID field
		}

		files, err := p.fileRepo.GetFilesByIdAndEntity(ctx, fEntity)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		productFileRes = append(productFileRes, &dtos.ProductFileRes{
			Product: product,
			Files:   files,
		})
	}

	return productFileRes, http.StatusOK, nil
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
	}, http.StatusOK, nil
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

// OnDeleteProductById implements product.ProductUsecase.
func (p *productUsecase) OnDeleteProductById(ctx context.Context, productId int64) (int, error) {
	err := p.productRepo.DeleteProductById(ctx, &productId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// OnGetProductsByOrderId implements product.ProductUsecase.
func (p *productUsecase) OnGetProductsByOrderId(ctx context.Context, orderId int64) ([]*dtos.ProductFileRes, int, error) {
	products, err := p.productRepo.GetProductsByOrderId(ctx, &orderId)
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

func (p *productUsecase) OnUpdateProductById(c *fiber.Ctx, ctx context.Context, productId int64, productDatReq *_prodEntities.ProductUpdateReq, fileDatReq []*_fileEntities.FileUploaderReq) (*dtos.ProductFileRes, int, error) {
	err := p.productRepo.UpdateProductById(ctx, productId, productDatReq)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	fEntity := &_fileEntities.FileEntityReq{
		EntityType: "PRODUCT",
		EntityId:   productId,
	}

	for _, fDatReq := range fileDatReq {
		file := &_fileEntities.File{
			Type:       fDatReq.FileType,
			PathUrl:    fDatReq.FileData,
			Name:       fDatReq.FileName,
			EntityType: "PRODUCT",
			EntityId:   productId,
		}

		_, fUrl, status, errOnCreatedFile := file.UpdateFile(c, true)
		if errOnCreatedFile != nil {
			return nil, status, errOnCreatedFile
		}

		fDatReq.FileData = *fUrl
		errOnInsertFile := p.fileRepo.UpdateFileByIdAndEntity(ctx, fEntity, fDatReq)
		if errOnInsertFile != nil {
			return nil, http.StatusInternalServerError, errOnInsertFile
		}

	}

	filesRes, errOnGetFiles := p.fileRepo.GetFilesByIdAndEntity(ctx, fEntity)
	if errOnGetFiles != nil {
		return nil, http.StatusInternalServerError, errOnGetFiles
	}

	newProduct, err := p.productRepo.GetProductById(ctx, &productId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &dtos.ProductFileRes{
		Product: newProduct,
		Files:   filesRes,
	}, http.StatusOK, nil
}
