package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkLeft(starMap []string, x int, y int, rows int, columns int) bool {
	if y == 0 {
		return false
	}
	if starMap[x][y-1] == '*' {
		return true
	}
	return false
}

func checkRight(starMap []string, x int, y int, rows int, columns int) bool {
	if y == columns-1 {
		return false
	}
	if starMap[x][y+1] == '*' {
		return true
	}
	return false
}

func checkUp(starMap []string, x int, y int, rows int, columns int) bool {
	if x == 0 {
		return false
	}
	if starMap[x-1][y] == '*' {
		return true
	}
	return false
}

func checkDown(starMap []string, x int, y int, rows int, columns int) bool {
	if x == rows-1 {
		return false
	}
	if starMap[x+1][y] == '*' {
		return true
	}
	return false
}

func checkSolo(starMap []string, x int, y int, rows int, columns int) bool {
	res := 0
	if checkLeft(starMap, x, y, rows, columns) {
		res += 1
	}
	if checkRight(starMap, x, y, rows, columns) {
		res += 1
	}
	if checkUp(starMap, x, y, rows, columns) {
		res += 1
	}
	if checkDown(starMap, x, y, rows, columns) {
		res += 1
	}
	if res == 1 {
		return true
	}
	return false
}

func findSolo(starMap []string, rows int, columns int) (x int, y int) {
	for row, str := range starMap {
		for column, value := range str {
			if value == '*' {
				if checkSolo(starMap, row, column, rows, columns) {
					return row, column
				}
			}
		}
	}
	panic("COULD NOT FIND START POINT")
	return 0, 0
}

func solution(starMap []string, rows int, columns int) string {
	x, y := findSolo(starMap, rows, columns)
	conFlag := true
	prevType := -1
	var res string
	for conFlag {
		if prevType != 2 && checkRight(starMap, x, y, rows, columns) {
			prevType = 0
			y += 1
			res += "R"
		} else if prevType != 0 && checkLeft(starMap, x, y, rows, columns) {
			prevType = 2
			y -= 1
			res += "L"
		} else if prevType != 3 && checkUp(starMap, x, y, rows, columns) {
			prevType = 1
			x -= 1
			res += "U"
		} else if prevType != 1 && checkDown(starMap, x, y, rows, columns) {
			prevType = 3
			x += 1
			res += "D"
		} else {
			conFlag = false
		}
	}
	return res
}
func main() {
	var numOfMaps int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &numOfMaps)
	var rows, columns int
	for ; numOfMaps > 0; numOfMaps-- {
		fmt.Fscan(in, &rows, &columns)
		starMap := make([]string, rows)
		for i := 0; i < rows; i++ {
			fmt.Fscan(in, &starMap[i])
		}
		fmt.Println(solution(starMap, rows, columns))
	}

}
