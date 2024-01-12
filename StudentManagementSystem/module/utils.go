package module

import "log"

func checkError(err error) bool {
	if err != nil {
		log.Fatal(err.Error())
		return true //如果有错误，函数返回 true，表示发生了错误
	}
	return false
}
