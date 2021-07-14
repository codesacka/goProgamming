package main

import "fmt"

func main() {
	// y := string("Hello World")
	var s = make([]string, 0)
	var m = make(map[string]string)
	s = append(s, "some string")
	m["some"] = "some value"
	fmt.Println(s)
}
