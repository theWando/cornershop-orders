package usecase

import (
	"github.com/theWando/cornershop-orders/model"
	"github.com/theWando/cornershop-orders/repositories"
	"sync"
)

func GetAllOrders() ([]model.Order, error) {
	empty := make([]model.Order, 0, 0)
	initialResult, ordersErr := repositories.GetOrders()
	if ordersErr != nil {
		return empty, ordersErr
	}

	pages := &initialResult.Pages
	if pages.Total == 1 {
		return initialResult.Results, nil
	}
	orders := &initialResult.Results

	channel := make(chan []model.Order, int(pages.Total))

	var wg sync.WaitGroup
	wg.Add(int(pages.Total))
	for p := pages.Current; p <= pages.Total; p++ {
		go func(page int, channel chan []model.Order) {
			defer wg.Done()
			pagedOrders, _ := repositories.GetPagedOrders(repositories.GetOrderRequest{Page: page})
			channel <- pagedOrders.Results
		}(int(p), channel)
	}

	wg.Wait()
	close(channel)
	for orderCombo := range channel {
		*orders = append(*orders, orderCombo...)
	}
	return *orders, nil
}
