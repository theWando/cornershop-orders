package usecase

import (
	"errors"
	"fmt"
	"github.com/theWando/cornershop-orders/model"
	"github.com/theWando/cornershop-orders/repositories"
	"sync"
)

func Get() (int, error) {
	orders, err := GetAllOrders()
	if err != nil {
		return 0, errors.New("error getting response")
	}

	var wg sync.WaitGroup
	wg.Add(len(orders))

	channel := make(chan float32, len(orders))

	var inter model.Order
	for _, inter = range orders {
		go func(order model.Order, channel chan float32) {
			defer wg.Done()

			orderDetail, err := repositories.GetOrder(order.Uuid)
			if err != nil {
				fmt.Printf("error getting data for %s. err %s\n", order.Uuid, err)
				wg.Done()
				return
			}

			if evalCriteria(orderDetail) {
				channel <- orderDetail.Total
			}
		}(inter, channel)
	}

	total := 0
	var value float32
	wg.Wait()
	close(channel)
	for value = range channel {
		total += int(value)
	}
	return total, nil
}

func evalCriteria(order model.OrderDetail) bool {
	orderBreakDown := order.Breakdown
	for _, breakdownItem := range orderBreakDown {
		if "payment_method" == breakdownItem.Type {
			if "Cobrado ···· 8398" == breakdownItem.Name {
				return true
			}
		}
	}
	return false
}
