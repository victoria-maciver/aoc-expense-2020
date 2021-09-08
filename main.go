package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input, err := readInputFromFile()
	if err != nil {
		fmt.Println(err.Error())
	}
	a, b := expenseReportSolution(input)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
	fmt.Printf("Answer: %v\n", a*b)
}

// O(n*log(n))
func expenseReportSolution(input []int) (int, int) {
	sort.Ints(input)

	for i := range input {
		required := 2020 - input[i]
		if search(input, required) {
			return input[i], required
		}
	}
	return 0, 0
}

func readInputFromFile() ([]int, error) {
	path := "/Users/vmaciver/codestuff/adventcodego/input"
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, i)
	}
	return lines, scanner.Err()
}

func search(input []int, x int) bool {
	return binarySearch(input, 0, len(input)-1, x)
}

func binarySearch(input []int, low int, high int, x int) bool {
	if low > high {
		return false
	}

	mid := (low + high) / 2

	if x == input[mid] {
		return true
	} else if x < input[mid] {
		return binarySearch(input, low, mid-1, x)
	} else {
		return binarySearch(input, mid+1, high, x)
	}
}
