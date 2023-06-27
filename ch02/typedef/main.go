package main

import "fmt"

type Celsius float64
type Fahrenherit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 0
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenherit) String() string {
	return fmt.Sprintf("%g°C", f)
}
