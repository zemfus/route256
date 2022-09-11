package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func checkEnoughUniqInSorted(array []int) bool {
	counter := 1
	prev := array[0]
	for _, value := range array[1:] {
		if value != prev {
			counter++
			prev = value
			if counter >= 4 {
				return true
			}
		}
	}
	return false
}

//accepts only sorted array
func findSteps(array []int) ([]int, bool) {
	step := 1
	curLen := 0
	result := make([]int, 4)
	prev := array[len(array)-1]
	for i := len(array) - 2; i >= 0 && step < 5; i-- {
		curLen++
		if array[i] != prev {
			prev = array[i]
			if curLen >= step {
				result[4-step] = prev
				step++
				curLen = 0
			}
		}
	}
	if step < 5 {
		return result, true
	}
	return result, false
}

func solution(array []int, len int) {
	sortedArray := make([]int, len, len)
	copy(sortedArray, array)
	sort.Ints(sortedArray)
	steps, errorFlag := findSteps(sortedArray)
	markedStars := make([]int, 5)
	if !checkEnoughUniqInSorted(sortedArray) {
		errorFlag = true
	} else {
		for index, value := range array {
			star := 0
			for ; star < 4 && value > steps[star]; star++ {
			}
			array[index] = star + 1
			markedStars[star]++
		}
	}
	for index, value := range markedStars[1:] {
		if markedStars[index] < value {
			errorFlag = true
			break
		}
	}
	if errorFlag {
		for index, _ := range array {
			array[index] = -1
		}
	}
	fmt.Println(strings.Trim(fmt.Sprint(array), "[]"))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var numOfTests, len int
	fmt.FScan(in, &numOfTests)
	for ; numOfTests > 0; numOfTests-- {
		fmt.FScan(in, &len)
		array := make([]int, len)
		for i := 0; i < len; i++ {
			fmt.FScan(in, &array[i])
		}
		solution(array, len)
	}
}
