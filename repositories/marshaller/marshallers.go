package marshaller

type Marshaller interface {
	Marshall(value map[string]interface{}) (interface{}, error)
}

func NewOrdersMarshaller() *ordersMarshaller {
	return &ordersMarshaller{}
}

func NewPagesMarshaller() *pagesMarshaller {
	return &pagesMarshaller{}
}

func NewResultsMarshaller() *resultsMarshaller {
	return &resultsMarshaller{}
}
