package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("task1Input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	task1(string(data))
	task2(string(data))
}

func task1(input string) {
	inputArrays := processListIntoArrays(input, true)
	sum := 0
	for i := range inputArrays[0] {
		sum += absInt(inputArrays[1][i] - inputArrays[0][i])
	}
	fmt.Println("Task 1 Answer ", sum)
}

func task2(input string) {
	inputArrays := processListIntoArrays(input, false)
	sum := 0
	for i := range inputArrays[0] {
		array2Count := 0
		for _, num := range inputArrays[1] {
			if num == inputArrays[0][i] {
				array2Count += 1
			}
		}

		sum += inputArrays[0][i] * array2Count
	}
	fmt.Println("Task 2 Answer ", sum)
}

func processListIntoArrays(input string, sortArray bool) [2][]int {
	array1 := []int{}
	array2 := []int{}
	list := strings.Split(input, "\n")
	for i := range list {
		splitLine := strings.Split(list[i], "  ")
		value1, err1 := strconv.Atoi(strings.TrimSpace(splitLine[0]))
		if err1 != nil {
			panic(err1)
		}
		value2, err2 := strconv.Atoi(strings.TrimSpace(splitLine[1]))
		if err2 != nil {
			panic(err2)
		}
		array1 = append(array1, value1)
		array2 = append(array2, value2)
	}
	if sortArray {
		sort.Ints(array1)
		sort.Ints(array2)
	}

	return [2][]int{array1, array2}
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
