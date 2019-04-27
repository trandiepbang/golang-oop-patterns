package main

import "fmt"

var data string

func GetData() string {
	if data == "" {
		data = "data"
	}
	return data
}

func main() {
	dataGetter := GetData()
	fmt.Println("receive data : ", dataGetter)
	data = ""
	dataGetter = GetData()
	fmt.Println("receive data 2 : ", dataGetter)
}
