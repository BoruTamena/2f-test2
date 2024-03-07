package main

import (
	"fmt"
)

type Message struct {
	Text string
}

func sender(msgCh chan<- Message, doneCh chan<- bool, messages []Message) {

	for _, msg := range messages {
		msgCh <- msg // sending message over channel
	}
	doneCh <- true // completing
	close(msgCh)   // closing the msg channel

}

func main() {

	msgCh := make(chan Message) // Channel for sending messages
	doneCh := make(chan bool)   // Channel for signaling completion

	msg1 := []Message{
		{
			Text: "Message one ",
		},
	}

	go sender(msgCh, doneCh, msg1)

	go func() {

		defer close(doneCh) // closing the done channel
		<-msgCh             // receiving the message

	}()

	done := <-doneCh

	if done {
		fmt.Println("successfully completed")
	}

}
