package repositories

import (
	"fmt"
	"github.com/theWando/cornershop-orders/rest/client"
)

const url = "https://cornershopapp.com/api/v3/orders/"

func GetOrder(id string) (map[string]interface{}, error) {
	headers := DefaultHeaders()
	fullUrl := fmt.Sprint(url, id)
	return client.Get(fullUrl, headers)
}
