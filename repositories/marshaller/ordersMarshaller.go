package marshaller

import "github.com/theWando/cornershop-orders/model"

type ordersMarshaller struct{}

func (om *ordersMarshaller) Marshall(value map[string]interface{}) ([]model.Order, error) {
	results := value["results"].([]interface{})
	orders := make([]model.Order, len(results))
	for i, anOrder := range results {
		orders[i].Uuid = om.getUuid(anOrder)
	}
	return orders, nil
}

func (om *ordersMarshaller) getUuid(anOrder interface{}) string {
	innerOrders := (anOrder.(map[string]interface{}))["orders"].([]interface{})
	inter := innerOrders[0].(map[string]interface{})
	return inter["uuid"].(string)
}
