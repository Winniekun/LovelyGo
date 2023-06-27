package main

import "fmt"

type Age uint

func (age Age) GetName()  {
	fmt.Println("this age is : ", age)
}

func (age *Age) Modify()  {
	*age = Age(100)
}

