package main

import "testing"

func TestBinarySearch(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	testSearchFor(1, input, t)
	testSearchFor(5, input, t)
	testSearchFor(9, input, t)
}

func testSearchFor(value int, input []int, t *testing.T) {
	found := search(input, value)
	if !found {
		t.Errorf("Did not find %d", value)
	}
}

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

	a, b := expenseReportSolution(input)

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
