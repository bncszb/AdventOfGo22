package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("input")
	check(err)

	bags := string(dat)

	split_bags := strings.Split(bags, "\n\n")
	num_of_bags := int(len(split_bags))
	cal_per_bag := make([]int, num_of_bags)

	for i := 0; i < num_of_bags; i++ {
		content_of_bag := strings.Split(split_bags[i], "\n")
		for j := 0; j < len(content_of_bag); j++ {
			if content_of_bag[j] != "" {
				calorie, err := strconv.Atoi(content_of_bag[j])
				check(err)

				cal_per_bag[i] += calorie
			}
		}
	}

	sort.Ints(cal_per_bag)

	fmt.Print(cal_per_bag[num_of_bags-1], "\n")

	last_N := 3
	sum := 0
	for i := num_of_bags - (last_N); i < num_of_bags; i++ {
		sum += cal_per_bag[i]
	}
	fmt.Print(sum, "\n")

}
