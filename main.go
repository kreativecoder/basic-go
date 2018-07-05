//lessons are from https://gobyexample.com
package main

import "fmt"
import "time"

func main() {
	fmt.Println("Hello go!")

	//values
	fmt.Println(20*10, 20-10, 15+3, 100/4, true && !true)

	//variables
	var a = 12
	var b int
	c := "confirm"
	fmt.Println(a, b, c)

	//constants
	const s = "test"
	fmt.Println(s)

	//Loops
	for i := 2; i <= 10; i++ {
		fmt.Println(i)
	}

	//condition
	if true {
		fmt.Println("Yes")
	}

	//switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//arrays
	arr := [4]string{"i", "am", "an", "array"}
	fmt.Println(arr)

	var arra [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			arra[i][j] = i + j
		}
	}
	fmt.Println(arra)

}
