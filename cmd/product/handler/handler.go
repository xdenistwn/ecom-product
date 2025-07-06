package handler

import (
	"fmt"
	"net/http"
	"product/cmd/product/usecase"
	"product/infrastructure/log"
	"product/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: productUsecase,
	}
}

// handler product category
func (h *ProductHandler) GetProductCategoryByID(c *gin.Context) {
	productCategoryIDstr := c.Param("id")

	productCategoryID, err := strconv.ParseInt(productCategoryIDstr, 10, 64)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"productCategoryID": productCategoryID,
		}).Errorf("strconv.ParseInt got error %v", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid Product Category ID",
		})

		return
	}

	product, err := h.ProductUsecase.GetProductCategoryByID(c.Request.Context(), productCategoryID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"productCategoryID": productCategoryID,
		}).Errorf("h.ProductUsecase.GetProductCategoryByID %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "Success",
		"product_category": product,
	})
}

func (h *ProductHandler) ProductCategoryManagement(c *gin.Context) {
	var param models.ProductCategoryManagementParameter

	if err := c.ShouldBindJSON(&param); err != nil {
		log.Logger.Error(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid Input",
		})

		return
	}

	if param.Action == "" {
		log.Logger.Error("Missing required action parameter")

		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Missing required action parameter",
		})

		return
	}

	switch param.Action {
	case "add":
		ProductCategoryID, err := h.ProductUsecase.CreateNewProductCategory(c.Request.Context(), &param.ProductCategory)

		if err != nil {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Errorf("h.ProductUsecase.CreateNewProductCategory got error %v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully create new product category: %d", ProductCategoryID),
		})

		return
	case "edit":
		if param.ID == 0 {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Error("Invalid request - product is empty")

			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Invalid request",
			})

			return
		}

		ProductCategory, err := h.ProductUsecase.EditProductCategory(c.Request.Context(), &param.ProductCategory)

		if err != nil {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Errorf("h.ProductUsecase.EditProductCategory got error %v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":          "Successfully edit product category.",
			"product_category": ProductCategory,
		})

		return
	case "delete":
		if param.ID == 0 {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Error("Invalid request - product is empty")

			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Invalid request",
			})

			return
		}

		err := h.ProductUsecase.DeleteProductCategory(c.Request.Context(), param.ID)

		if err != nil {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Errorf("h.ProductUsecase.DeleteProductCategory got error %v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully delete product category ID %d.", param.ID),
		})

		return

	default:
		log.Logger.Error("Invalid Action")
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid Action",
		})

		return
	}
}

// handler product
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	productIDstr := c.Param("id")

	productID, err := strconv.ParseInt(productIDstr, 10, 64)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"productID": productID,
		}).Errorf("strconv.ParseInt got error %v", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid Product ID",
		})

		return
	}

	product, err := h.ProductUsecase.GetProductByID(c.Request.Context(), productID)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"productID": productID,
		}).Errorf("h.ProductUsecase.GetProductByID %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"product": product,
	})
}

func (h *ProductHandler) ProductManagement(c *gin.Context) {
	var param models.ProductManagementParameter

	if err := c.ShouldBindJSON(&param); err != nil {
		log.Logger.Error(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid Input",
		})

		return
	}

	if param.Action == "" {
		log.Logger.Error("Missing required action parameter")

		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Missing required action parameter",
		})

		return
	}

	switch param.Action {
	case "add":
		productID, err := h.ProductUsecase.CreateNewProduct(c.Request.Context(), &param.Product)

		if err != nil {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Errorf("h.ProductUsecase.CreateNewProduct got error %v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully create new product: %d", productID),
		})

		return
	case "edit":
		if param.ID == 0 {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Error("Invalid request - product id is empty")

			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Invalid request",
			})

			return
		}

		product, err := h.ProductUsecase.EditProduct(c.Request.Context(), &param.Product)

		if err != nil {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Errorf("h.ProductUsecase.EditProduct got error %v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully edit product.",
			"product": product,
		})

		return
	case "delete":
		if param.ID == 0 {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Error("Invalid request - product id is empty")

			c.JSON(http.StatusBadRequest, gin.H{
				"error_message": "Invalid request",
			})

			return
		}

		err := h.ProductUsecase.DeleteProduct(c.Request.Context(), param.ID)

		if err != nil {
			log.Logger.WithFields(logrus.Fields{
				"param": param,
			}).Errorf("h.ProductUsecase.DeleteProduct got error %v", err)

			c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully delete product ID %d.", param.ID),
		})

		return

	default:
		log.Logger.Error("Invalid Action")
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid Action",
		})

		return
	}
}

func (h *ProductHandler) SearchProduct(c *gin.Context) {
	name := c.Query("name")
	category := c.Query("category")

	minPrice, _ := strconv.ParseFloat(c.Query("min_price"), 64)
	maxPrice, _ := strconv.ParseFloat(c.Query("max_price"), 64)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	orderBy := c.Query("order_by")
	sort := c.Query("sort")

	param := &models.SearchProductParameter{
		Name:     name,
		Category: category,
		MinPrice: minPrice,
		MaxPrice: maxPrice,
		Page:     page,
		PageSize: pageSize,
		OrderBy:  orderBy,
		Sort:     sort,
	}

	products, totalCount, err := h.ProductUsecase.SearchProduct(c.Request.Context(), param)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{
			"param": param,
		}).Errorf("h.ProductUsecase.SearchProduct got error %v", err)

		c.JSON(http.StatusOK, gin.H{
			"error_message": err.Error(),
		})

		return
	}

	// next page url
	totalPages := (totalCount + pageSize - 1) / pageSize

	var nextPageUrl *string

	if page < totalPages {
		url := fmt.Sprintf("%s/v1/product/search?name%s&category=%s&min_price=%0.f&max_price=%0.f&page=%d&page_size=%d", c.Request.Host, name, category, minPrice, maxPrice, page+1, pageSize)

		nextPageUrl = &url
	}

	c.JSON(http.StatusOK, gin.H{
		"data": models.SearchProductResponse{
			Products:    products,
			Page:        page,
			PageSize:    pageSize,
			TotalCount:  totalCount,
			TotalPages:  totalPages,
			NextPageUrl: nextPageUrl,
		},
	})
}
