package repositories

import (
	"fmt"
	"github.com/theWando/cornershop-orders/model"
	"github.com/theWando/cornershop-orders/repositories/marshaller"
	"github.com/theWando/cornershop-orders/rest/client"
)

func GetOrders() (model.Result, error) {
	return GetPagedOrders(GetOrderRequest{Page: 1})
}

type GetOrderRequest struct {
	Page int
}

func GetPagedOrders(request GetOrderRequest) (model.Result, error) {
	empty := model.Result{}

	results, err := client.Get(fmt.Sprint("https://cornershopapp.com/api/v3/order_groups?page=", request.Page), DefaultHeaders())
	if err != nil {
		return empty, err
	}
	marshalledOrders, marshallingErrors := marshallResponse(results)
	if marshallingErrors != nil {
		return empty, marshallingErrors
	}
	return marshalledOrders, err
}

func marshallResponse(results map[string]interface{}) (model.Result, error) {
	ordersMarshaller := marshaller.NewResultsMarshaller()
	return ordersMarshaller.Marshall(results)
}
