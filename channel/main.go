package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan string) // unbuffered channel

	go uploadFile(ch)

	fileUrl := <-ch // blocking...

	fmt.Println("File url :", fileUrl)

}

func uploadFile(c chan string) {
	fmt.Println("uploading file...")
	time.Sleep(time.Second * 3)
	fmt.Println("file upload done!")

	fileUrl := "https://www.google.com"
	c <- fileUrl
}
