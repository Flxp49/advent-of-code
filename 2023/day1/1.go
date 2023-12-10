package day1

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var engnums = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var nums = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func Part1() int {
	file, _ := os.Open("inputs/input1.txt")
	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	file.Close()

	var foundL, foundR bool
	var lDigit, rDigit int
	val := 0
	for _, line := range input {
		foundL, foundR = false, false
		lDigit = 0
		rDigit = 0
		j := len(line) - 1
		for i := 0; i < len(line) && !(foundL && foundR); i++ {
			if !foundL {
				num, err := strconv.Atoi(string(line[i]))
				if err == nil {
					lDigit = num
					foundL = true
				}
			}
			if !foundR {
				num, err := strconv.Atoi(string(line[j]))
				if err == nil {
					rDigit = num
					foundR = true
				}
			}
			j--
		}
		val += lDigit*10 + rDigit
	}
	return val
}

func Part2() int {
	file, _ := os.Open("inputs/input1.txt")
	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	file.Close()

	var foundL, foundR bool
	var lDigit, rDigit int
	var lDigitIndex, rDigitIndex int
	val := 0
	for _, line := range input {
		foundL, foundR = false, false
		lDigitIndex = 0
		rDigitIndex = 0
		lDigit = 0
		rDigit = 0
		j := len(line) - 1
		for i := 0; i < len(line) && !(foundL && foundR); i++ {
			if !foundL {
				num, err := strconv.Atoi(string(line[i]))
				if err == nil {
					lDigitIndex = i
					lDigit = num
					foundL = true
				}
			}
			if !foundR {
				num, err := strconv.Atoi(string(line[j]))
				if err == nil {
					rDigitIndex = j
					rDigit = num
					foundR = true
				}
			}
			j--
		}
		q := find(line)
		if q != nil {
			if q[0] < lDigitIndex {
				lDigit = q[1]
			}
			if q[2] > rDigitIndex {
				rDigit = q[3]
			}
		}
		val += (lDigit*10 + rDigit)
	}
	return val

}

func find(line string) []int {
	var vals []int = nil
	lEngNumIndex, rEngNumIndex := -2, -2
	lEngNumDigit, rEngNumDigit := 0, 0
	for _, v := range engnums {
		i := strings.Index(line, v)
		j := strings.LastIndex(line, v)
		if lEngNumIndex == -2 && i != -1 {
			lEngNumIndex = i
			lEngNumDigit = nums[v]
		} else {
			if i != -1 && lEngNumIndex > i {
				lEngNumIndex = i
				lEngNumDigit = nums[v]
			}
		}
		if rEngNumIndex == -2 && j != -1 {
			rEngNumIndex = j
			rEngNumDigit = nums[v]
		} else {
			if j != -1 && rEngNumIndex < j {
				rEngNumIndex = j
				rEngNumDigit = nums[v]
			}
		}
	}
	if lEngNumIndex == -2 {
		return nil
	}
	return append(vals, lEngNumIndex, lEngNumDigit, rEngNumIndex, rEngNumDigit)
}
