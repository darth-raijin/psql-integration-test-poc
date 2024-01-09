package models

import (
	"gorm.io/gorm"
	"time"
)

type Jobs struct {
	gorm.Model
	JobID         uint      `gorm:"primaryKey;autoIncrement"`
	Name          string    `gorm:"index;type:varchar(255);not null"`
	Status        string    `gorm:"index;type:varchar(100);not null"`
	Schedule      string    `gorm:"type:varchar(100)"`
	LastRun       time.Time `gorm:"index"`
	NextRun       time.Time `gorm:"index"`
	MaxRetries    int       `gorm:"default:3"`
	RetryCount    int       `gorm:"default:0"`
	IsRecurring   bool      `gorm:"index"`
	Handler       string    `gorm:"type:text"`
	ExecutionTime int       `gorm:"index"`
	FailureReason string    `gorm:"type:text"`

	ApiKeys []ApiKey `gorm:"many2many:api_keys_jobs;"`
}
