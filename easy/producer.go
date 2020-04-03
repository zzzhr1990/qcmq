package easy

import (
	"encoding/json"
	"errors"

	cmq "github.com/zzzhr1990/qcmq"
)

// Producer a producer
type Producer struct {
	client *cmq.Client
	queue  string
}

// NewProducer proc
func NewProducer(secretID string, secretKey string, region string, net string, queueName string) *Producer {
	proc := &Producer{
		client: cmq.NewClient(secretID, secretKey, region, net),
		queue:  queueName,
	}
	return proc
}

// Send send to ps
func (p *Producer) Send(message string) error {
	//
	sendMsgReq := &cmq.SendMessageRequest{
		QueueName:   cmq.String(p.queue),
		MessageBody: cmq.String(message),
	}
	sendMsgResp, err := p.client.SendMessage(sendMsgReq)
	if err != nil {
		return err
	}
	if *sendMsgResp.Code != 0 {
		err, _ := json.Marshal(*sendMsgResp)
		return errors.New("Cannot send message: " + string(err))
	}
	/*
		data, err := json.Marshal(*sendMsgResp)
		if err != nil {
			panic(err)
		}
	*/
	return nil
}
