package main

import (
	"fmt"
	"github.com/theWando/cornershop-orders/usecase"
)

func main() {
	fmt.Println("Starting Applications")
	total, err := usecase.Get()

	if err != nil {
		fmt.Println("error getting order ", err)
	}
	fmt.Println(total)
}
