package main

import "fmt"

//* any and interface{} are same but interface{} is more generic
//* type assertion
//* ok idiom

func printData(data any) {
	strData, ok := data.(string)
	if ok {
		fmt.Println(len(strData))
	}

	intData, ok := data.(int)
	if ok {
		fmt.Println(intData + 100)
	}

}

func main() {
	var data = "100"
	printData(data)
}
