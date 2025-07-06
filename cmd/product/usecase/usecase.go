package usecase

import (
	"context"
	"product/cmd/product/service"
	"product/infrastructure/log"
	"product/models"

	"github.com/sirupsen/logrus"
)

type ProductUsecase struct {
	ProductService service.ProductService
}

func NewProductUsecase(productService service.ProductService) *ProductUsecase {
	return &ProductUsecase{
		ProductService: productService,
	}
}

func (uc *ProductUsecase) GetProductByID(ctx context.Context, productID int64) (*models.Product, error) {
	product, err := uc.ProductService.GetProductByID(ctx, productID)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (uc *ProductUsecase) GetProductCategoryByID(ctx context.Context, productCategoryID int64) (*models.ProductCategory, error) {
	productCategory, err := uc.ProductService.GetProductCategoryByID(ctx, productCategoryID)

	if err != nil {
		return nil, err
	}

	return productCategory, nil
}

func (uc *ProductUsecase) CreateNewProduct(ctx context.Context, param *models.Product) (int64, error) {
	productID, err := uc.ProductService.CreateNewProduct(ctx, param)

	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"name":     param.Name,
			"category": param.CategoryID,
		}).Errorf("uc.ProductService.CreateNewProduct got error %v", err)

		return 0, err
	}

	return productID, nil
}

func (uc *ProductUsecase) CreateNewProductCategory(ctx context.Context, param *models.ProductCategory) (int64, error) {
	productCategoryID, err := uc.ProductService.CreateNewProductCategory(ctx, param)

	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"name": param.Name,
		}).Errorf("uc.ProductService.CreateNewProductCategory got error %v", err)

		return 0, err
	}

	return productCategoryID, nil
}

func (uc *ProductUsecase) EditProduct(ctx context.Context, param *models.Product) (*models.Product, error) {
	product, err := uc.ProductService.UpdateProduct(ctx, param)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (uc *ProductUsecase) EditProductCategory(ctx context.Context, param *models.ProductCategory) (*models.ProductCategory, error) {
	productCategory, err := uc.ProductService.UpdateProductCategory(ctx, param)

	if err != nil {
		return nil, err
	}

	return productCategory, nil
}

// nyoba return langsung
func (uc *ProductUsecase) DeleteProduct(ctx context.Context, productID int64) error {
	return uc.ProductService.DeleteProductByID(ctx, productID)
}

func (uc *ProductUsecase) DeleteProductCategory(ctx context.Context, productCategoryID int64) error {
	return uc.ProductService.DeleteProductCategoryByID(ctx, productCategoryID)
}

// search
func (s *ProductUsecase) SearchProduct(ctx context.Context, param *models.SearchProductParameter) ([]models.Product, int, error) {
	products, totalCount, err := s.ProductService.SearchProduct(ctx, param)
	if err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}
