package main

import (
	"fmt"

	"github.com/as-ifn-at/glofox/internal/config"
)

func main()  {
	config := config.Load()

	fmt.Printf("config: %v\n", config)
}