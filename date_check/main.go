package main

import "fmt"

var mounths_size = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func main() {
	var size int
	fmt.Scan(&size)
	var day, month, year int
	for i := 0; i < size; i++ {
		fmt.Scan(&day, &month, &year)
		if month <= 12 && month > 0 {
			if (day > 0 && day <= mounths_size[month-1]) || (month == 2 && (year%400 == 0 || (year%4 == 0 && year%100 != 0)) && day > 0 && day <= 29) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		} else {
			fmt.Println("NO")
		}
	}
}
