package out_traffic

import (
	"bytes"
	"contributions_alerm_discord/json_type"
	"contributions_alerm_discord/response_base"
	"encoding/json"
	"net/http"
)

func DiscordSend(UserName string, Content string, WebHookURL string) {
	// 構造体を生成してその中に要素を保存する
	payloadBase := new(json_type.DiscordJson)
	payloadBase.UserName = UserName
	payloadBase.Content = Content

	// json形式に変える
	payload, err := json.Marshal(payloadBase)
	response_base.ErrorHandler(err)

	// jsonを添えてPostで送信する
	res, err := http.Post(WebHookURL, "application/json", bytes.NewBuffer(payload))
	response_base.ErrorHandler(err)
	defer res.Body.Close()
}