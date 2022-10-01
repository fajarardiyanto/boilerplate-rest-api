package service

import repo "github.com/fajarardiyanto/boilerplate-rest-api/pkg/repository/service"

type service struct{}

func NewHelloService() repo.HelloRepositoryService {
	return &service{}
}

func (*service) Hello(name string) string {
	return name
}
