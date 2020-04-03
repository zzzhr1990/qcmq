package cmq_test

import (
	"fmt"
	"testing"

	"github.com/zzzhr1990/qcmq"
)

func publishMessage(topicName string) (*cmq.PublishMessageResponse, error) {
	req := &cmq.PublishMessageRequest{
		TopicName:   cmq.String(topicName),
		MessageBody: cmq.String("hello"),
	}
	return client.PublishMessage(req)
}

func TestPublishMessage(t *testing.T) {
	resp, err := listTopic()
	if err != nil {
		panic(fmt.Sprintf("list topic failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list topic failed: %s", serialize(*resp)))
	}

	topicName := (*resp.TopicList)[0].TopicName
	pubResp, err := publishMessage(topicName)
	if err != nil {
		panic(fmt.Sprintf("publish message failed: %v", err))
	}
	if *pubResp.Code != 0 {
		panic(fmt.Sprintf("publish message failed: %s", serialize(*pubResp)))
	}
}

func batchPublishMessage(topicName string) (*cmq.BatchPublishMessageResponse, error) {
	req := &cmq.BatchPublishMessageRequest{
		TopicName: cmq.String(topicName),
		Messages:  cmq.StringSlice([]string{"hello", "world"}),
	}
	return client.BatchPublishMessage(req)
}

func TestBatchPublishMessage(t *testing.T) {
	resp, err := listTopic()
	if err != nil {
		panic(fmt.Sprintf("list topic failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list topic failed: %s", serialize(*resp)))
	}

	topicName := (*resp.TopicList)[0].TopicName
	pubResp, err := batchPublishMessage(topicName)
	if err != nil {
		panic(fmt.Sprintf("batch publish message failed: %v", err))
	}
	if *pubResp.Code != 0 {
		panic(fmt.Sprintf("batch publish message failed: %s", serialize(*pubResp)))
	}
}
