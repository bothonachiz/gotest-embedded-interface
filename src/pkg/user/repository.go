package user

type UserRepository interface {
	GetRepo(number int) int
}

type Repository struct {
	UserRepository
}

func NewRepository() *Repository {
	return &Repository{
		UserRepository: &Repository{},
	}
}

func (s *Repository) GetRepo(number int) int {
	return number % 2
}
