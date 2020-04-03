package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	cmq "github.com/zzzhr1990/qcmq"
	"github.com/zzzhr1990/qcmq/easy"
)

const (
	secretID           = ""
	secretKey          = ""
	queueName          = ""
	pollingWaitSeconds = 1
)

func main() {
	// create a CMQ client
	client := cmq.NewClient(secretID, secretKey, cmq.RegionGuangzhou, cmq.InternetAccess)

	// send a message
	if false {
		sendMsgReq := &cmq.SendMessageRequest{
			QueueName:   cmq.String(queueName),
			MessageBody: cmq.String("hello world"),
		}
		sendMsgResp, err := client.SendMessage(sendMsgReq)
		if err != nil {
			panic(err)
		}
		data, err := json.Marshal(*sendMsgResp)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}

	pss := easy.NewProducer(secretID, secretKey, cmq.RegionGuangzhou, cmq.InternetAccess, queueName)
	css := easy.NewConsumer(secretID, secretKey, cmq.RegionGuangzhou, cmq.InternetAccess, queueName)
	ch := make(chan *easy.Message)
	css.StartListen(ch)
	var count int64 = 0

	go func() {
		for bb := range ch {
			// fmt.Println("Recv: -> " + bb.Body)
			css.Ack(bb)
			count++
		}
	}()

	startTime := time.Now()
	// WAIT SIG
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	loop := true
	go func() {
		for loop {
			pss.Send("css")
		}
	}()
	<-quit
	fmt.Println("Shutdown Server ...")
	loop = false
	css.Shutdown()
	endTime := time.Now()
	timeDiff := endTime.Sub(startTime)
	fmt.Printf("%v tests in %v => %v/s", count, timeDiff.Seconds(), int64(float64(count)/timeDiff.Seconds()))
	// consume a message
	/*
		recvMsgReq := &cmq.ReceiveMessageRequest{
			QueueName:          cmq.String(queueName),
			PollingWaitSeconds: cmq.Int(pollingWaitSeconds),
		}
		recvMsgResp, err := client.ReceiveMessage(recvMsgReq)
		if err != nil {
			panic(err)
		}
		data, err := json.Marshal(recvMsgResp)
		if err != nil {
			panic(err)
		}
		//
		fmt.Println(string(data))
		if *recvMsgResp.Code != 0 {
			println("No data")
			return
		}
		recvDeleteMsgReq := &cmq.DeleteMessageRequest{
			QueueName:     cmq.String(queueName),
			ReceiptHandle: recvMsgResp.ReceiptHandle,
		}
		deleteResp, err := client.DeleteMessage(recvDeleteMsgReq)
		data, err = json.Marshal(*deleteResp)
		if err != nil {
			panic(err)
		}
		//
		fmt.Println(string(data))
	*/

}
