package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"oapi-codegen-layout/pkg/api"
)

// ProductHandler implements the product-related ServerInterface methods
type ProductHandler struct {
	// Add dependencies like database, logger, etc.
}

// NewProductHandler creates a new product handler
func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

// ListProducts returns a list of products
// (GET /products)
func (h *ProductHandler) ListProducts(c *gin.Context, params api.ListProductsParams) {
	// Mock data for demonstration
	products := []api.Product{
		{
			Id:          openapi_types.UUID(uuid.New()),
			Name:        "Laptop",
			Description: strPtr("High-performance laptop"),
			Price:       1299.99,
			Category:    "Electronics",
			Stock:       int32Ptr(10),
			CreatedAt:   time.Now().Add(-24 * time.Hour),
			UpdatedAt:   timePtr(time.Now()),
		},
		{
			Id:          openapi_types.UUID(uuid.New()),
			Name:        "Wireless Mouse",
			Description: strPtr("Ergonomic wireless mouse"),
			Price:       29.99,
			Category:    "Electronics",
			Stock:       int32Ptr(50),
			CreatedAt:   time.Now().Add(-48 * time.Hour),
			UpdatedAt:   timePtr(time.Now()),
		},
		{
			Id:          openapi_types.UUID(uuid.New()),
			Name:        "Office Chair",
			Description: strPtr("Comfortable office chair"),
			Price:       199.99,
			Category:    "Furniture",
			Stock:       int32Ptr(15),
			CreatedAt:   time.Now().Add(-72 * time.Hour),
			UpdatedAt:   timePtr(time.Now()),
		},
	}

	// Apply category filter if provided
	if params.Category != nil && *params.Category != "" {
		filtered := []api.Product{}
		for _, p := range products {
			if p.Category == *params.Category {
				filtered = append(filtered, p)
			}
		}
		products = filtered
	}

	// Apply limit if provided
	if params.Limit != nil && int(*params.Limit) < len(products) {
		products = products[:*params.Limit]
	}

	c.JSON(http.StatusOK, products)
}

// CreateProduct creates a new product
// (POST /products)
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req api.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.Error{
			Code:    "invalid_request",
			Message: err.Error(),
		})
		return
	}

	// Create product (mock implementation)
	product := api.Product{
		Id:          openapi_types.UUID(uuid.New()),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		Stock:       req.Stock,
		CreatedAt:   time.Now(),
		UpdatedAt:   timePtr(time.Now()),
	}

	c.JSON(http.StatusCreated, product)
}

// GetProductById retrieves a product by ID
// (GET /products/{productId})
func (h *ProductHandler) GetProductById(c *gin.Context, productId openapi_types.UUID) {
	// Mock product retrieval
	product := api.Product{
		Id:          productId,
		Name:        "Sample Product",
		Description: strPtr("This is a sample product"),
		Price:       99.99,
		Category:    "Electronics",
		Stock:       int32Ptr(25),
		CreatedAt:   time.Now().Add(-24 * time.Hour),
		UpdatedAt:   timePtr(time.Now()),
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct updates an existing product
// (PUT /products/{productId})
func (h *ProductHandler) UpdateProduct(c *gin.Context, productId openapi_types.UUID) {
	var req api.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.Error{
			Code:    "invalid_request",
			Message: err.Error(),
		})
		return
	}

	// Mock product update
	product := api.Product{
		Id:          productId,
		Name:        getOrDefault(req.Name, "Updated Product"),
		Description: req.Description,
		Price:       getOrDefaultFloat(req.Price, 0),
		Category:    getOrDefault(req.Category, "General"),
		Stock:       req.Stock,
		CreatedAt:   time.Now().Add(-24 * time.Hour),
		UpdatedAt:   timePtr(time.Now()),
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct deletes a product
// (DELETE /products/{productId})
func (h *ProductHandler) DeleteProduct(c *gin.Context, productId openapi_types.UUID) {
	// Mock product deletion - productId is already validated by the generated middleware
	c.Status(http.StatusNoContent)
}

// Helper functions
func strPtr(s string) *string {
	return &s
}

func int32Ptr(i int32) *int32 {
	return &i
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func getOrDefault(ptr *string, def string) string {
	if ptr != nil {
		return *ptr
	}
	return def
}

func getOrDefaultFloat(ptr *float64, def float64) float64 {
	if ptr != nil {
		return *ptr
	}
	return def
}
