package readProduct

import (
	model "github.com/firmanJS/boilerplate-gin/model"
	util "github.com/firmanJS/boilerplate-gin/util"
)

type Service interface {
	ReadProductService() (*[]model.EntityProduct, *util.CatchError)
}

type service struct {
	repository Repository
}

func NewServiceRead(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ReadProductService() (*[]model.EntityProduct, *util.CatchError) {

	resultProduct, errProduct := s.repository.ReadProductRepository()

	return resultProduct, errProduct
}