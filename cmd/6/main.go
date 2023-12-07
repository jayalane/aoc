// -*- tab-width: 2 -*-

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	two = 2
)

var fields = map[string]int{
	"Time:":     0,
	"Distance:": 1,
}

type race struct {
	fields [two]uint64
}

const data0 = `Time:      7  15   30
Distance:  9  40  200`

const data1 = `Time:        49     87     78     95
Distance:   356   1378   1502   1882`

var errParseLine = errors.New("parse line failed")

func newParseLineError(msg string) error {
	return fmt.Errorf("%w: %s", errParseLine, msg)
}

func handleLine(line string, races *[]race, skipSpaces bool) error {
	theFieldStr := ""

	for k := range fields {
		if strings.HasPrefix(line, k) {
			theFieldStr = k
		}
	}

	if theFieldStr == "" {
		return newParseLineError("field not found for line")
	}

	segs := strings.Split(line, theFieldStr)
	if len(segs) != two {
		return newParseLineError("Couldn't work out parse" + line)
	}

	numbers := segs[1]
	j := 0

	if skipSpaces {
		numbers = strings.ReplaceAll(numbers, " ", "")
	}
	
	for i, n := range strings.Split(numbers, " ") {
		if skipSpaces && i > 1 {
			return newParseLineError("SkipSpaces has 2 races" + line)
		}
		if n == "" {
			continue
		}

		n = strings.Trim(n, " ")

		N, err := strconv.ParseUint(n, 10, 64)
		if err != nil {
			return err
		}

		if len((*races)) < j+1 {
			(*races) = append((*races), race{})
		}
		fmt.Println("Races len, j", len(*races), j)
		(*races)[j].fields[fields[theFieldStr]] = N

		fmt.Println("setting", j, theFieldStr, N)
		j++
	}

	return nil
}

func countOptions(r race) uint64 {
	t := r.fields[fields["Time:"]]
	d := r.fields[fields["Distance:"]]

	numWins := uint64(0)

	for i := uint64(1); i < t; i++ {
		travelled := i * (t - i)
		if travelled > d {
			numWins++
		}
	}

	fmt.Println("For d, t", d, t, "NumWins is", numWins)

	return numWins
}

func analyzeRaces(races []race) uint64 {
	fmt.Println("Races are")

	product := uint64(1)
	for _, r := range races {
		numOpts := countOptions(r)
		product *= numOpts
		fmt.Println("Races are", numOpts)
	}

	return product
}

func doLines(text string, skipSpaces bool) error {
	lines := strings.Split(text, "\n")

	races := make([]race, 0, len(lines)>>1)

	for _, l := range lines {
		err := handleLine(l, &races, skipSpaces)
		if err != nil {
			return err
		}
	}

	total := analyzeRaces(races)

	fmt.Println("Total is ", total)

	return nil
}

func main() {
	err := doLines(data0, false)
	if err != nil {
		fmt.Println("data0", data0, "got err", err)
	}

	err = doLines(data1, false)
	if err != nil {
		fmt.Println("data1", data1, "got err", err)
	}
	err = doLines(data0, true)
	if err != nil {
		fmt.Println("data0", data0, "got err", err)
	}

	err = doLines(data1, true)
	if err != nil {
		fmt.Println("data1", data1, "got err", err)
	}
}
