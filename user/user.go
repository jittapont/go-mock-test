package user

type User struct {
	Name string
}

//go:generate mockgen -destination=../mocks/mock_repository.go -package=mocks github.com/jittapont/go-mock-test/user UserRepository
type UserRepository interface {
	InsertUser(user *User) error
	GetUsers() ([]*User, error)
}

func InsertUser(repo UserRepository, u *User) error {
	return repo.InsertUser(u)
}
