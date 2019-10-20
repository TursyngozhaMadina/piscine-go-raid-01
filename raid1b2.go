package student

import "github.com/01-edu/z01"

//Raid1b Bruce Wayne is Batman
func Raid1b(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if j == 1 && i == 1 {
				z01.PrintRune('/')
			} else if (j == x && i == 1) || (j == 1 && i == y) {
				z01.PrintRune(92)
			} else if j == x && i == y {
				z01.PrintRune('/')
			} else if j == 1 || j == x {
				z01.PrintRune('*')
			} else if i == 1 || i == y {
				z01.PrintRune('*')
			} else {
				z01.PrintRune(' ')
			}
			if j == x {
				z01.PrintRune('\n')

			}

		}
	}
}
