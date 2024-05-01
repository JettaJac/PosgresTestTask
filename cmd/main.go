package main

import (
	"fmt"
	"main/internal/config"

)
func main() {
	fmt.Println("Start server")
	config := config.NewConfig()
	fmt.Println(config)
}