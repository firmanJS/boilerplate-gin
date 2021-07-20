package readProduct

import (
	model "github.com/firmanJS/boilerplate-gin/model"
	util "github.com/firmanJS/boilerplate-gin/util"
	"gorm.io/gorm"
)

type Repository interface {
	ReadProductRepository() (*[]model.EntityProduct, *util.CatchError)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRead(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ReadProductRepository() (*[]model.EntityProduct, *util.CatchError) {

	var products []model.EntityProduct
	db := r.db.Model(&products)

	ReadProducts := db.Debug().Select("id, id_product, name, price, quantity, created_at, updated_at").Find(&products)

	if ReadProducts.Error != nil {
		return &products, &util.CatchError{
			Code:    util.FAILED,
			Message: ReadProducts.Error.Error(),
		}
	} else {
		return &products, &util.CatchError{
			Code:    "",
			Message: "Data Deleted",
		}
	}
}
