package main

import "fmt"

func main() {
	var size_g, size int
	var buf1, buf2, buf3 int
	var ans_mass []int

	fmt.Scan(&size_g)
	for i := 0; i < size_g; i++ {
		fmt.Scan(&size)
		var mass [][3]int
		for k := 0; k < size; k++ {
			fmt.Scan(&buf1, &buf2, &buf3)
			var mas [3]int
			mas[0] = buf1
			mas[1] = buf2
			mas[2] = buf3
			mass = append(mass, mas)
		}
		ans_mass = append(ans_mass, mass[0][1], mass[0][0], mass[0][2])
		mass = append(mass[0:0], mass[1:]...)
		for len(mass) != 0 {
			for i, k := range mass {
				if k[0] == ans_mass[len(ans_mass)-1] && (k[1] == ans_mass[len(ans_mass)-2] || k[2] == ans_mass[len(ans_mass)-2]) {
					if k[1] == ans_mass[len(ans_mass)-2] {
						ans_mass = append(ans_mass, k[2])
						mass = append(mass[0:i], mass[i+1:]...)
					} else {
						ans_mass = append(ans_mass, k[1])
						mass = append(mass[0:i], mass[i+1:]...)
					}

				}
			}
			fmt.Println(ans_mass)
		}
		fmt.Println(ans_mass)

	}
}
