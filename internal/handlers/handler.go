package handlers

import (
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
