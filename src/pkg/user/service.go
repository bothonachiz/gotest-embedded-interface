package user

import "fmt"

type UserService interface {
	GetInfo() bool
	Login() bool
	GetUserInfo() bool
}

type Service struct {
	UserService
	Repo UserRepository
}

func NewService(repository UserRepository) *Service {
	service := &Service{
		Repo: repository,
	}
	service.UserService = service
	return service
}

func (s *Service) GetInfo() bool {
	fmt.Println("GetInfo: ", *s)
	s.UserService.Login()
	res2 := s.Repo.GetRepo(20)
	res3 := s.Repo.GetRepo(20)
	fmt.Println(res2, res3)
	return true
}

func (s *Service) Login() bool {
	fmt.Println("Login: ", *s)
	return true
}

func (s *Service) GetUserInfo() bool {
	fmt.Println("GetUserInfo: ", *s)
	s.UserService.GetInfo()
	return true
}
