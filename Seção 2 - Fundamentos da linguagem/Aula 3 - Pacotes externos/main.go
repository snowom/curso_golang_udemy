package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	err := checkmail.ValidateFormat("teste@gmail.com")
	fmt.Println(err)
}
