package auth

import (
	"github.com/firmanJS/boilerplate-gin/model"
)

type ServiceLogin interface {
	LoginService(input *model.InputLogin) (*model.EntityUsers, string)
}

type serviceLogin struct {
	repositoryLogin RepositoryLogin
}

func NewServiceLogin(repository RepositoryLogin) *serviceLogin {
	return &serviceLogin{repositoryLogin: repository}
}

func (s *serviceLogin) LoginService(input *model.InputLogin) (*model.EntityUsers, string) {

	user := model.EntityUsers{
		Username: input.Username,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repositoryLogin.LoginRepository(&user)

	return resultLogin, errLogin
}
