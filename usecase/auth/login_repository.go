package auth

import (
	"github.com/firmanJS/boilerplate-gin/model"
	"github.com/firmanJS/boilerplate-gin/util"
	"gorm.io/gorm"
)

type RepositoryLogin interface {
	LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repositoryLogin struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repositoryLogin {
	return &repositoryLogin{db: db}
}

func (r *repositoryLogin) LoginRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Username = input.Username
	users.Password = input.Password

	checkUserAccount := db.Debug().Select("*").Where("username = ?", input.Username).Find(&users)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- util.NOT_FOUND
		return &users, <-errorCode
	}

	comparePassword := util.ComparePassword(users.Password, input.Password)

	if comparePassword != nil {
		errorCode <- util.FAILED
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
