package main

import (
	"fmt"

	"github.com/tieba/function"
)

func main() {
	list := function.GetLikeList()
	fmt.Println(list)
}
