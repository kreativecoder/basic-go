package integers

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(3, 7)
	expected := 10

	if sum != expected {
		t.Errorf("sum: '%d' expected: '%d'", sum, expected)

	}

}
