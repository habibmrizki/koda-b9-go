package messaging

import "fmt"

type Message struct {
	Sender  string
	Content string
}

func SendMessage(channel chan Message, sender, content string) {
	message := Message{Sender: sender, Content: content}
	channel <- message
}

func ReceivedMessage(chn chan Message) {
	for msg := range chn {
		fmt.Printf("Pengirim: %s, Pesan: %s\n", msg.Sender, msg.Content)
	}
}
