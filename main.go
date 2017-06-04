package main

import (
	"fmt"
	external_resources "github.com/mymemory/go-mock-testing/external_resources"
	model "github.com/mymemory/go-mock-testing/model"

	"github.com/mymemory/go-mock-testing/mock_resources"
)

func main() {
	//get_user(1)
	add_user()
	get_user("1")
}

func get_user(id string){
	var database external_resources.Database
	database = new(mock_resources.Database)
	user := database.GetUserById(id)
	fmt.Println("get user : ", user.Name)
}

func add_user(){
	var database external_resources.Database
	database = new(mock_resources.Database)
	user := new(model.User)
	user.Id = "1"
	user.Name = "Hello"
	user.Email = "me@me.com"

	database.AddUser(user)

	fmt.Println("add user : ", user.Name)
}
