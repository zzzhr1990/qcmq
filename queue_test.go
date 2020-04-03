package cmq_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	cmq "github.com/zzzhr1990/qcmq"
)

func createQueue() (*cmq.CreateQueueResponse, error) {
	queueName := "test-" + strconv.FormatInt(time.Now().Unix(), 10)
	req := &cmq.CreateQueueRequest{
		QueueName: cmq.String(queueName),
	}
	return client.CreateQueue(req)
}

func TestCreateQueue(t *testing.T) {
	resp, err := createQueue()
	if err != nil {
		panic(fmt.Sprintf("create queue failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("create queue failed: %s", serialize(*resp)))
	}
}

func listQueue() (*cmq.ListQueueResponse, error) {
	return client.ListQueue(&cmq.ListQueueRequest{})
}

func TestListQueue(t *testing.T) {
	resp, err := listQueue()
	if err != nil {
		panic(fmt.Sprintf("list queue failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list queue failed: %s", serialize(*resp)))
	}
}

func getQueueAttributes(queueName string) (*cmq.GetQueueAttributesResponse, error) {
	req := &cmq.GetQueueAttributesRequest{
		QueueName: cmq.String(queueName),
	}
	return client.GetQueueAttributes(req)
}

func TestGetQueueAttributes(t *testing.T) {
	resp, err := listQueue()
	if err != nil {
		panic(fmt.Sprintf("list queue failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list queue failed: %s", serialize(*resp)))
	}

	queueName := (*resp.QueueList)[0].QueueName
	attrResp, err := getQueueAttributes(queueName)
	if err != nil {
		panic(fmt.Sprintf("get queue attributes failed: %v", err))
	}
	if *attrResp.Code != 0 {
		panic(fmt.Sprintf("get queue attributes failed: %s", serialize(*attrResp)))
	}
}

func setQueueAttributes(queueName string) (*cmq.SetQueueAttributesResponse, error) {
	req := &cmq.SetQueueAttributesRequest{
		QueueName:          cmq.String(queueName),
		PollingWaitSeconds: cmq.Int(3),
	}
	return client.SetQueueAttributes(req)
}

func TestSetQueueAttributes(t *testing.T) {
	resp, err := listQueue()
	if err != nil {
		panic(fmt.Sprintf("list queue failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list queue failed: %s", serialize(*resp)))
	}

	queueName := (*resp.QueueList)[0].QueueName
	attrResp, err := setQueueAttributes(queueName)
	if err != nil {
		panic(fmt.Sprintf("set queue attributes failed: %v", err))
	}
	if *attrResp.Code != 0 {
		panic(fmt.Sprintf("set queue attributes failed: %s", serialize(*attrResp)))
	}
}

func deleteQueue(queueName string) (*cmq.DeleteQueueResponse, error) {
	req := &cmq.DeleteQueueRequest{
		QueueName: cmq.String(queueName),
	}
	return client.DeleteQueue(req)
}

func TestDeleteQueue(t *testing.T) {
	resp, err := listQueue()
	if err != nil {
		panic(fmt.Sprintf("list queue failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list queue failed: %s", serialize(*resp)))
	}

	queueName := (*resp.QueueList)[0].QueueName
	delResp, err := deleteQueue(queueName)
	if err != nil {
		panic(fmt.Sprintf("delete queue failed: %v", err))
	}
	if *delResp.Code != 0 {
		panic(fmt.Sprintf("delete queue failed: %s", serialize(*delResp)))
	}
}
