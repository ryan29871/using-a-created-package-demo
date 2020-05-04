package main

import (
	"fmt"

	"github.com/ryan29871/hello-world/calc"
	calcNew "github.com/ryan29871/hello-world/v2/calc"
)

func main() {
	result := calc.Add(1, 2, 3)
	fmt.Println("calc.Add(1, 2, 3) => ", result)

	// v2 of hello-world
	newResult, err := calcNew.Add(1, 2, 3, 4)

	if err != nil {
		fmt.Println("Error: => ", err)
	} else {
		fmt.Println("calcNew.Add(1, 2, 3, 4) => ", newResult)
	}

}
