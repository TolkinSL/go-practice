package main

import "fmt"

func main() {
	var mySlice1 []int
	fmt.Printf("mySlice1 %#v\n", mySlice1)
	fmt.Println(len(mySlice1))

	mySlice2 := []int{}
	fmt.Printf("mySlice2 %#v\n", mySlice2)
	fmt.Println(len(mySlice2))

	mySlice1 = append(mySlice1, 4)
	fmt.Printf("mySlice1 %#v\n", mySlice1)
	fmt.Println(len(mySlice1))

	mySlice3 := make([]int, 0, 4)
	fmt.Printf("mySlice3 %#v\n", mySlice3)

	mySlice1 = []int{1, 2, 3}
	mySlice2 = make([]int, len(mySlice1))
	copy(mySlice2, mySlice1)
	fmt.Printf("mySlice2 %#v\n", mySlice2)

	mySlice4 := []rune("Привет")
	fmt.Printf("mySlice4 %#v\n", mySlice4)

	fmt.Printf("hackSlice2----\n")
	hackSlice1 := []string{"a", "b", "c", "d"}
	hackSlice2 := hackSlice1[0:2]
	fmt.Printf("hackSlice2 %#v\n", hackSlice2)
	fmt.Printf("hackSlice2 %#v\n", hackSlice2[:3])
	hackSlice2 = hackSlice1[0:2:2]
	fmt.Printf("hackSlice2 %#v\n", hackSlice2)
	// fmt.Printf("hackSlice2 %#v\n", hackSlice2[:3]) //panic: runtime error: slice bounds out of range [:3]
	hackSlice2[1] = "z"
	fmt.Printf("hackSlice1 %#v\n", hackSlice1)
	fmt.Printf("hackSlice2 %#v\n", hackSlice2)

	// Удаление
	idx := 2
	hackSlice1 = []string{"a", "b", "c", "d"}
	fmt.Printf("hackSlice1 %#v\n", hackSlice1)
	hackSlice1 = append(hackSlice1[:idx],hackSlice1[idx+1:]... )
	fmt.Printf("hackSlice1 %#v\n", hackSlice1)

}
