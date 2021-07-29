package updateProduct

import (
	model "github.com/firmanJS/boilerplate-gin/model"
	structs "github.com/firmanJS/boilerplate-gin/usecase/product"
	util "github.com/firmanJS/boilerplate-gin/util"
)

type Service interface {
	UpdateProductService(input *structs.InputUpdateProduct) (*model.EntityProduct, *util.CatchError)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateProductService(input *structs.InputUpdateProduct) (*model.EntityProduct, *util.CatchError) {

	Products := model.EntityProduct{
		Id:       input.Id,
		Name:     input.Name,
		Quantity: input.Quantity,
		Price:    input.Price,
	}

	resultUpdateProduct, errUpdateProduct := s.repository.UpdateProductRepository(&Products)

	return resultUpdateProduct, errUpdateProduct
}
