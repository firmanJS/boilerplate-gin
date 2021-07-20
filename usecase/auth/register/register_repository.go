package usecaseRegister

import (
	model "github.com/firmanJS/boilerplate-gin/model"
	util "github.com/firmanJS/boilerplate-gin/util"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

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
