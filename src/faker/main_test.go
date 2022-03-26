package faker

import (
	"testing"
)

func TestNumberBetween(t *testing.T) {
	from := 0
	to := 10
	quantity := 100000
	for i := 0; i < quantity; i++ {
		number := NumberBetween(from, to)
		if number < from || number > to {
			t.Errorf("Expected a number between \"%d\" and \"%d\" but was \"%d\"", from, to, number)
		}
	}
}

func TestNumberBetweenReturnsAlmostOnce(t *testing.T) {
	from := 0
	to := 10
	quantity := 100000
	appearances := make(map[int]int)
	for i := 0; i < quantity; i++ {
		number := NumberBetween(from, to)
		appearances[number]++
	}
	for i := 0; i <= to; i++ {
		if appearances[i] == 0 {
			t.Errorf("Expected \"%d\" to appear at least once", i)
		}
	}
}

func TestBoolean(t *testing.T) {
	percentage := 0.25
	quantity := 100000
	isTrue := 0.0
	isFalse := 0.0
	for i := 0; i < quantity; i++ {
		boolean := Boolean(percentage)
		if boolean {
			isTrue++
		} else {
			isFalse++
		}
	}

	maxAllowed := percentage + 0.01
	minAllowed := percentage - 0.01
	obtainerPercentage := isTrue / float64(quantity)
	if obtainerPercentage > maxAllowed || obtainerPercentage < minAllowed {
		t.Errorf("Expected at max \"%f\" and at least \"%f\" true values but wast \"%f\"", maxAllowed, minAllowed, obtainerPercentage)
	}
}
