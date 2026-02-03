package main

import "fmt"

func main() {
	num := 2

	switch {
	case num == 1:
		fmt.Println("num = 1")
	case num == 2:
		fmt.Println("num = 2")
	case num == 3:
		fmt.Println("num = 3")
	default:
		fmt.Println("num = ?")
	}

	num = 4

	switch num {
	case 1:
		fmt.Println("num = 1")
	case 2:
		fmt.Println("num = 2")
	case 3:
		fmt.Println("num = 3")
	default:
		fmt.Println("num = ?")
	}

	num = 2

	switch num {
	case 1:
		fmt.Println("num = 1")
	case 2:
		fmt.Println("num = 2")
		fallthrough
	case 3:
		fmt.Println("num = 3")
	default:
		fmt.Println("num = ?")
	}

	num = 1
	switch num {
	case 1, 2, 3:
		fmt.Println("num = 1-2-3")
	case 4:
		fmt.Println("num = 4")
	default:
		fmt.Println("num = ?")
	}

	for i := range []int{1, 2, 3} {
		fmt.Println("Hello", i)
	}

}
