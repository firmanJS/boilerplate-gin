package deleteProduct

import (
	model "github.com/firmanJS/boilerplate-gin/model"
	util "github.com/firmanJS/boilerplate-gin/util"
	"gorm.io/gorm"
)

type Repository interface {
	DeletedProductRepository(input *model.EntityProduct) (*model.EntityProduct, *util.CatchError)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDelete(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) DeletedProductRepository(input *model.EntityProduct) (*model.EntityProduct, *util.CatchError) {

	var products model.EntityProduct
	db := r.db.Model(&products)

	checkdProductId := db.Debug().Select("*").Where("id = ?", input.Id).Find(&products)

	if checkdProductId.RowsAffected < 1 {
		return &products, &util.CatchError{
			Code:    util.CONFLICT,
			Message: "Product data is not exist or deleted",
		}
	}

	deletedProductId := db.Debug().Select("name").Where("id = ?", input.Id).Find(&products).Delete(&products)

	if deletedProductId.Error != nil {
		return &products, &util.CatchError{
			Code:    util.FAILED,
			Message: deletedProductId.Error.Error(),
		}
	} else {
        return &products, &util.CatchError{
			Code:    "",
			Message: "Data Deleted",
		}
    }	
}
