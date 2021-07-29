package auth

import (
	"github.com/firmanJS/boilerplate-gin/model"
)

type ServiceRegister interface {
	RegisterService(input *model.InputRegister) (*model.EntityUsers, string)
}

type serviceRegister struct {
	repositoryRegister RepositoryRegister
}

func NewServiceRegister(repository RepositoryRegister) *serviceRegister {
	return &serviceRegister{repositoryRegister: repository}
}

func (s *serviceRegister) RegisterService(input *model.InputRegister) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		Username: input.Username,
		Password: input.Password,
	}

	resultRegister, errRegister := s.repositoryRegister.RegisterRepository(&users)

	return resultRegister, errRegister
}
