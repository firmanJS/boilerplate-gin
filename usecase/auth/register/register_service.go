package usecaseRegister

import (
	model "github.com/firmanJS/boilerplate-gin/model"
	structs "github.com/firmanJS/boilerplate-gin/usecase/auth"
)

type Service interface {
	RegisterService(input *structs.InputRegister) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input *structs.InputRegister) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		Username: input.Username,
		Password: input.Password,
	}

	resultRegister, errRegister := s.repository.RegisterRepository(&users)

	return resultRegister, errRegister
}
