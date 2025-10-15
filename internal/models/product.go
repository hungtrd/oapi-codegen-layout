package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Name        string    `gorm:"type:varchar(200);not null"`
	Description *string   `gorm:"type:varchar(1000)"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	Category    string    `gorm:"type:varchar(100);not null;index"`
	Stock       int32     `gorm:"type:int;not null;default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate hook to generate UUID before creating
func (p *Product) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
