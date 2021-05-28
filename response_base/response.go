package response_base

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ErrorHandler(err error) {
	if err != nil{
		fmt.Println(err)
		return
	}
}


func ContributionsDiff(filename string) int{
	file, err := os.Open(filename)
	ErrorHandler(err)

	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	ErrorHandler(err)

	// データが存在しない場合は0を返す
	if len(record) <= 1{
		return 0
	}

	yData := record[len(record) - 2]
	tData := record[len(record) - 1]

	yInt, _ := strconv.Atoi(yData[1])
	tInt, _ := strconv.Atoi(tData[1])

	diff := tInt - yInt
	fmt.Println(yData, tData)
	fmt.Println(diff)

	return diff
}