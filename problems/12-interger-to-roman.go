package main

import (
	"math"
	"strconv"
	"strings"
)

type intRoman struct {
	number int
	char   string
}

func intToRoman(input int) string {
	aTable := []intRoman{
		{
			number: 9,
			char:   "IX",
		},
		{
			number: 5,
			char:   "V",
		},
		{
			number: 4,
			char:   "IV",
		},
		{
			number: 1,
			char:   "I",
		},
	}

	bTable := []intRoman{
		{
			number: 90,
			char:   "XC",
		},
		{
			number: 50,
			char:   "L",
		},
		{
			number: 40,
			char:   "XL",
		},
		{
			number: 10,
			char:   "X",
		},
	}

	cTable := []intRoman{
		{
			number: 900,
			char:   "CM",
		},
		{
			number: 500,
			char:   "D",
		},
		{
			number: 400,
			char:   "CD",
		},
		{
			number: 100,
			char:   "C",
		},
	}

	dTable := []intRoman{
		{
			number: 1000,
			char:   "M",
		},
	}

	//input.
	result := ""

	strInput := strconv.Itoa(input)
	for idx, value := range strings.Split(strInput, "") {
		pos := (len(strInput) - 1) - idx
		data, _ := strconv.Atoi(value)
		char := data * int(math.Pow(10, float64(pos)))

		var table *[]intRoman
		if pos == 3 {
			table = &dTable
		} else if pos == 2 {
			table = &cTable
		} else if pos == 1 {
			table = &bTable
		} else {
			table = &aTable
		}

		for {
			for _, v := range *table {
				if v.number <= char {
					char = char - v.number
					result = result + v.char
					break
				}
			}
			if char == 0 {
				break
			}
		}
	}

	return result
}

func main() {
	print(intToRoman(60))
}
