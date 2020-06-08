package repositories

import (
	"fmt"
	"github.com/theWando/cornershop-orders/model"
	"github.com/theWando/cornershop-orders/rest/client"
)

const url = "https://cornershopapp.com/api/v3/orders/"

func GetOrder(id string) (model.OrderDetail, error) {
	headers := DefaultHeaders()
	fullUrl := fmt.Sprint(url, id)
	var order model.OrderDetail
	if err := client.Get(fullUrl, headers, &order); err != nil {
		return model.OrderDetail{}, err
	}
	return order, nil
}
