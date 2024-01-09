package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username    string
	Email       string `gorm:"type:varchar(100);unique_index"`
	Password    string `gorm:"type:varchar(100);unique_index"`
	FirstName   string `gorm:"type:varchar(100);unique_index"`
	LastName    string `gorm:"type:varchar(100);unique_index"`
	DateOfBirth time.Time

	ApiKeyID uuid.UUID `gorm:"type:uuid;uniqueIndex"`
	ApiKeys  []ApiKey  `gorm:"many2many:user_api_keys;"`

	JobsID uuid.UUID `gorm:"type:uuid;uniqueIndex"`
	Jobs   []Jobs    `gorm:"many2many:user_jobs;"`
}
