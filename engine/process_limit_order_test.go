package engine

import (
	"fmt"
	"testing"
)

func TestProcessLimitOrder(t *testing.T) {
	var tests = []struct {
		bookGen        []*Order
		input          *Order
		processedOrder []*Order
		partialOrder   *Order
	}{
		{
			[]*Order{
				NewOrder("b1", Buy, "5.0", "7000.0"),
			},
			NewOrder("s2", Sell, "5.0", "8000.0"),
			[]*Order{},
			nil,
		},
		{
			[]*Order{
				NewOrder("s2", Sell, "5.0", "8000.0"),
			},
			NewOrder("b1", Buy, "5.0", "7000.0"),
			[]*Order{},
			nil,
		},
		////////////////////////////////////////////////////////////////////////
		{
			[]*Order{
				NewOrder("b1", Buy, "5.0", "7000.0"),
			},
			NewOrder("s2", Sell, "5.0", "7000.0"),
			[]*Order{
				NewOrder("b1", Buy, "5.0", "7000.0"),
				NewOrder("s2", Sell, "5.0", "7000.0"),
			},
			nil,
		},
		{
			[]*Order{
				NewOrder("s1", Sell, "5.0", "7000.0"),
			},
			NewOrder("b2", Buy, "5.0", "7000.0"),
			[]*Order{
				NewOrder("s1", Sell, "5.0", "7000.0"),
				NewOrder("b2", Buy, "5.0", "7000.0"),
			},
			nil,
		},
		////////////////////////////////////////////////////////////////////////
		{
			[]*Order{
				NewOrder("b1", Buy, "5.0", "7000.0"),
			},
			NewOrder("s2", Sell, "1.0", "7000.0"),
			[]*Order{
				NewOrder("s2", Sell, "1.0", "7000.0"),
			},
			NewOrder("b1", Buy, "4.0", "7000.0"),
		},
		{
			[]*Order{
				NewOrder("s1", Sell, "5.0", "7000.0"),
			},
			NewOrder("b2", Buy, "1.0", "7000.0"),
			[]*Order{
				NewOrder("b2", Buy, "1.0", "7000.0"),
			},
			NewOrder("s1", Sell, "4.0", "7000.0"),
		},
		////////////////////////////////////////////////////////////////////////
		{
			[]*Order{
				NewOrder("b1", Buy, "5.0", "7000.0"),
			},
			NewOrder("s2", Sell, "1.0", "7000.0"),
			[]*Order{
				NewOrder("s2", Sell, "1.0", "7000.0"),
			},
			NewOrder("b1", Buy, "4.0", "7000.0"),
		},
	}

	for i, tt := range tests {
		ob := NewOrderBook()

		// Order book generation.
		for _, o := range tt.bookGen {
			ob.Process(*o)
		}

		processedOrder, partialOrder := ob.Process(*tt.input)
		fmt.Println("result ", i, processedOrder, partialOrder)
		for i, po := range processedOrder {
			if po.String() != tt.processedOrder[i].String() {
				fmt.Println(i, po, tt.processedOrder[i], len(po.String()), len(tt.processedOrder[i].String()))
				t.Fatalf("Incorrect processedOrder: (have: \n%s\n, want: \n%s\n)", processedOrder, tt.processedOrder)
			}
		}

		if tt.partialOrder == nil {
			if partialOrder != tt.partialOrder {
				fmt.Println(len(partialOrder.String()), len((tt.partialOrder.String())))
				t.Fatalf("Incorrect partialOrder: (have: \n%s\n, want: \n%s)", partialOrder, tt.partialOrder)
			}
		} else {
			if partialOrder.String() != tt.partialOrder.String() {
				// fmt.Println(len(partialOrder.String()), len((tt.partialOrder.String())))
				t.Fatalf("Incorrect partialOrder: (have: \n%s\n, want: \n%s)", partialOrder, tt.partialOrder)
			}
		}
	}
}