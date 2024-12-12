package main

import (
	"fmt"
	"learn-go/1_basic/interface/services"
)

func main() {
	var db []*services.User

	userServices := services.NewUserServices(db)
	user1 := userServices.Register(&services.User{
		ID:   1,
		Name: "Zakaria",
	})
	fmt.Println(user1)
	user2 := userServices.Register(&services.User{
		ID:   2,
		Name: "Wahyu",
	})
	fmt.Println(user2)
	result := userServices.GetUser()
	fmt.Println("----HASIL GET USER---")
	for _, v := range result {
		fmt.Println(v.Name)
	}
	fmt.Println("----GET BY ID----")
	getIndex := userServices.GetUserByID(2)
	fmt.Println(getIndex)
}
