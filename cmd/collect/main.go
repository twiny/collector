package main

import (
	"collector/pkg/config"
	"fmt"
)

// here we go!
func main() {
	c, err := config.ParseConfig("config/config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", c)
}
