package marshaller

import "github.com/theWando/cornershop-orders/model"

type pagesMarshaller struct{}

func (pm *pagesMarshaller) Marshall(value map[string]interface{}) (model.Page, error) {
	iPage := value["pages"].(map[string]interface{})
	page := model.Page{
		Current: iPage["current"].(float64),
		Total:   iPage["total"].(float64),
		Size:    iPage["size"].(float64),
	}
	return page, nil
}
