package main

import "qqbot/bot"

func main() {
	bot.SetupDatabase()
	bot.StartListen()
}
