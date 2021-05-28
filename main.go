package main

import (
	"contributions_alerm_discord/out_traffic"
	"contributions_alerm_discord/response_base"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)


func main(){
	// .envファイルから環境変数を取得する
	err := godotenv.Load(".env")
	response_base.ErrorHandler(err)

	// 環境変数からkey取り出す
	ApiKey := "Bearer " + os.Getenv("GITHUB_KEY")
	DiscordURL := os.Getenv("DISCORD_URL")
	UserName := os.Getenv("GITHUB_NAME")
	FileName := os.Getenv("FILE_NAME")
	if FileName == ""{
		FileName = "contributions.csv"
	}

	// Contributionを取得してcsvに書き込む
	// 将来的にはSQLとかに…（SQLiteがいいかな）
	out_traffic.ContributionGet(ApiKey, UserName, FileName)

	// Contributionの差異を取得する
	diff := response_base.ContributionsDiff(FileName)

	// 前回取得データとの差異が検出出来たらDiscordに送信する
	if diff > 0{
		message := strconv.Itoa(diff) + "回草を生やしました\n" + "GitHubURL:https://github.com/" + UserName
		fmt.Println(message)
		out_traffic.DiscordSend(UserName, message, DiscordURL)
	}
}