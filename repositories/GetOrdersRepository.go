package repositories

import (
	"fmt"
	"github.com/theWando/cornershop-orders/model"
	"github.com/theWando/cornershop-orders/rest/client"
)

const uri = "https://cornershopapp.com/api/v3/order_groups?page="

func GetOrders() (model.Result, error) {
	return GetPagedOrders(1)
}

type GetOrderRequest int

func GetPagedOrders(request GetOrderRequest) (model.Result, error) {
	empty := model.Result{}

	var result model.Result
	if err := client.Get(fmt.Sprint(uri, request), DefaultHeaders(), &result); err != nil {
		return empty, err
	}
	return result, nil
}
