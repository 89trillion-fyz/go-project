package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type LogData struct {
	TraceId  string `json:"traceId"`  //日志唯一id
	EndPoint string `json:"endPoint"` //服务器ip:port
	LogInfo  string `json:"logInfo"`  //日志信息堆栈
	Ext      string `json:"ext"`
	Time     int    `json:"time"` //
}

type myMessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil
	}

	err := processMessage(m.Body)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return err
}

func processMessage(body []byte) error {
	fmt.Println("消息内容：", string(body))
	return nil
}

func main() {
	//for i := 0; i < 5; i++ {
	//	product()
	//}
	// Instantiate a consumer that will subscribe to the provided channel.
	config := nsq.NewConfig()
	//channel-test 是在客户端指定的，消费者连接topic时没有channel会自动创建
	// //设置重连时间
	config.LookupdPollInterval = time.Second
	consumer, err := nsq.NewConsumer("test-topic2", "channel-test-2", config)
	if err != nil {
		log.Fatal(err)
	}
	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.AddHandler(&myMessageHandler{})
	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err = consumer.ConnectToNSQLookupds([]string{"192.168.56.104:4161", "192.168.56.104:4163"})
	if err != nil {
		log.Fatal(err)
	}
	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	// Gracefully stop the consumer.
	consumer.Stop()
}
