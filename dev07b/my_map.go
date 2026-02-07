package main

import "fmt"

func main() {
	var userId1 map[string]int
	fmt.Printf("%#v\n", userId1)
	// Паника при записи значений

	userId2 := map[string]int{}
	userId2["Vasia"] = 1236
	fmt.Printf("%#v\n", userId2)

	val := userId2["Vasia"]
	val, ok := userId2["Olya"]
	fmt.Printf("val: %d, ok: %t\n", val, ok)

	userId2["Olya"] = 14
	userId2["Tanya"] = 746

	fmt.Printf("%#v\n", userId2)
	delete(userId2, "Vasia")
	fmt.Printf("%#v\n", userId2)

	for key, val := range userId2 {
		fmt.Println(key, val)
	}

	// Создание сета с литералом пустой структуры, ничего не весит
	// Дедупликация
	
	fmt.Println("\nCreate Set")

	userId3 := map[string]struct{}{}
	userId3["Vasia"] = struct{}{}
	userId3["Olya"] = struct{}{}
	userId3["Vasia"] = struct{}{}
	fmt.Printf("%#v\n", userId3)
}
