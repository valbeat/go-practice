package main

import "fmt"

func main() {
	month := map[int]string{
		1: "Jan",
		2: "Feb",
	}
	fmt.Println(month)

	_, ok := month[1]
	fmt.Println(ok)
}
