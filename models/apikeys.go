package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ApiKey struct {
	gorm.Model
	Key         string    `gorm:"index;type:varchar(255);not null;unique"`
	UserID      uuid.UUID `gorm:"index"`
	Expiration  time.Time `gorm:"index"`
	IsActive    bool      `gorm:"index"`
	UsageCount  int       `gorm:"index"`
	Description string    `gorm:"index;type:varchar(255);not null;unique"`

	Jobs []Jobs `gorm:"many2many:api_keys_jobs;"`
}
