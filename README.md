# Tencent Cloud CMQ SDK For Go

## Getting Started

```go
package main

import (
	"encoding/json"
	"fmt"

	cmq "github.com/zzzhr1990/qcmq"
)

const (
	secretID           = "YOUR_SECRET_ID"
	secretKey          = "YOUR_SECRET_KEY"
	queueName          = "YOUR_QUEUE_NAME"
	pollingWaitSeconds = 1
)

func main() {
	// create a CMQ client
	client := cmq.NewClient(secretID, secretKey, cmq.RegionGuangzhou, cmq.InternetAccess)

	// send a message
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

	// consume a message
	recvMsgReq := &cmq.ReceiveMessageRequest{
		QueueName:          cmq.String(queueName),
		PollingWaitSeconds: cmq.Int(pollingWaitSeconds),
	}
	recvMsgResp, err := client.ReceiveMessage(recvMsgReq)
	if err != nil {
		panic(err)
	}
	data, err = json.Marshal(*recvMsgResp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

```

## Documetation

- [Introduction](https://cloud.tencent.com/document/api/406/5851)
- [API Overview](https://cloud.tencent.com/document/api/406/5852)

## APIs

- [x] [CreateQueue](https://cloud.tencent.com/document/api/431/5832)
- [x] [ListQueue](https://cloud.tencent.com/document/api/431/5833)
- [x] [GetQueueAttributes](https://cloud.tencent.com/document/api/431/5834)
- [x] [SetQueueAttributes](https://cloud.tencent.com/document/api/431/5835)
- [x] [DeleteQueue](https://cloud.tencent.com/document/api/431/5836)
- [x] [SendMessage](https://cloud.tencent.com/document/api/431/5837)
- [x] [BatchSendMessage](https://cloud.tencent.com/document/api/431/5838)
- [x] [ReceiveMessage](https://cloud.tencent.com/document/api/431/5839)
- [x] [BatchReceiveMessage](https://cloud.tencent.com/document/api/431/5924)
- [x] [DeleteMessage](https://cloud.tencent.com/document/api/431/5840)
- [x] [BatchDeleteMessage](https://cloud.tencent.com/document/api/431/5841)
- [x] [CreateTopic](https://cloud.tencent.com/document/api/406/7405)
- [x] [SetTopicAttributes](https://cloud.tencent.com/document/api/406/7406)
- [x] [ListTopic](https://cloud.tencent.com/document/api/406/7407)
- [x] [GetTopicAttributes](https://cloud.tencent.com/document/api/406/7408)
- [x] [DeleteTopic](https://cloud.tencent.com/document/api/406/7409)
- [x] [PublishMessage](https://cloud.tencent.com/document/api/406/7411)
- [x] [BatchPublishMessage](https://cloud.tencent.com/document/api/406/7412)
- [x] [Subscribe](https://cloud.tencent.com/document/api/406/7414)
- [x] [ListSubscriptionByTopic](https://cloud.tencent.com/document/api/406/7415)
- [x] [SetSubscriptionAttributes](https://cloud.tencent.com/document/api/406/7416)
- [x] [GetSubscriptionAttributes](https://cloud.tencent.com/document/api/406/7418)
- [x] [Unsubscribe](https://cloud.tencent.com/document/api/406/7417)
