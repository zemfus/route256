package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	var str string
	for i := 0; i < size; i++ {
		fmt.Scan(&str)
		for str != "" {
			if "00" == str[:2] {
				fmt.Print("a")
				str = str[2:]
				continue
			}
			if "11" == str[:2] {
				fmt.Print("d")
				str = str[2:]
				continue
			}
			if "100" == str[:3] {
				fmt.Print("b")
				str = str[3:]
				continue
			}
			if "101" == str[:3] {
				fmt.Print("c")
				str = str[3:]
			}

		}
		fmt.Println()
	}
}
