package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"created_at" gorm:"autoCreateTime"`
	Completed   bool      `json:"completed" gorm:"default:false"`
}

// BeforeCreate GORM hook to set UUID before inserting a new record
func (t *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}
