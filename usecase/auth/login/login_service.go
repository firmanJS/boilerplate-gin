package usecaseLogin

import (
	"github.com/firmanJS/boilerplate-gin/model"
	structs "github.com/firmanJS/boilerplate-gin/usecase/auth"
)

type Service interface {
	LoginService(input *structs.InputLogin) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginService(input *structs.InputLogin) (*model.EntityUsers, string) {

	user := model.EntityUsers{
		Username: input.Username,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&user)

	return resultLogin, errLogin
}
