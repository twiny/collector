package main

import (
	"fmt"

	"github.com/twiny/collector/pkg/config"
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
