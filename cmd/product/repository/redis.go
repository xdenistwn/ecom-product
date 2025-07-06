package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"product/models"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	cacheKeyProductInfo         = "product:%d"
	cacheKeyProductCateogryInfo = "product_category:%d"
)

func (r *ProductRepository) GetProductByIDFromRedis(ctx context.Context, productID int64) (*models.Product, error) {
	cacheKey := fmt.Sprintf(cacheKeyProductInfo, productID)

	var product models.Product

	productStr, err := r.Redis.Get(ctx, cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			return &product, nil
		}

		return nil, err
	}

	// unmarshal redis string to model struct
	err = json.Unmarshal([]byte(productStr), &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) GetProductCategoryByIDFromRedis(ctx context.Context, productCategoryID int64) (*models.ProductCategory, error) {
	cacheKey := fmt.Sprintf(cacheKeyProductCateogryInfo, productCategoryID)

	var productCategory models.ProductCategory

	productCategoryStr, err := r.Redis.Get(ctx, cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, err
	}

	// unmarshal redis string to model struct
	err = json.Unmarshal([]byte(productCategoryStr), &productCategory)
	if err != nil {
		return nil, err
	}

	return &productCategory, nil
}

func (r *ProductRepository) SetProductByID(ctx context.Context, product *models.Product, productID int64) error {
	cacheKey := fmt.Sprintf(cacheKeyProductInfo, productID)

	productJSON, err := json.Marshal(product)

	// untuk debug
	// deadline, _ := ctx.Deadline()
	// time.Sleep(8 * time.Second)
	// fmt.Printf("context deadline %v \n", deadline)
	// fmt.Printf("context deadline Timerun %v \n", time.Until(deadline))

	if err != nil {
		return err
	}

	err = r.Redis.SetEx(ctx, cacheKey, productJSON, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) SetProductCategoryByID(ctx context.Context, productCategory *models.ProductCategory, productCategoryID int64) error {
	cacheKey := fmt.Sprintf(cacheKeyProductCateogryInfo, productCategoryID)

	productCategoryJSON, err := json.Marshal(productCategory)
	if err != nil {
		return err
	}

	err = r.Redis.SetEx(ctx, cacheKey, productCategoryJSON, 1*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}
