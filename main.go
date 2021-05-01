package main

import (
	log "log"

	bot "github.com/rvramesh/tg-microwin-bot/src/bot"
	helper "github.com/rvramesh/tg-microwin-bot/src/bot/helpers"
	db "github.com/rvramesh/tg-microwin-bot/src/db"
)

// main function start the application.
// First load the file .env with the environment variables
// and ask for the token to call the main bot initializer
func main() {

	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function
	db.Init()

	if err := helper.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	token := helper.GetBotToken()

	log.Fatal(bot.StartBot(token))
}
