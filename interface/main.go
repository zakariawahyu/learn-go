package main

import (
	"fmt"
	"learn-go/interface/services"
)

func main() {
	var db []*services.User

	userServices := services.NewUserServices(db)
	user1 := userServices.Register(&services.User{
		Name: "Zakaria",
	})
	fmt.Println(user1)
	user2 := userServices.Register(&services.User{
		Name: "Wahyu",
	})
	fmt.Println(user2)
	result := userServices.GetUser()
	fmt.Println("----HASIL GET USER---")
	for _, v := range result {
		fmt.Println(v.Name)
	}
}
