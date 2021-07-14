package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Bar string
	Baz string
}

func main() {

	f := Foo{"Juniour", "Hello Shabado"}

	b, _ := json.Marshal(f)

	fmt.Println(string(b))
	json.Unmarshal(b, &f)

	fmt.Println(string(b))
}
