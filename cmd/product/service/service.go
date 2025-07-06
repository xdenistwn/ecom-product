package service

import (
	"context"
	"fmt"
	"product/cmd/product/repository"
	"product/infrastructure/log"
	"product/models"
	"time"

	"github.com/sirupsen/logrus"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

// function contructor
func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

// service
func (s *ProductService) GetProductByID(ctx context.Context, productID int64) (*models.Product, error) {
	// get cache from redis
	product, err := s.ProductRepository.GetProductByIDFromRedis(ctx, productID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"productID": productID,
		}).Errorf("s.ProductRepository.GetProductByIDFromRedis got error %v", err)
	}

	if product.ID != 0 {
		return product, nil
	}

	// get from db
	product, err = s.ProductRepository.FindProductByID(ctx, productID)

	if err != nil {
		return nil, err
	}

	// go routine usecase here
	// so every time redis cache is missed, system need to refetch from database
	// after that set it into Redis with expiry date
	// but I want to set the redis cache in the background so user does not have to wait.
	go func(product *models.Product, productID int64) {
		ctxDetach, cancelRedis := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelRedis()
		err := s.ProductRepository.SetProductByID(ctxDetach, product, productID)
		if err != nil {
			log.Logger.WithFields(logrus.Fields{
				"productID": productID,
			}).Errorf("s.ProductRepository.SetProductByID got error %v", ctx.Err())
		}
		fmt.Println("ctxDetached after call:", ctx.Err())
	}(product, productID)

	return product, nil
}

func (s *ProductService) GetProductCategoryByID(ctx context.Context, productCategoryID int64) (*models.ProductCategory, error) {
	productCategory, err := s.ProductRepository.FindProductCategoryByID(ctx, productCategoryID)

	if err != nil {
		return nil, err
	}

	return productCategory, nil
}

func (s *ProductService) CreateNewProduct(ctx context.Context, param *models.Product) (int64, error) {
	productID, err := s.ProductRepository.InsertNewProduct(ctx, param)

	if err != nil {
		return 0, err
	}

	return productID, nil
}

func (s *ProductService) CreateNewProductCategory(ctx context.Context, param *models.ProductCategory) (int64, error) {
	productCategoryID, err := s.ProductRepository.InsertNewProductCategory(ctx, param)

	if err != nil {
		return 0, err
	}

	return productCategoryID, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, param *models.Product) (*models.Product, error) {
	product, err := s.ProductRepository.UpdateProduct(ctx, param)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) UpdateProductCategory(ctx context.Context, param *models.ProductCategory) (*models.ProductCategory, error) {
	productCategory, err := s.ProductRepository.UpdateProductCategory(ctx, param)

	if err != nil {
		return nil, err
	}

	return productCategory, nil
}

func (s *ProductService) DeleteProductByID(ctx context.Context, productID int64) error {
	err := s.ProductRepository.DeleteProduct(ctx, productID)

	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) DeleteProductCategoryByID(ctx context.Context, productCategoryID int64) error {
	err := s.ProductRepository.DeleteProductCategory(ctx, productCategoryID)

	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) SearchProduct(ctx context.Context, param *models.SearchProductParameter) ([]models.Product, int, error) {
	products, totalCount, err := s.ProductRepository.SearchProduct(ctx, param)
	if err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}
