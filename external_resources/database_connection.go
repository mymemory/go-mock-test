package external_resources

import model "github.com/mymemory/go-mock-testing/model"

type Database interface {
	GetUserById(id string) model.User
	AddUser(user *model.User)
	DeleteUser(user *model.User)
	UpdateUser(user *model.User)
}