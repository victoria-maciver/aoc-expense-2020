package main

import "testing"

func TestExpenseReportSolution(t *testing.T) {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	correct1, correct2 := 1721, 299

	a, b, err := expenseReportSolution(input)

	if err != nil {
		t.Errorf("Error encountered in expenseReportSolution")
	}

	if (a + b) != 2020 {
		t.Errorf("Does not sum to 2020")
	}

	if a != correct1 && a != correct2 {
		t.Errorf("Incorrect choice of %d, should be %d or %d", a, correct1, correct2)
	}
	if b != correct1 && b != correct2 {
		t.Errorf("Incorrect choice of %d, should be %d or %d", b, correct1, correct2)
	}
}
