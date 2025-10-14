package handlers

import (
	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"oapi-codegen-layout/pkg/api"
)

// Handler implements the full ServerInterface by embedding specialized handlers
type Handler struct {
	*UserHandler
	*ProductHandler
}

// NewHandler creates a new unified handler
func NewHandler() *Handler {
	return &Handler{
		UserHandler:    NewUserHandler(),
		ProductHandler: NewProductHandler(),
	}
}

// Ensure Handler implements ServerInterface
var _ api.ServerInterface = (*Handler)(nil)

// GetHealth implements the health check endpoint (required by ServerInterface)
// (GET /health)
func (h *Handler) GetHealth(c *gin.Context) {
	h.UserHandler.GetHealth(c)
}

// ListUsers implements listing users (required by ServerInterface)
// (GET /users)
func (h *Handler) ListUsers(c *gin.Context, params api.ListUsersParams) {
	h.UserHandler.ListUsers(c, params)
}

// CreateUser implements creating a user (required by ServerInterface)
// (POST /users)
func (h *Handler) CreateUser(c *gin.Context) {
	h.UserHandler.CreateUser(c)
}

// GetUserById implements getting a user by ID (required by ServerInterface)
// (GET /users/{userId})
func (h *Handler) GetUserById(c *gin.Context, userId openapi_types.UUID) {
	h.UserHandler.GetUserById(c, userId)
}

// UpdateUser implements updating a user (required by ServerInterface)
// (PUT /users/{userId})
func (h *Handler) UpdateUser(c *gin.Context, userId openapi_types.UUID) {
	h.UserHandler.UpdateUser(c, userId)
}

// DeleteUser implements deleting a user (required by ServerInterface)
// (DELETE /users/{userId})
func (h *Handler) DeleteUser(c *gin.Context, userId openapi_types.UUID) {
	h.UserHandler.DeleteUser(c, userId)
}

// ListProducts implements listing products (required by ServerInterface)
// (GET /products)
func (h *Handler) ListProducts(c *gin.Context, params api.ListProductsParams) {
	h.ProductHandler.ListProducts(c, params)
}

// CreateProduct implements creating a product (required by ServerInterface)
// (POST /products)
func (h *Handler) CreateProduct(c *gin.Context) {
	h.ProductHandler.CreateProduct(c)
}

// GetProductById implements getting a product by ID (required by ServerInterface)
// (GET /products/{productId})
func (h *Handler) GetProductById(c *gin.Context, productId openapi_types.UUID) {
	h.ProductHandler.GetProductById(c, productId)
}

// UpdateProduct implements updating a product (required by ServerInterface)
// (PUT /products/{productId})
func (h *Handler) UpdateProduct(c *gin.Context, productId openapi_types.UUID) {
	h.ProductHandler.UpdateProduct(c, productId)
}

// DeleteProduct implements deleting a product (required by ServerInterface)
// (DELETE /products/{productId})
func (h *Handler) DeleteProduct(c *gin.Context, productId openapi_types.UUID) {
	h.ProductHandler.DeleteProduct(c, productId)
}
