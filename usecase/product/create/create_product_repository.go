package createProduct

import (
	model "github.com/firmanJS/boilerplate-gin/model"
	util "github.com/firmanJS/boilerplate-gin/util"
	"gorm.io/gorm"
)

type Repository interface {
	CreateProductRepository(input *model.EntityProduct) (*model.EntityProduct, *util.CatchError)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateProductRepository(input *model.EntityProduct) (*model.EntityProduct, *util.CatchError) {

	var products model.EntityProduct
	db := r.db.Model(&products)

	checkProductExist := db.Debug().Select("id_product").Where("id_product = ?", input.Id_Product).Find(&products)

	if checkProductExist.RowsAffected > 0 {
		return &products, &util.CatchError{
			Code:    util.CONFLICT,
			Message: "ID Product already exist",
		}
	}

	products.Id_Product = input.Id_Product
	products.Name = input.Name
	products.Price = input.Price
	products.Quantity = input.Quantity

	addNewProduct := db.Debug().Create(&products)
	db.Commit()

	if addNewProduct.Error != nil {
		return &products, &util.CatchError{
			Code:    util.FAILED,
			Message: addNewProduct.Error.Error(),
		}
	}

	return &products, &util.CatchError{}
}
