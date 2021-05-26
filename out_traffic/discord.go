package out_traffic

import (
	"contributions_alerm_discord/response_base"
	"net/http"
	"os"
	"strings"
)

func DiscordSend(UserName string, Content string) {
	url := os.Getenv("DISCORD_URL")
	method := "POST"

	payloadBase := `{`+""+`"username": "`+ UserName +`",`+""+`"content": "` + Content + `"`+""+`}`
	payload := strings.NewReader(payloadBase)

	// httpリクエスト用意
	client := &http.Client {}
	req, err := http.NewRequest(method, url, payload)
	response_base.ErrorHandler(err)

	// headerに付与
	req.Header.Add("Content-Type", "application/json")

	// API通信
	res, err := client.Do(req)
	response_base.ErrorHandler(err)
	defer res.Body.Close()
}