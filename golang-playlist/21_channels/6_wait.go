package main

import (
    "fmt"
    "sync"
    "time"
)

func emailSender(emailChan <-chan string) {
    for email := range emailChan {
        fmt.Println("sending email to", email)
        time.Sleep(time.Second)
    }
}

func main() {
    emailChan := make(chan string, 100)
    var wg sync.WaitGroup

    wg.Go(func() {
        emailSender(emailChan)
    })
	//wg.Go(emailSender(emailChan)) // this calls emailSender immediately
	// we pass the function reference to wg.Go not the function call

    for i := 0; i < 5; i++ {
        emailChan <- fmt.Sprintf("%d@gmail.com", i)
    }

    fmt.Println("done sending.")
    close(emailChan)

    wg.Wait()
}