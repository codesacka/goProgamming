package main

import "fmt"

func main() {

	var count = int32(42)
	ptr := &count

	fmt.Println(*ptr)

	*ptr = 100

	fmt.Println(count)

}
