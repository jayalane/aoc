// -*- tab-width: 2 -*-

package main

import (
	"fmt"
	"strings"
)

const data0 = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

var period = "."

const (
	five = 5
	ten  = 10
)

type number struct {
	value    int64
	startRow int64
	startCol int64
	endRow   int64
	endCol   int64
}

var digits = []string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
}

func isSymbol(ch string) bool {
	if ch != period {
		isDigit, _ := isDigitValue(ch)
		if !isDigit {
			return true
		}
	}

	return false
}

func anySymbols(i int64, j int64, d [][]string) bool {
	if i > 0 && j > 0 && i < int64(len(d)) && j < int64(len(d[i])) {
		ch := d[i][j]
		if isSymbol(ch) {
			return true
		}
	}

	return false // outof bounds is false
}

// checkNumbers returns true if the number is next to a non-period
// for now assuming in a row.
func checkNumbers(n number, d [][]string) bool {
	fmt.Println("Checking", n)

	// check corners
	if anySymbols(n.startRow-1, n.startCol-1, d) {
		return true
	}

	if anySymbols(n.startRow+1, n.startCol-1, d) {
		return true
	}

	if anySymbols(n.endRow-1, n.endCol+1, d) {
		return true
	}

	if anySymbols(n.endRow+1, n.endCol+1, d) {
		return true
	}

	// check ends
	if anySymbols(n.startRow, n.startCol-1, d) {
		return true
	}

	fmt.Println("Checking", n.startRow, n.endCol, d[n.startRow][n.endCol+1])
	if anySymbols(n.startRow, n.endCol+1, d) {
		return true
	}

	// check middle
	for i := n.startCol; i < n.endCol+1; i++ {
		if anySymbols(n.startRow-1, i, d) {
			return true
		}

		if anySymbols(n.startRow+1, i, d) {
			return true
		}
	}

	return false
}

func isDigitValue(d string) (bool, int64) {
	if len(d) > 1 {
		return false, -1
	}

	for i, ch := range digits {
		if ch == d {
			return true, int64(i)
		}
	}

	return false, -1
}

/*
func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
*/

func getNums(l string, row int64, d *[][]string) ([]number, error) { //nolint:unparam
	// first digit
	totals := make([]number, 0, five)

	startCol := int64(-1)
	total := int64(0)

	for i := int64(0); i < int64(len(l)); i++ {
		ch := string(l[i])
		(*d)[row][i] = ch // save the char for detection

		isDigit, value := isDigitValue(ch)
		if isDigit {
			if startCol == -1 {
				startCol = i
			}

			total = total*ten + value
		} else if startCol >= 0 {
			number := number{total, row, startCol, row, i - 1}
			totals = append(totals, number)
			startCol = -1
			total = 0
		}
	}

	return totals, nil
}

func doLines(text string) error {
	lines := strings.Split(text, "\n")
	numbers := make([]number, 0, five)

	data := make([][]string, 0)

	for row, l := range lines {
		data = append(data, make([]string, len(l)))

		someNumbers, err := getNums(l, int64(row), &data)
		if err != nil {
			fmt.Println("Got error", err)

			return err
		}

		numbers = append(numbers, someNumbers...)
	}

	gTot := int64(0)

	for _, n := range numbers {
		if checkNumbers(n, data) {
			fmt.Println("Addiing one", n.value)
			gTot += n.value
		} else {
			fmt.Println("Skipping on", n.value)
		}
	}

	fmt.Println("Total:", gTot)
	fmt.Println("Got numbers", numbers)
	fmt.Println("Got data", data)

	return nil
}

func main() {
	err := doLines(data0)
	if err != nil {
		fmt.Println("data0", data0, "got err", err)
	}
}
