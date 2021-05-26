package main

import (
	"contributions_alerm_discord/out_traffic"
	"contributions_alerm_discord/response_base"
	"github.com/joho/godotenv"
	"os"
)


func main() {
	err := godotenv.Load(".env")
	response_base.ErrorHandler(err)

	// 環境変数からkey取り出す
	ApiKey := "Bearer " + os.Getenv("GITHUB_KEY")
	UserName := os.Getenv("GITHUB_NAME")

	out_traffic.ContributionGet(ApiKey, UserName)

	out_traffic.DiscordSend(UserName, "golang")
}