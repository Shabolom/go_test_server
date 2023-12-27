package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;" json:"id"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}
