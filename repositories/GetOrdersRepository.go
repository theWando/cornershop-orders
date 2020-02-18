package repositories

import "github.com/theWando/cornershop-orders/rest/client"

func GetOrders() (map[string]interface{}, error) {
	return client.Get("https://cornershopapp.com/api/v3/order_groups", DefaultHeaders())
}
