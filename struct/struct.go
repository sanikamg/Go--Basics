package main

import "fmt"

//defining struct
type car struct {
	Name, model, color string
	car_no             int
}

func main() {
	//assigning values
	car1 := car{"Ferrari", "svh", "red", 323564}
	//accessing values
	fmt.Println("car name : ", car1.Name)
}
