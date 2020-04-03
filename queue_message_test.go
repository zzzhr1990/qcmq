package cmq_test

import (
	"fmt"
	"testing"

	"github.com/zzzhr1990/qcmq"
)

func sendMessage(queueName string) (*cmq.SendMessageResponse, error) {
	req := &cmq.SendMessageRequest{
		QueueName:   cmq.String(queueName),
		MessageBody: cmq.String("hello"),
	}
	return client.SendMessage(req)
}

func TestSendMessage(t *testing.T) {
	resp, err := listQueue()
	if err != nil {
		panic(fmt.Sprintf("list queue failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list queue failed: %s", serialize(*resp)))
	}

	queueName := (*resp.QueueList)[0].QueueName
	sendResp, err := sendMessage(queueName)
	if err != nil {
		panic(fmt.Sprintf("send message failed: %v", err))
	}
	if *sendResp.Code != 0 {
		panic(fmt.Sprintf("send message failed: %s", serialize(*sendResp)))
	}
}

func batchSendMessage(queueName string) (*cmq.BatchSendMessageResponse, error) {
	req := &cmq.BatchSendMessageRequest{
		QueueName: cmq.String(queueName),
		Messages:  cmq.StringSlice([]string{"hello", "world"}),
	}
	return client.BatchSendMessage(req)
}

func TestBatchSendMessage(t *testing.T) {
	resp, err := listQueue()
	if err != nil {
		panic(fmt.Sprintf("list queue failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list queue failed: %s", serialize(*resp)))
	}

	queueName := (*resp.QueueList)[0].QueueName
	sendResp, err := batchSendMessage(queueName)
	if err != nil {
		panic(fmt.Sprintf("batch send message failed: %v", err))
	}
	if *sendResp.Code != 0 {
		panic(fmt.Sprintf("batch send message failed: %s", serialize(*sendResp)))
	}
}

func receiveMessage(queueName string) (*cmq.ReceiveMessageResponse, error) {
	req := &cmq.ReceiveMessageRequest{
		QueueName: cmq.String(queueName),
	}
	return client.ReceiveMessage(req)
}

func TestReceiveMessage(t *testing.T) {
	resp, err := listQueue()
	if err != nil {
		panic(fmt.Sprintf("list queue failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list queue failed: %s", serialize(*resp)))
	}

	queueName := (*resp.QueueList)[0].QueueName
	recvResp, err := receiveMessage(queueName)
	if err != nil {
		panic(fmt.Sprintf("receive message failed: %v", err))
	}
	if *recvResp.Code != 0 {
		panic(fmt.Sprintf("receive message failed: %s", serialize(*recvResp)))
	}
}

func batchReceiveMessage(queueName string) (*cmq.BatchReceiveMessageResponse, error) {
	req := &cmq.BatchReceiveMessageRequest{
		QueueName: cmq.String(queueName),
		Count:     cmq.Int(2),
	}
	return client.BatchReceiveMessage(req)
}

func TestBatchReceiveMessage(t *testing.T) {
	resp, err := listQueue()
	if err != nil {
		panic(fmt.Sprintf("list queue failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list queue failed: %s", serialize(*resp)))
	}

	queueName := (*resp.QueueList)[0].QueueName
	recvResp, err := batchReceiveMessage(queueName)
	if err != nil {
		panic(fmt.Sprintf("batch receive message failed: %v", err))
	}
	if *recvResp.Code != 0 {
		panic(fmt.Sprintf("batch receive message failed: %s", serialize(*recvResp)))
	}
}

func deleteMessage(queueName string, handle string) (*cmq.DeleteMessageResponse, error) {
	req := &cmq.DeleteMessageRequest{
		QueueName:     cmq.String(queueName),
		ReceiptHandle: cmq.String(handle),
	}
	return client.DeleteMessage(req)
}

func TestDeleteMessage(t *testing.T) {
	resp, err := listQueue()
	if err != nil {
		panic(fmt.Sprintf("list queue failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list queue failed: %s", serialize(*resp)))
	}

	queueName := (*resp.QueueList)[0].QueueName
	recvResp, err := receiveMessage(queueName)
	if err != nil {
		panic(fmt.Sprintf("receive message failed: %v", err))
	}
	if *recvResp.Code != 0 {
		panic(fmt.Sprintf("receive message failed: %s", serialize(*recvResp)))
	}

	handle := *recvResp.ReceiptHandle
	delResp, err := deleteMessage(queueName, handle)
	if err != nil {
		panic(fmt.Sprintf("delete message failed: %v", err))
	}
	if *delResp.Code != 0 {
		panic(fmt.Sprintf("delete message failed: %s", serialize(*delResp)))
	}
}

func batchDeleteMessage(queueName string, handles []string) (*cmq.BatchDeleteMessageResponse, error) {
	req := &cmq.BatchDeleteMessageRequest{
		QueueName:      cmq.String(queueName),
		ReceiptHandles: cmq.StringSlice(handles),
	}
	return client.BatchDeleteMessage(req)
}

func TestBatchDeleteMessage(t *testing.T) {
	resp, err := listQueue()
	if err != nil {
		panic(fmt.Sprintf("list queue failed: %v", err))
	}
	if *resp.Code != 0 {
		panic(fmt.Sprintf("list queue failed: %s", serialize(*resp)))
	}

	queueName := (*resp.QueueList)[0].QueueName
	recvResp, err := batchReceiveMessage(queueName)
	if err != nil {
		panic(fmt.Sprintf("batch receive message failed: %v", err))
	}
	if *recvResp.Code != 0 {
		panic(fmt.Sprintf("batch receive message failed: %s", serialize(*recvResp)))
	}

	var handles []string
	for _, msg := range *recvResp.MessageList {
		handles = append(handles, msg.ReceiptHandle)
	}
	delResp, err := batchDeleteMessage(queueName, handles)
	if err != nil {
		panic(fmt.Sprintf("batch delete message failed: %v", err))
	}
	if *delResp.Code != 0 {
		panic(fmt.Sprintf("batch delete message failed: %s", serialize(*delResp)))
	}
}
