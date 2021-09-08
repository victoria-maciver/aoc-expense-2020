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
	correct1, correct2, correct3 := 979, 366, 675

	a, b, c, err := expenseReportSolution(input)

	if err != nil {
		t.Errorf("Error encountered in expenseReportSolution")
	}

	if (a + b + c) != 2020 {
		t.Errorf("Does not sum to 2020")
	}

	if a != correct1 && a != correct2 && a != correct3 {
		t.Errorf("Incorrect choice of %d, should be either: %d, %d or %d", a, correct1, correct2, correct3)
	}
	if b != correct1 && b != correct2 && b != correct3 {
		t.Errorf("Incorrect choice of %d, should be either %d, %d or %d", b, correct1, correct2, correct3)
	}
	if c != correct1 && c != correct2 && c != correct3 {
		t.Errorf("Incorrect choice of %d, should be either %d, %d or %d", c, correct1, correct2, correct3)
	}
}
