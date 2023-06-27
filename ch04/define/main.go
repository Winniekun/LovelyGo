package main

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

func main() {
	p := Person{
		age:   30,
		name:  "飞雪无情",
		habit: "code, ACMer",
		address: Address{
			province: "北京",
			city: "北京",
		},
	}
	println(p.habit)
	println(p.address.province)
}
