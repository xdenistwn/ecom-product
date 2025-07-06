package models

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  int     `json:"category_id"`
}

type ProductManagementParameter struct {
	Action string `json:"action"`
	Product
}

type ProductCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ProductCategoryManagementParameter struct {
	Action string `json:"action"`
	ProductCategory
}

type SearchProductParameter struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
	OrderBy  string  `json:"order_by"`
	Sort     string  `json:"sort"`
}

type SearchProductResponse struct {
	Products    []Product `json:"products"`
	Page        int       `json:"page"`
	PageSize    int       `json:"page_size"`
	TotalCount  int       `json:"total_count"`
	TotalPages  int       `json:"total_pages"`
	NextPageUrl *string   `json:"next_page_url"`
}
