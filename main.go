package main

import (
	"fmt"
	"github.com/theWando/conershopOrders/repositories"
)

func main() {
	fmt.Println("Starting Applications")
	total, err := repositories.GetOrder("CL-750625-9154211")

	if err != nil {
		fmt.Println("error getting order ", err)
	}
	fmt.Println(total)
}
