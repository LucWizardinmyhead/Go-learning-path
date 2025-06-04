package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("! WHATT")

	os.Exit(4)
}
