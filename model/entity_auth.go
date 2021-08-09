package model

import (
	"time"

	util "github.com/firmanJS/boilerplate-gin/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityUsers struct {
	ID        string `gorm:"primaryKey;"`
	Username  string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InputLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type InputRegister struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,gte=8"`
}

func (entity *EntityUsers) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.Password = util.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityUsers) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
