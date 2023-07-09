package main

import (
	"fmt"

	"github.com/juliofilizzola/bot_discord/initializers"
)

func init() {
	initializers.Env()
}

func main() {
	fmt.Println("init server")
}
