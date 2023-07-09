package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port         = ""
	TokenDiscord = ""
	WebhookId    = ""
	WebhookToken = ""
	AvatarURL    = ""
	Username     = ""
)

func Env() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Port = fmt.Sprint(os.Getenv("PORT"))
	TokenDiscord = fmt.Sprint(os.Getenv("TOKEN_DISCORD"))
	WebhookId = fmt.Sprint(os.Getenv("WEBHOOK_ID"))
	WebhookToken = fmt.Sprint(os.Getenv("WEBHOOK_TOKEN"))
	Username = fmt.Sprint(os.Getenv("USER_NAME"))
	AvatarURL = fmt.Sprint(os.Getenv("AVATAR_URL"))
}
