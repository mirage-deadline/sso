package main

import (
	"fmt"
	"github.com/mirage-deadline/sso/internal/config"
)

func main() {
	cfg := config.NewConfig()
	fmt.Println(cfg)
}
