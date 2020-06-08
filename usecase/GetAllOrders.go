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
	orders := &initialResult.Results[0].Orders
	if pages.Total == 1 {
		return *orders, nil
	}

	size := initialResult.Pages.Total - 1
	channel := make(chan []model.BaseOrder, size)

	var wg sync.WaitGroup
	wg.Add(int(size))
	for p := pages.Current + 1; p <= pages.Total; p++ {
		go func(page int, channel chan []model.BaseOrder) {
			defer wg.Done()
			pagedOrders, _ := repositories.GetPagedOrders(repositories.GetOrderRequest(page))
			channel <- pagedOrders.Results
		}(int(p), channel)
	}

	wg.Wait()
	close(channel)
	for orderCombo := range channel {
		for _, baseOrder := range orderCombo {
			*orders = append(*orders, baseOrder.Orders...)
		}
	}
	return *orders, nil
}
