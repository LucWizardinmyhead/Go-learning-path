package main

import ( // factored import statement, good practice
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("I love the number", rand.Intn(100002)) //rand picks a random number from the given variable
}
