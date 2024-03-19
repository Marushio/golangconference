package main

import "fmt"
import "time"

//#5-Data orientation
type Person struct {
	name string
	age int
}
//#5.1-method that belong to Person
func (p Person) Walk() {
	fmt.Println(p.name + " is walking")
}

//#6-Parallelism
func counter(count int) {
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

//#6-Parallelism
// GO ROUTINE -> 1
func main() {
	//#1-Hello world
	println("Hello World!!!")
	//#2-vars declarations
	var num int
	num = 10
	fmt.Println(num)

	//#3-array size fixed
	//var intArr [5]int

	//#4-slice example
	//var slc []int
	//#type search how slies work with memory and garbage collector 
	
	//#5-Data orientation
	var person Person
	person.name = "Lucas"
	person.age = 25
	person.Walk()

	//#6-Parallelism
	go counter(10) // routine(thread) 2 
	go counter(10) // routine(thread) 3 ->start with 2kb | other applications normaly start with 1mb
	counter(10)
	

}

