package main

import (
	"e-mar404/gator/internal/config"
	"fmt"
)

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.SetUser("emar")
	c, err = config.Read()
	fmt.Println(c)
}
