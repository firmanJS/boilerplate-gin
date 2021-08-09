package readCategory

import (
	model "github.com/firmanJS/boilerplate-gin/model"
	util "github.com/firmanJS/boilerplate-gin/util"
	"gorm.io/gorm"
)

type Repository interface {
	ReadCategoryRepository() (*[]model.EntityCategory, *util.CatchError)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRead(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ReadCategoryRepository() (*[]model.EntityCategory, *util.CatchError) {

	var Categorys []model.EntityCategory
	db := r.db.Model(&Categorys)

	ReadCategorys := db.Debug().Select("id, name, created_at, updated_at").Find(&Categorys)

	if ReadCategorys.Error != nil {
		return &Categorys, &util.CatchError{
			Code:    util.FAILED,
			Message: ReadCategorys.Error.Error(),
		}
	} else {
		return &Categorys, &util.CatchError{
			Code:    "",
			Message: "success",
		}
	}
}
