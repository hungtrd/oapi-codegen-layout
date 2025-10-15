package handlers

import (
	"gorm.io/gorm"
	"oapi-codegen-layout/pkg/api"
)

// Handler implements the full ServerInterface by embedding specialized handlers
type Handler struct {
	*UserHandler
	*ProductHandler
}

// NewHandler creates a new unified handler
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		UserHandler:    NewUserHandler(db),
		ProductHandler: NewProductHandler(db),
	}
}

// Ensure Handler implements ServerInterface
var _ api.ServerInterface = (*Handler)(nil)
