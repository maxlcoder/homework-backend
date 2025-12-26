package main

import (
	"fmt"

	"github.com/maxlcoder/homework-backend/model"
)

func main() {
	password := "Admin@123"
	fmt.Println(model.HashPassword(password))
}
