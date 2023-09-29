package main

import (
	"fmt"
	"go-logger/config"
	"os"
)

func main() {
	fmt.Println("test")
	config.LoadENV()
	fmt.Println(os.Getenv("test"))
}
