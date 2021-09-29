package main

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"io"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

type Nodes struct {
	RemoteAddress    string   `json:"remote_address"`
	Hostname         string   `json:"hostname"`
	BroadcastAddress string   `json:"broadcast_address"`
	TCPPort          int      `json:"tcp_port"`
	HTTPPort         int      `json:"http_port"`
	Version          string   `json:"version"`
	Tombstones       []bool   `json:"tombstones"`
	Topics           []string `json:"topics"`
}

type LookupNodesRsp struct {
	Producers []*Nodes `json:"producers"`
}

func main() {
	for i := 0; i < 10; i++ {
		product(i)
	}
}

func product(i int) {
	lookupIp := []string{"http://192.168.56.104:4161", "http://192.168.56.104:4163"}

	resp, err := http.Get(lookupIp[0] + "/channels?topic=test-topic2")
	if err != nil {
		log.Println("服务发现失败4161挂了")
		resp, err = http.Get(lookupIp[1] + "/channels")
		if err != nil {
			log.Println("服务发现失败4163挂了")
			log.Fatal(err)
		}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	log.Println(string(body))
	var r LookupNodesRsp
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("结果 %#v", r)
	// Instantiate a producer.
	config := nsq.NewConfig()

	randI := rand.New(rand.NewSource(time.Now().Unix())).Intn(2)
	endPoint := r.Producers[randI].BroadcastAddress + ":" + strconv.FormatInt(reflect.ValueOf(r.Producers[randI].TCPPort).Int(), 10)
	log.Println(endPoint)
	producer, err := nsq.NewProducer(endPoint, config)
	if err != nil {
		log.Fatal(err)
	}
	messageBody := []byte("hello" + strconv.Itoa(i))
	topicName := "test-topic2"
	// Synchronously publish a single message to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	err = producer.Publish(topicName, messageBody)
	if err != nil {
		log.Println("sss222")
		log.Fatal(err)
	}
	// Gracefully stop the producer when appropriate (e.g. before shutting down the service)
	producer.Stop()

}
