package main 

import "github.com/01-edu/z01"

func Raid(x int) {
for i := 1 ; i <= x ; i++ {
 		if i == 1 || i == x {
 				z01.PrintRune('o')
 		} else {	
 				z01.PrintRune('-')
 		}
	}
 	z01.PrintRune(10)
}

func Raid1(x, y int) {
 	for i := 2 ; i <= y - 1 ; i ++ {
		for j := 1 ; j <= x ; j ++ {
			if j == 1 || j == x {
				z01.PrintRune('|')	
 			} else {
 				z01.PrintRune(' ')
			}	
 		}
 		z01.PrintRune(10)
 	}
}
 func Raid1a(x, y int) {
 	if x >= 1 && y >= 1 {
 		Raid(x)
 		Raid1(x,y)
 		if y != 1 {
 			Raid(x)
 		}
 	}
}


 func main() {
	 Raid1a(5,3)
}

