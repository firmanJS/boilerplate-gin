package updateProduct

import (
	model "github.com/firmanJS/boilerplate-gin/model"
	util "github.com/firmanJS/boilerplate-gin/util"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateProductRepository(input *model.EntityProduct) (*model.EntityProduct, *util.CatchError)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateProductRepository(input *model.EntityProduct) (*model.EntityProduct, *util.CatchError) {

	var products model.EntityProduct
	db := r.db.Model(&products)

	products.Id = input.Id

	checkProductId := db.Debug().Select("name").Where("id = ?", input.Id).Find(&products)

	if checkProductId.RowsAffected < 1 {
		return &products, &util.CatchError{
			Code:    util.NOT_FOUND,
			Message: "Product data is not exist or deleted",
		}
	}

	products.Name = input.Name
	products.Price = input.Price
	products.Quantity = input.Quantity

	updateProduct := db.Debug().Select("name", "price", "quantity", "updated_at").Where("id = ?", input.Id).Updates(products)

	if updateProduct.Error != nil {
		return &products, &util.CatchError{
			Code:    util.FAILED,
			Message: updateProduct.Error.Error(),
		}
	} else {
		return &products, &util.CatchError{
			Code:    "",
			Message: "Update Product data sucessfully",
		}
	}
}
