package marshaller

import (
	"errors"
	"github.com/theWando/cornershop-orders/model"
)

type resultsMarshaller struct{}

func (rm *resultsMarshaller) Marshall(value map[string]interface{}) (model.Result, error) {
	empty := model.Result{}
	pages, err := NewPagesMarshaller().Marshall(value)
	if err != nil {
		return empty, errors.New("error marshalling pages")
	}

	results, resultsErr := NewOrdersMarshaller().Marshall(value)
	if resultsErr != nil {
		return empty, errors.New("error marshalling results")
	}
	return model.Result{
		Pages:   pages,
		Count:   value["count"].(float64),
		Results: results,
	}, nil
}
