package main

import (
	"fmt"

	"github.com/uzzal71/podcast/auth"
	"github.com/uzzal71/podcast/user"
)

func main() {
	auth.LoginWithCredentials("uzzal@gmail.com", "123456")
	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email: "sujon@gmail.com",
		Name:  "Sujon Roy",
	}

	fmt.Println(user.Email, user.Name)
}
