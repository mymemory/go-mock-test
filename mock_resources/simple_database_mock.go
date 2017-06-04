package mock_resources

import (
	/*"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"*/
	model "github.com/mymemory/go-mock-testing/model"
)


type Database struct{

}

func (m *Database) GetUserById(id string) model.User {
	user := new(model.User)
	user.Id = id
	user.Name = "Test"
	user.Email = "me@me.com"
	return *user
}

func (m *Database) AddUser(user *model.User){

}

func (m *Database) DeleteUser(user *model.User){

}

func (m *Database) UpdateUser(user *model.User){

}
