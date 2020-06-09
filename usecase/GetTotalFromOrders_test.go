package usecase

import (
	"github.com/theWando/cornershop-orders/model"
	"testing"
)

func Benchmark_evalCriteria(b *testing.B) {
	tests := model.OrderDetail{Breakdown: []model.Breakdown{
		{Type: "subtotal", Name: "Lider"},
		{Type: "subscription", Name: "Envío"},
		{Type: "-", Name: "separator"},
		{Type: "store", Name: "Total del pedido"},
		{Type: "payment_method", Name: "Cobrado ···· 8394"},
		{Type: "line", Name: "Cobrado ···· 8398"}}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		evalCriteria(tests)
	}

}

func Test_evalCriteria(t *testing.T) {
	type args struct {
		order model.OrderDetail
	}
	tests := []struct {
		name string
		args args
		want bool
	}{{name: "Should be true", args: struct{ order model.OrderDetail }{order: model.OrderDetail{
		Breakdown: []model.Breakdown{
			{Type: "payment_method", Name: "Cobrado ···· 8398"},
			{Type: "line", Name: "Cobrado ···· 8398"},
		},
	}}, want: true}, {name: "Should be false", args: struct{ order model.OrderDetail }{order: model.OrderDetail{
		Breakdown: []model.Breakdown{
			{Type: "payment_method", Name: "Cobrado ···· 8394"},
			{Type: "payment_method", Name: "Cobrdo ···· 8398"},
			{Type: "line", Name: "Cobrado ···· 8398"}},
	}}, want: false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalCriteria(tt.args.order); got != tt.want {
				t.Errorf("evalCriteria() = %v, want %v", got, tt.want)
			}
		})
	}
}
