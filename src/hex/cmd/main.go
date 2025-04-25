package main

import (
	"fmt"
	"hex/internal/adapters/core/arithmetic"
)

func main() {
	arithAdapter := arithmetic.NewAdapter()
	result, err := arithAdapter.Addition(1, 3)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(result)

}
