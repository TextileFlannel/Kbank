package storage

type User struct {
	Login    string
	Password []byte
}

type Storage interface {
	CreateUser(login string, hashedPassword []byte) error
	GetUser(login string) (*User, error)
}