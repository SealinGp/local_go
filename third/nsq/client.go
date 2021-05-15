package nsq

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

func Consume(topic, channel string) {

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Println("newConsumer err:", err)
		return
	}

	consumer.AddHandler(&nsqClient{})

	err = consumer.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		fmt.Println("connect err:", err)
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Gracefully stop the consumer.
	consumer.Stop()
}

type nsqClient struct{}

func (h *nsqClient) HandleMessage(m *nsq.Message) error {

	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		// In this case, a message with an empty body is simply ignored/discarded.
		return nil
	}

	// do whatever actual message processing is desired
	err := h.processMessage(m.Body)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return err
}
func (h *nsqClient) processMessage(body []byte) error {
	fmt.Println("process message", string(body))
	return nil
}
