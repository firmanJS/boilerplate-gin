package usecaseLogin

import (
	"github.com/firmanJS/boilerplate-gin/model"
)

type Service interface {
	LoginService(input *model.InputLogin) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginService(input *model.InputLogin) (*model.EntityUsers, string) {

	user := model.EntityUsers{
		Username: input.Username,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.LoginRepository(&user)

	return resultLogin, errLogin
}
