package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var engnums = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var nums = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func main() {
	file, _ := os.Open("input1.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var l, r int
	val := 0
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	// part 1
	// 	for i := 0; i < len(input); i++ {
	// 		x, err := strconv.Atoi(string(input[i]))
	// 		if err == nil {
	// 			l = x
	// 			break
	// 		}
	// 	}
	// 	for i := len(input) - 1; i >= 0; i-- {
	// 		x, err := strconv.Atoi(string(input[i]))
	// 		if err == nil {
	// 			r = x
	// 			break
	// 		}
	// 	}
	// 	val += l*10 + r
	// }
	// fmt.Println(val)

	// part 2
	for _, v := range input {
		l = 99
		r = -1
		min, max := 0, 0
		for i := 0; i < len(v); i++ {
			x, err := strconv.Atoi(string(v[i]))
			if err == nil {
				l = i
				min = x
				break
			}
		}
		for i := len(v) - 1; i >= 0; i-- {
			x, err := strconv.Atoi(string(v[i]))
			if err == nil {
				r = i
				max = x
				break
			}
		}
		q := find(v)
		if q != nil {
			if q[0] < l {
				min = q[1]
			}
			if q[2] > r {
				max = q[3]
			}
		}
		val += (min*10 + max)
	}
	fmt.Println(val)

}

func find(n string) []int {
	var found []int
	var vals []int = nil
	var min, max int
	for _, v := range engnums {
		i := strings.Index(n, v)
		j := strings.LastIndex(n, v)
		if i != -1 {
			found = append(found, i, nums[v])
			if j != -1 && j != i {
				found = append(found, j, nums[v])
			}
		}
	}
	if len(found) == 0 {
		return nil
	}
	var minval, maxval int
	for q := 0; q < len(found); q += 2 {
		if q == 0 {
			min = found[q]
			max = found[q]
			minval = found[q+1]
			maxval = found[q+1]
			continue
		}
		if min > found[q] {
			min = found[q]
			minval = found[q+1]
		}
		if max < found[q] {
			max = found[q]
			maxval = found[q+1]
		}
	}
	return append(vals, min, minval, max, maxval)
}
