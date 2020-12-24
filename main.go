package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		First: "Jenny",
	}

	p2 := person{
		First: "James",
	}

	xp := []person{p1, p2}

	bytes, err := json.Marshal(xp)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
