package usecase

import (
	"errors"
	"fmt"
	"github.com/theWando/cornershop-orders/repositories"
	"sync"
)

func Get() (int, error) {
	orders, err := repositories.GetOrders()
	if err != nil {
		return 0, errors.New("error getting orders")
	}

	ordersSlice := orders["results"].([]interface{})
	var wg sync.WaitGroup
	wg.Add(len(ordersSlice))

	channel := make(chan float64, len(ordersSlice))

	var inter interface{}
	var i int
	for i, inter = range ordersSlice {
		go func(order map[string]interface{}, i int, channel chan float64) {
			defer wg.Done()

			ordersDetails := order["orders"].([]interface{})
			id := ordersDetails[0].(map[string]interface{})["uuid"].(string)
			orderDetail, err := repositories.GetOrder(id)
			if err != nil {
				fmt.Printf("%d error getting data for %s. err %s\n", i, id, err)
				wg.Done()
				return
			}
			total := orderDetail["total"]
			channel <- total.(float64)
		}(inter.(map[string]interface{}), i, channel)
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
