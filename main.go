package main

import "github.com/bothonachiz/gotest/src/pkg/user"

func main() {
	userRepo := user.NewRepository()
	userService := user.NewService(userRepo)

	userService.GetUserInfo()
}
