package createProduct

import (
	model "github.com/firmanJS/boilerplate-gin/model"
	structs "github.com/firmanJS/boilerplate-gin/usecase/product"
	utils "github.com/firmanJS/boilerplate-gin/util"
)

type Service interface {
	CreateProductService(input *structs.InputCreateProduct) (*model.EntityProduct, *utils.CatchError)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateProductService(input *structs.InputCreateProduct) (*model.EntityProduct, *utils.CatchError) {

	products := model.EntityProduct{
		Id_Product: input.Id_Product,
		Name:       input.Name,
		Price:      input.Price,
		Quantity:   input.Quantity,
	}

	resultCreateProduct, errCreateProduct := s.repository.CreateProductRepository(&products)

	return resultCreateProduct, errCreateProduct
}
