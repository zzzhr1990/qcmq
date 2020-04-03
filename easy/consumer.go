package easy

import (
	"encoding/json"
	"log"
	"sync/atomic"
	"time"

	cmq "github.com/zzzhr1990/qcmq"
)

// Consumer a simple consumer
type Consumer struct {
	client  *cmq.Client
	queue   string
	running int32
}

// NewConsumer proc
func NewConsumer(secretID string, secretKey string, region string, net string, queueName string) *Consumer {
	proc := &Consumer{
		client: cmq.NewClient(secretID, secretKey, region, net),
		queue:  queueName,
	}
	return proc
}

// StartListen lst
func (c *Consumer) StartListen(messageChan chan *Message) {
	atomic.StoreInt32(&c.running, 1)
	go func() {
		for atomic.LoadInt32(&c.running) == 1 {
			recvMsgReq := &cmq.ReceiveMessageRequest{
				QueueName:          cmq.String(c.queue),
				PollingWaitSeconds: cmq.Int(1),
			}
			recvMsgResp, err := c.client.ReceiveMessage(recvMsgReq)
			if err != nil {
				log.Printf("Req error: %v", err)
			} else {
				if *recvMsgResp.Code != 0 {
					if *recvMsgResp.Code != 7000 {
						j, _ := json.Marshal(recvMsgResp)
						log.Printf("Req failed: %v", string(j))
					} else {
						time.Sleep(time.Second * 2)
					}
				} else {
					// recvMsgResp.MessageBody
					msg := &Message{
						Body:          *recvMsgResp.MessageBody,
						ReceiptHandle: *recvMsgResp.ReceiptHandle,
					}
					messageChan <- msg
				}
			}
			// c.client.ReceiveMessage()
		}
		close(messageChan)
	}()

}

// Shutdown close it
func (c *Consumer) Shutdown() {
	// j.
	atomic.StoreInt32(&c.running, 0)
}

// Ack ack an message
func (c *Consumer) Ack(msg *Message) {
	// j.
	// atomic.StoreInt32(&c.running, 0)
	recvDeleteMsgReq := &cmq.DeleteMessageRequest{
		QueueName:     cmq.String(c.queue),
		ReceiptHandle: &msg.ReceiptHandle,
	}
	deleteResp, err := c.client.DeleteMessage(recvDeleteMsgReq)
	if err != nil {
		log.Printf("Req delete message error: %v", err)
		return
	}
	if *deleteResp.Code != 0 {
		j, _ := json.Marshal(deleteResp)
		log.Printf("Req failed: %v", string(j))

	}
}
