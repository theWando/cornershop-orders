package usecase

import (
	"errors"
	"fmt"
	"github.com/theWando/cornershop-orders/model"
	"github.com/theWando/cornershop-orders/repositories"
	"sync"
)

func Get() (int, error) {
	response, err := repositories.GetOrders()
	if err != nil {
		return 0, errors.New("error getting response")
	}

	var wg sync.WaitGroup
	orders := &response.Results
	wg.Add(len(*orders))

	channel := make(chan float64, len(*orders))

	var inter model.Order
	var i int
	for i, inter = range *orders {
		go func(order model.Order, i int, channel chan float64) {
			defer wg.Done()

			orderDetail, err := repositories.GetOrder(order.Uuid)
			if err != nil {
				fmt.Printf("%d error getting data for %s. err %s\n", i, order.Uuid, err)
				wg.Done()
				return
			}
			total := orderDetail["total"]
			channel <- total.(float64)
		}(inter, i, channel)
	}

	total := 0
	var value float64
	wg.Wait()
	close(channel)
	for value = range channel {
		total += int(value)
	}
	return total, nil
}
