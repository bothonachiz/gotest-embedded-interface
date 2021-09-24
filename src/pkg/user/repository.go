package user

type UserRepository interface {
	GetRepo(number int) int
}

type Repository struct {
	UserRepository
}

func NewRepository() *Repository {
	repository := &Repository{}
	repository.UserRepository = repository
	return repository
}

func (s *Repository) GetRepo(number int) int {
	return number % 2
}
