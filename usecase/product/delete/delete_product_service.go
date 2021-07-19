package deleteProduct

import (
	model "github.com/firmanJS/boilerplate-gin/model"
	structs "github.com/firmanJS/boilerplate-gin/usecase/product"
	util "github.com/firmanJS/boilerplate-gin/util"
)

type Service interface {
	DeleteProductService(input *structs.InputDeleteProduct) (*model.EntityProduct, *util.CatchError)
}

type service struct {
	repository Repository
}

func NewServiceDelete(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) DeleteProductService(input *structs.InputDeleteProduct) (*model.EntityProduct, *util.CatchError) {

	Products := model.EntityProduct{
		Id: input.Id,
	}

	resultDeleteProduct, errDeleteProduct := s.repository.DeletedProductRepository(&Products)

	return resultDeleteProduct, errDeleteProduct
}
