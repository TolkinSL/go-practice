package main

import "fmt"

func main() {
	var p *int
	fmt.Println(p == nil)

	x := 37
	p = &x

	fmt.Println(p)
	fmt.Println(*p)

	pp := &p
	fmt.Println(**pp)

	type User struct {
		Name string
	}

	myUser := &User{
		Name: "Vasia",
	}

	fmt.Printf("%#v\n", myUser)
	myUser.Name = "Olia"
	fmt.Printf("%#v\n", myUser)

	myFunc := func(u *User, n string) {
		u.Name = n
		fmt.Printf("Func %#v\n", u)
	}
	myFunc(myUser, "Petia")
	fmt.Printf("Main %#v\n", myUser)


}