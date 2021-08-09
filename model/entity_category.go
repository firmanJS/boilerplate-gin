package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityCategory struct {
	Id        string    `gorm:"primaryKey;" json:"id"`
	Name      string    `gorm:"type:varchar(50);unique;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (entity *EntityCategory) BeforeCreate(db *gorm.DB) error {
	entity.Id = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityCategory) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
