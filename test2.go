package main
import (
	"fmt"
	"errors"
	"math"
)

//introducing structures
type person struct{
	name string
	age int
}

func main() {
	fmt.Println("Hello, World")
	x:= 8
	var y int = 6
	var sum int = x + y
	if y < x {
		fmt.Println("x is greater than y")
		}else if y > x {
			fmt.Println("y is greater than x")
				}else {
				fmt.Println("y is equal to x")
					}
	fmt.Println(sum)
//different methods of initializing arrays
	var a [5]int
	a[2]=7
	fmt.Println(a)
	b := [5]int{4,5,6,7,8}
	fmt.Println(b)
	c := []int{2,3,4}
//this is a slice. As in, variable length array
	fmt.Println(c)
	c = append(c,13)
	fmt.Println(c)
//Maps are dictionaries in python. key value pairs
	vertices := make(map[string]int)
	vertices["triangle"]=2
	vertices["square"]=3
	vertices["circle"]=4
	fmt.Println(vertices)
	fmt.Println(vertices["circle"])
	delete(vertices,"circle")
	fmt.Println(vertices)

//for loop is the only loop in Go
	for i := 0; i < 5; i++ {
	fmt.Println(i)
	}

//using for loop as a while loop
	i := 0
	for i<5{
	fmt.Println(i)
	i++
	}

	arr := []string{"a","b","c"}

	for index, value := range arr {
		fmt.Println("index:",index,"value:",value)
	}

//calling a function example
	result := sum1(2, 3)
	fmt.Println(result)

//another function example
	result1, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("The sqrt is:", result1)
	}

//calling structures
	p := person{name: "Jake", age: 23}
	fmt.Println(p)
	fmt.Println(p.age)
}

func sum1(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error){
	if x<0{
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}
