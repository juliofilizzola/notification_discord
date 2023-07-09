package main

import (
	"fmt"

	"github.com/juliofilizzola/bot_discord/initializers"
	"github.com/juliofilizzola/bot_discord/internal/send"
)

func init() {
	initializers.Env()
}

func main() {
	send.MessageDiscord()
	fmt.Println("init server")
}
