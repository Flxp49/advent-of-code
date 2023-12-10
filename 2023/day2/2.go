package day2

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// var input = [...]string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
// 	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
// 	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
// 	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
// 	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}

func Day2() {
	file, _ := os.Open("input2.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	// vals := map[string]int{"red": 0, "blue": 0, "green": 0}
	// maxVals := map[string]int{"red": 12, "blue": 14, "green": 13}
	regNums, _ := regexp.Compile(`[0-9]+`)
	regColor, _ := regexp.Compile(`[a-z]+`)
	var num, color []string
	var d int
	// var add bool
	// game := 0

	// Part 1
	// for i, x := range input {
	// 	r := strings.Split(x, ":")
	// 	rr := strings.Split(r[1], ";")
	// 	add = true
	// b:
	// 	for _, v := range rr {
	// 		vals["red"] = 0
	// 		vals["blue"] = 0
	// 		vals["green"] = 0
	// 		num = regNums.FindAllString(v, -1)
	// 		color = regColor.FindAllString(v, -1)
	// 		for j := range num {
	// 			d, _ = strconv.Atoi(num[j])
	// 			vals[color[j]] += d
	// 			if vals[color[j]] > maxVals[color[j]] {
	// 				add = false
	// 				break b
	// 			}
	// 		}
	// 	}
	// 	if add {
	// 		game += (i + 1)
	// 	}
	// }

	// Part 2
	minVals := map[string]int{"red": 0, "blue": 0, "green": 0}
	power := 0

	for _, x := range input {
		r := strings.Split(x, ":")
		rr := strings.Split(r[1], ";")
		minVals["red"] = 0
		minVals["green"] = 0
		minVals["blue"] = 0
		for _, v := range rr {
			num = regNums.FindAllString(v, -1)
			color = regColor.FindAllString(v, -1)
			for j := range num {
				d, _ = strconv.Atoi(num[j])
				if minVals[color[j]] < d {
					minVals[color[j]] = d
				}
			}
		}
		power += (minVals["red"] * minVals["green"] * minVals["blue"])
	}
	log.Println(power)
}
