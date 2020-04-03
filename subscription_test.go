package cmq_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/zzzhr1990/qcmq"
)

func subscribe(topicName, queueName string) (*cmq.SubscribeResponse, error) {
	req := &cmq.SubscribeRequest{
		TopicName:        cmq.String(topicName),
		SubscriptionName: cmq.String("test-" + strconv.FormatInt(time.Now().Unix(), 10)),
		Protocol:         cmq.String(cmq.SubscribeViaQueue),
		Endpoint:         cmq.String(queueName),
		BindingKeys:      cmq.StringSlice([]string{"key1"}),
	}
	return client.Subscribe(req)
}

func TestSubscribe(t *testing.T) {
	topicResp, err := listTopic()
	if err != nil {
		panic(fmt.Sprintf("list topic failed: %v", err))
	}
	if *topicResp.Code != 0 {
		panic(fmt.Sprintf("list topic failed: %s", serialize(*topicResp)))
	}

	queueResp, err := listQueue()
	if err != nil {
		panic(fmt.Sprintf("list queue failed: %v", err))
	}
	if *queueResp.Code != 0 {
		panic(fmt.Sprintf("list queue failed: %s", serialize(*queueResp)))
	}

	topicName := (*topicResp.TopicList)[0].TopicName
	queueName := (*queueResp.QueueList)[0].QueueName
	resp, err := subscribe(topicName, queueName)
	if err != nil {
		panic(fmt.Sprintf("subscribe failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("subscribe failed: %s", serialize(*resp)))
	}
}

func listSubscriptionByTopic(topicName string) (*cmq.ListSubscriptionByTopicResponse, error) {
	req := &cmq.ListSubscriptionByTopicRequest{
		TopicName: cmq.String(topicName),
	}
	return client.ListSubscriptionByTopic(req)
}

func TestListSubscriptionByTopic(t *testing.T) {
	resp, err := listTopic()
	if err != nil {
		panic(fmt.Sprintf("list topic failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list topic failed: %s", serialize(*resp)))
	}

	topicName := (*resp.TopicList)[0].TopicName
	listResp, err := listSubscriptionByTopic(topicName)
	if err != nil {
		panic(fmt.Sprintf("list subscription by topic failed: %v", err))
	}
	if *listResp.Code != 0 {
		panic(fmt.Sprintf("list subscription by topic failed: %s", serialize(*listResp)))
	}
}

func setSubscriptionAttributes(topicName, subscriptionName string) (*cmq.SetSubscriptionAttributesResponse, error) {
	req := &cmq.SetSubscriptionAttributesRequest{
		TopicName:        cmq.String(topicName),
		SubscriptionName: cmq.String(subscriptionName),
		BindingKeys:      cmq.StringSlice([]string{"key2"}),
	}
	return client.SetSubscriptionAttributes(req)
}

func TestSetSubscriptionAttributes(t *testing.T) {
	resp, err := listTopic()
	if err != nil {
		panic(fmt.Sprintf("list topic failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list topic failed: %s", serialize(*resp)))
	}

	topicName := (*resp.TopicList)[0].TopicName
	listResp, err := listSubscriptionByTopic(topicName)
	if err != nil {
		panic(fmt.Sprintf("list subscription by topic failed: %v", err))
	}
	if *listResp.Code != 0 {
		panic(fmt.Sprintf("list subscription by topic failed: %s", serialize(*listResp)))
	}

	subscriptionName := (*listResp.SubscriptionList)[0].SubscriptionName

	attrResp, err := setSubscriptionAttributes(topicName, subscriptionName)
	if err != nil {
		panic(fmt.Sprintf("set subscription attributes failed: %v", err))
	}
	if *attrResp.Code != 0 {
		panic(fmt.Sprintf("set subscription attributes failed: %s", serialize(*attrResp)))
	}
}

func getSubscriptionAttributes(topicName, subscriptionName string) (*cmq.GetSubscriptionAttributesResponse, error) {
	req := &cmq.GetSubscriptionAttributesRequest{
		TopicName:        cmq.String(topicName),
		SubscriptionName: cmq.String(subscriptionName),
	}
	return client.GetSubscriptionAttributes(req)
}

func TestGetSubscriptionAttributes(t *testing.T) {
	resp, err := listTopic()
	if err != nil {
		panic(fmt.Sprintf("list topic failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list topic failed: %s", serialize(*resp)))
	}

	topicName := (*resp.TopicList)[0].TopicName
	listResp, err := listSubscriptionByTopic(topicName)
	if err != nil {
		panic(fmt.Sprintf("list subscription by topic failed: %v", err))
	}
	if *listResp.Code != 0 {
		panic(fmt.Sprintf("list subscription by topic failed: %s", serialize(*listResp)))
	}

	subscriptionName := (*listResp.SubscriptionList)[0].SubscriptionName

	attrResp, err := getSubscriptionAttributes(topicName, subscriptionName)
	if err != nil {
		panic(fmt.Sprintf("get subscription attributes failed: %v", err))
	}
	if *attrResp.Code != 0 {
		panic(fmt.Sprintf("get subscription attributes failed: %s", serialize(*attrResp)))
	}
}

func unsubscribe(topicName, subscriptionName string) (*cmq.UnsubscribeResponse, error) {
	req := &cmq.UnsubscribeRequest{
		TopicName:        cmq.String(topicName),
		SubscriptionName: cmq.String(subscriptionName),
	}
	return client.Unsubscribe(req)
}

func TestUnsubscribe(t *testing.T) {
	topicResp, err := listTopic()
	if err != nil {
		panic(fmt.Sprintf("list topic failed: %v", err))
	}
	if *topicResp.Code != 0 {
		panic(fmt.Sprintf("list topic failed: %s", serialize(*topicResp)))
	}

	topicName := (*topicResp.TopicList)[0].TopicName
	listResp, err := listSubscriptionByTopic(topicName)
	if err != nil {
		panic(fmt.Sprintf("list subscription by topic failed: %v", err))
	}
	if *listResp.Code != 0 {
		panic(fmt.Sprintf("list subscription by topic failed: %s", serialize(*listResp)))
	}

	subscriptionName := (*listResp.SubscriptionList)[0].SubscriptionName

	resp, err := unsubscribe(topicName, subscriptionName)
	if err != nil {
		panic(fmt.Sprintf("unsubscribe failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("unsubscribe failed: %s", serialize(*resp)))
	}
}
