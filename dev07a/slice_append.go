package main

import "fmt"

func main() {
	var user []string

	user = append(user, "Vasia")
	fmt.Printf("cap: %d, addr: %p\n", cap(user), &user[0])

	user = append(user, "Olia")
	fmt.Printf("cap: %d, addr: %p\n", cap(user), &user[0])

	user = append(user, "Tania")
	fmt.Printf("cap: %d, addr: %p\n", cap(user), &user[0])

	user = append(user, "Petia")
	fmt.Printf("cap: %d, addr: %p\n", cap(user), &user[0])

	user2 := make([]string, 0, 4)
	fmt.Printf("%#v\n", user2)
	// user2[0] = "Vasia" // panic: runtime error: index out of range [0] with length 0
	user2 = append(user2, "Vasia")
	fmt.Printf("%#v\n", user2)
}