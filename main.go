package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := readInputFromFile()
	if err != nil {
		fmt.Println("Input file could not be read")
		os.Exit(1)
	}
	a, b, err := expenseReportSolution(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
	fmt.Printf("Answer: %v\n", a*b)
}

func expenseReportSolution(input []int) (int, int, error) {
	expenses := generateMap(input)

	for i := range input {
		required := 2020 - input[i]
		_, found := expenses[required]
		if found {
			return input[i], required, nil
		}
	}
	return 0, 0, fmt.Errorf("valid expenses could not be found")
}

func generateMap(input []int) map[int]int {
	gen := make(map[int]int)

	for i := range input {
		gen[input[i]] = input[i]
	}

	return gen
}

// // O(n*log(n))
// func expenseReportSolution(input []int) (int, int) {
// 	sort.Ints(input)

// 	for i := range input {
// 		required := 2020 - input[i]
// 		if search(input, required) {
// 			return input[i], required
// 		}
// 	}
// 	return 0, 0
// }

// func search(input []int, x int) bool {
// 	return binarySearch(input, 0, len(input)-1, x)
// }

// func binarySearch(input []int, low int, high int, x int) bool {
// 	if low > high {
// 		return false
// 	}

// 	mid := (low + high) / 2

// 	if x == input[mid] {
// 		return true
// 	} else if x < input[mid] {
// 		return binarySearch(input, low, mid-1, x)
// 	} else {
// 		return binarySearch(input, mid+1, high, x)
// 	}
// }

func readInputFromFile() ([]int, error) {
	path := "input"
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
	return lines, nil
}
