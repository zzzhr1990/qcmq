package cmq_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	cmq "github.com/zzzhr1990/qcmq"
)

func createTopic() (*cmq.CreateTopicResponse, error) {
	req := &cmq.CreateTopicRequest{
		TopicName: cmq.String("test-" + strconv.FormatInt(time.Now().Unix(), 10)),
	}
	return client.CreateTopic(req)
}

func TestCreateTopic(t *testing.T) {
	resp, err := createTopic()
	if err != nil {
		panic(fmt.Sprintf("create topic failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("create topic failed: %s", serialize(*resp)))
	}
}

func listTopic() (*cmq.ListTopicResponse, error) {
	return client.ListTopic(&cmq.ListTopicRequest{})
}

func TestListTopic(t *testing.T) {
	resp, err := listTopic()
	if err != nil {
		panic(fmt.Sprintf("list topic failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list topic failed: %s", serialize(*resp)))
	}
}

func setTopicAttributes(topicName string) (*cmq.SetTopicAttributesResponse, error) {
	req := &cmq.SetTopicAttributesRequest{
		TopicName:      cmq.String(topicName),
		MaxMessageSize: cmq.Int(1024),
	}
	return client.SetTopicAttributes(req)
}

func TestSetTopicAttributes(t *testing.T) {
	resp, err := listTopic()
	if err != nil {
		panic(fmt.Sprintf("list topic failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list topic failed: %s", serialize(*resp)))
	}

	topicName := (*resp.TopicList)[0].TopicName
	attrResp, err := setTopicAttributes(topicName)
	if err != nil {
		panic(fmt.Sprintf("set topic attributes failed: %v", err))
	}
	if *attrResp.Code != 0 {
		panic(fmt.Sprintf("set topic attributes failed: %s", serialize(*attrResp)))
	}
}

func getTopicAttributes(topicName string) (*cmq.GetTopicAttributesResponse, error) {
	req := &cmq.GetTopicAttributesRequest{
		TopicName: cmq.String(topicName),
	}
	return client.GetTopicAttributes(req)
}

func TestGetTopicAttributes(t *testing.T) {
	resp, err := listTopic()
	if err != nil {
		panic(fmt.Sprintf("list topic failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list topic failed: %s", serialize(*resp)))
	}

	topicName := (*resp.TopicList)[0].TopicName
	attrResp, err := getTopicAttributes(topicName)
	if err != nil {
		panic(fmt.Sprintf("get topic attributes failed: %v", err))
	}
	if *attrResp.Code != 0 {
		panic(fmt.Sprintf("get topic attributes failed: %s", serialize(*attrResp)))
	}
}

func deleteTopic(topicName string) (*cmq.DeleteTopicResponse, error) {
	req := &cmq.DeleteTopicRequest{
		TopicName: cmq.String(topicName),
	}
	return client.DeleteTopic(req)
}

func TestDeleteTopic(t *testing.T) {
	resp, err := listTopic()
	if err != nil {
		panic(fmt.Sprintf("list topic failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list topic failed: %s", serialize(*resp)))
	}

	topicName := (*resp.TopicList)[0].TopicName
	delResp, err := deleteTopic(topicName)
	if err != nil {
		panic(fmt.Sprintf("delete topic failed: %v", err))
	}
	if *delResp.Code != 0 {
		panic(fmt.Sprintf("delete topic failed: %s", serialize(*delResp)))
	}
}
