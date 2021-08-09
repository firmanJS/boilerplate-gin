package auth

import (
	"github.com/firmanJS/boilerplate-gin/model"
	"github.com/firmanJS/boilerplate-gin/util"
	"gorm.io/gorm"
)

type RepositoryRegister interface {
	RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repositoryRegister struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repositoryRegister {
	return &repositoryRegister{db: db}
}

func (r *repositoryRegister) RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	checkUserAccount := db.Debug().Select("*").Where("username = ?", input.Username).Find(&users)

	if checkUserAccount.RowsAffected > 0 {
		errorCode <- util.CONFLICT
		return &users, <-errorCode
	}

	users.Username = input.Username
	users.Password = input.Password

	addNewUser := db.Debug().Create(&users)

	db.Commit()

	if addNewUser.Error != nil {
		errorCode <- util.FAILED
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
