package readCategory

import (
	model "github.com/firmanJS/boilerplate-gin/model"
	util "github.com/firmanJS/boilerplate-gin/util"
)

type Service interface {
	ReadCategoryService() (*[]model.EntityCategory, *util.CatchError)
}

type service struct {
	repository Repository
}

func NewServiceRead(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ReadCategoryService() (*[]model.EntityCategory, *util.CatchError) {

	resultCategory, errCategory := s.repository.ReadCategoryRepository()

	return resultCategory, errCategory
}
