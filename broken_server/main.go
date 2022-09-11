// a 00
// b 100
// c 101
// d 11

package main

import "fmt"

func main() {
	var size_g, size, ans, ans_f int
	var buf, buf1, buf2 int

	fmt.Scan(&size_g)
	for i1 := 0; i1 < size_g; i1++ {
		fmt.Scan(&size)
		var mass []int
		for k := 0; k < size; k++ {
			fmt.Scan(&buf)
			mass = append(mass, buf)
		}
		for i, k := range mass {
			buf1 = k
			ii := i
			for ii < len(mass) && buf1 == mass[ii] {
				ii++
				ans++
			}
			if ii >= len(mass) {
				if buf2 == 0 {
					ans_f = ans
				}

				break

			}
			buf2 = mass[ii]
			for ii < len(mass) && (buf1 == mass[ii] || buf2 == mass[ii]) {
				ii++
				ans++
			}
			if ans_f < ans {
				ans_f = ans
			}
			ans = 0
		}
		fmt.Println(ans_f)
		ans = 0
		ans_f = 0
		buf1 = 0
		buf2 = 0
	}
}
