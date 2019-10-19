package student

import "github.com/01-edu/z01"

//Raid1a Bruce Wayne is Batman
func Raid1a(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (j == 1 && i == 1) || (j == x && i == y) || (j == x && i == 1) || (j == 1 && i == y) {
				z01.PrintRune('o')
			} else if j == 1 || j == x {
				z01.PrintRune('|')
			} else if (i == 1 || i == y) || (j == 1 || j == x) {
				z01.PrintRune('-')
			} else {
				z01.PrintRune(' ')
			}
			if j == x {
				z01.PrintRune('\n')

			}

		}
	}
}