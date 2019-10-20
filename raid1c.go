package main

import "github.com/01-edu/z01"

func Raid1c(x,y int) {
	for i := 1 ; i <= y ; i ++ {
		for j := 1 ; j <= x ; j ++ {
			if (i == 1 && j == 1) || (j == x && i == 1 ) {
				z01.PrintRune('A') 
			} else if (i == y && j == 1) || (j == x && i == y ) {
				z01.PrintRune('C')
			} else if i == 1 || i == y {
				z01.PrintRune('B')
			} else if j == 1  || j == x {
				z01.PrintRune('B')
			}   else { 
				z01.PrintRune(' ')		
			} 
			if j == x {
				z01.PrintRune('\n')
			}
			
		}
		
	}
}
func main() {
	Raid1c(1,5)
}
