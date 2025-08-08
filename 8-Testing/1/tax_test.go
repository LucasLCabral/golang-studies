package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected: %f | Received: %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount float64
		expect float64
	}

	table := []calcTax {
		{500.0, 5.0},
		{1000.0, 10.0},
		{10.0, 5.0},
		{0, 0},
	}

	for _, test := range table {
		result := CalculateTax(test.amount)
		if result != test.expect {
			t.Errorf("Expected: %f | Received: %f", test.expect, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}