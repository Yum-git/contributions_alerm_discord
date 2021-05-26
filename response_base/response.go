package response_base

import "fmt"

func ErrorHandler(err error) {
	if err != nil{
		fmt.Println(err)
		return
	}
}
