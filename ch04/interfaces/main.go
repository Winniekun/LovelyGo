package main

import "fmt"

type Person struct {
	age     int
	name    string
	habit   string
	address Address
}

type Address struct {
	province string
	city     string
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func (p Person) String() string {
	return fmt.Sprintf("the name is %s, age is %d", p.name, p.age)
}

func main() {
	p := Person{
		age:   10,
		name:  "wkk",
		habit: "code",
		address: Address{
			province: "北京",
			city:     "北京",
		},
	}
	printString(p)
	printString(&p)
}
