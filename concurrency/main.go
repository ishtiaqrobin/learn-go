package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var fileUrl string

func main() {
	var start = time.Now()

	wg.Go(uploadFile)
	wg.Go(saveToDB)
	wg.Go(sendEmail)

	wg.Wait() // wait for all goroutines to finish. waiting... until counter is 0

	fmt.Println("file url", fileUrl)
	fmt.Println("all task completed")
	fmt.Println("time taken: ", time.Since(start))
}

func uploadFile() {
	fmt.Println("uploading file...")
	time.Sleep(time.Second * 3)
	fmt.Println("file upload done!")

	fileUrl = "https://www.google.com"
}

func saveToDB() {
	fmt.Println("saving to db...")
	time.Sleep(time.Second * 2)
	fmt.Println("saved!")
}

func sendEmail() {
	fmt.Println("sending email...")
	time.Sleep(time.Second * 1)
	fmt.Println("email sent!")
}
