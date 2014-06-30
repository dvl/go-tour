package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
    fmt.Printf("%v is my favorite number", rand.Intn(10))
}
