package main

import "fmt"

func main(){
	fmt.Print("Enter a Number: ")
	var input int
	fmt.Scanln(&input)

	if input%2 == 0	{
		fmt.Println("Even")
    } else {
		fmt.Println("Odd")
    }
	/*
	x := 10000
	if x = 10; 100%2 == 0 {
		fmt.Println("Even")
		x = 200
	} else {
		fmt.Println("Odd")
    }
	fmt.Println(x)
	*/
	 data := 10 
	 switch data {
	 case 10:
		data = 100
		fmt.Println(data)
	case 100:
		data = 103
		fmt.Println(data)
		fallthrough
	case 11:
		data = 104
		fmt.Println(data)
		fallthrough
	case 15:
		data = 1002
		fmt.Println(data)
		fallthrough
	default:
		data = 10909
		fmt.Println(data)
	 }

}