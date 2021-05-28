package out_traffic

import (
	"contributions_alerm_discord/json_type"
	"contributions_alerm_discord/response_base"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func ContributionGet(ApiKey string, UserName string, fileName string)  {
	// urlとかmethodの形式書く
	url := "https://api.github.com/graphql"
	method := "POST"

	// GraphQLの構文を無理やり書く
	// 後日GraphQLのライブラリ使って何とかします
	payloadBase := "{\"query\":\"query($userName:String!) {\\r\\n  user(login: $userName){\\r\\n    contributionsCollection {\\r\\n      contributionCalendar {\\r\\n        totalContributions\\r\\n      }\\r\\n    }\\r\\n  }\\r\\n}\",\"variables\":{\"userName\":\"" + UserName + "\"}}"
	payload := strings.NewReader(payloadBase)

	// httpリクエスト用意
	client := &http.Client {}
	req, err := http.NewRequest(method, url, payload)
	response_base.ErrorHandler(err)

	// headerに付与
	req.Header.Add("Authorization", ApiKey)
	req.Header.Add("Content-Type", "application/json")

	// API通信
	res, err := client.Do(req)
	response_base.ErrorHandler(err)
	defer res.Body.Close()

	// responseデータを読み取れるようにする
	body, err := ioutil.ReadAll(res.Body)
	response_base.ErrorHandler(err)

	// json形式で読み込む
	var bytes = []byte(body)
	var response json_type.ContributionsJson
	err = json.Unmarshal(bytes, &response)
	response_base.ErrorHandler(err)
	contributions := response.Data.User.Contributionscollection.Contributioncalendar.Totalcontributions
	contributionsStr := strconv.Itoa(contributions)

	// 取得時間を表示
	today := time.Now()
	today_ := today.Truncate(time.Second)
	strToday := today_.Format("2006/01/02")

	fmt.Println(strToday, contributionsStr)

	// csvにて保存
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	response_base.ErrorHandler(err)
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{strToday, contributionsStr})
	writer.Flush()
}
