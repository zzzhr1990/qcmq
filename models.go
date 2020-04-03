package cmq

import "strconv"

// specify how to access CMQ (via internet or intranet)
// @see https://cloud.tencent.com/document/api/406/5853
const (
	InternetAccess = "qcloud"
	IntranetAccess = "tencentyun"
)

// specify how to route topic messages
const (
	FilterByTags        = 1
	FilterByBindingKeys = 2
)

// specify subscription protocol
const (
	SubscribeViaHTTP  = "http"
	SubscribeViaQueue = "queue"
)

// specify notify strategy
const (
	BackoffRetry            = "BACKOFF_RETRY"
	ExponentialBackoffRetry = "EXPONENTIAL_DECAY_RETRY"
)

// specify content type
const (
	ContentTypeJSON = "JSON"
	ContentTypeRaw  = "SIMPLIFIED"
)

// @see https://cloud.tencent.com/document/product/406/12667
const (
	RegionGuangzhou       = "gz"
	RegionShanghai        = "sh"
	RegionBeijing         = "bj"
	RegionHongKong        = "hk"
	RegionChengdu         = "cd"
	RegionCanada          = "ca"
	RegionUSAWest         = "usw"
	RegionUSAEast         = "use"
	RegionIndia           = "in"
	RegionThailand        = "th"
	RegionSingapore       = "sg"
	RegionShanghaiFinance = "shjr"
	RegionShenzhenFinance = "szjr"
)

// @see https://cloud.tencent.com/document/product/406/5852
const (
	createQueue               = "CreateQueue"
	listQueue                 = "ListQueue"
	getQueueAttributes        = "GetQueueAttributes"
	setQueueAttributes        = "SetQueueAttributes"
	deleteQueue               = "DeleteQueue"
	sendMessage               = "SendMessage"
	batchSendMessage          = "BatchSendMessage"
	receiveMessage            = "ReceiveMessage"
	batchReceiveMessage       = "BatchReceiveMessage"
	deleteMessage             = "DeleteMessage"
	batchDeleteMessage        = "BatchDeleteMessage"
	createTopic               = "CreateTopic"
	setTopicAttributes        = "SetTopicAttributes"
	listTopic                 = "ListTopic"
	getTopicAttributes        = "GetTopicAttributes"
	deleteTopic               = "DeleteTopic"
	publishMessage            = "PublishMessage"
	batchPublishMessage       = "BatchPublishMessage"
	subscribe                 = "Subscribe"
	listSubscriptionByTopic   = "ListSubscriptionByTopic"
	setSubscriptionAttributes = "SetSubscriptionAttributes"
	getSubscriptionAttributes = "GetSubscriptionAttributes"
	unsubscribe               = "Unsubscribe"
)

type cmqRequest interface {
	getParams() map[string]string
}

// CMQ API Request/Response
// @see https://cloud.tencent.com/document/product/406/5852

// CreateQueueRequest represents a CreateQueue request
type CreateQueueRequest struct {
	QueueName               *string `json:"queueName"`
	MaxMessagePileUpNum     *int    `json:"maxMsgHeapNum"`
	PollingWaitSeconds      *int    `json:"pollingWaitSeconds"`
	VisibilityTimeout       *int    `json:"visibilityTimeout"`
	MaxMessageSize          *int    `json:"maxMsgSize"`
	MessageRetentionSeconds *int    `json:"msgRetentionSeconds"`
	RewindSeconds           *int    `json:"rewindSeconds"`
}

func (r *CreateQueueRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["queueName"] = *r.QueueName
	if r.MaxMessagePileUpNum != nil {
		params["maxMsgHeapNum"] = strconv.Itoa(*r.MaxMessagePileUpNum)
	}
	if r.PollingWaitSeconds != nil {
		params["pollingWaitSeconds"] = strconv.Itoa(*r.PollingWaitSeconds)
	}
	if r.VisibilityTimeout != nil {
		params["visibilityTimeout"] = strconv.Itoa(*r.VisibilityTimeout)
	}
	if r.MaxMessageSize != nil {
		params["maxMsgSize"] = strconv.Itoa(*r.MaxMessageSize)
	}
	if r.MessageRetentionSeconds != nil {
		params["msgRetentionSeconds"] = strconv.Itoa(*r.MessageRetentionSeconds)
	}
	if r.RewindSeconds != nil {
		params["rewindSeconds"] = strconv.Itoa(*r.RewindSeconds)
	}
	return params
}

// CreateQueueResponse represents a CreateQueue response
type CreateQueueResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
	QueueID   *string `json:"queueId"`
}

// ListQueueRequest represents a ListQueue request
type ListQueueRequest struct {
	Keyword *string `json:"searchWord"`
	Offset  *int    `json:"offset"`
	Limit   *int    `json:"limit"`
}

func (r *ListQueueRequest) getParams() map[string]string {
	params := make(map[string]string)
	if r.Keyword != nil {
		params["searchWord"] = *r.Keyword
	}
	if r.Offset != nil {
		params["offset"] = strconv.Itoa(*r.Offset)
	}
	if r.Limit != nil {
		params["limit"] = strconv.Itoa(*r.Limit)
	}
	return params
}

// ListQueueResponse represents a ListQueue response
type ListQueueResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
	Count     *int    `json:"totalCount"`
	QueueList *[]struct {
		QueueID   string `json:"queueId"`
		QueueName string `json:"queueName"`
	} `json:"queueList"`
}

// GetQueueAttributesRequest represents a GetQueueAttributes request
type GetQueueAttributesRequest struct {
	QueueName *string `json:"queueName"`
}

func (r *GetQueueAttributesRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["queueName"] = *r.QueueName
	return params
}

// GetQueueAttributesResponse represents a GetQueueAttributes response
type GetQueueAttributesResponse struct {
	Code                    *int    `json:"code"`
	Message                 *string `json:"message"`
	RequestID               *string `json:"requestId"`
	MaxMessagePileUpNum     *int    `json:"maxMsgHeapNum"`
	PollingWaitSeconds      *int    `json:"pollingWaitSeconds"`
	VisibilityTimeout       *int    `json:"visibilityTimeout"`
	MaxMessageSize          *int    `json:"maxMsgSize"`
	MessageRetentionSeconds *int    `json:"msgRetentionSeconds"`
	CreateTime              *int64  `json:"createTime"`
	LastModifyTime          *int64  `json:"lastModifyTime"`
	ActiveMessageNum        *int    `json:"activeMsgNum"`
	InactiveMessageNum      *int    `json:"inactiveMsgNum"`
	RewindSeconds           *int    `json:"rewindSeconds"`
	RewindMessageNum        *int    `json:"rewindmsgNum"`
	MinMessageTime          *int    `json:"minMsgTime"`
}

// SetQueueAttributesRequest represents a SetQueueAttributes request
type SetQueueAttributesRequest struct {
	QueueName               *string `json:"queueName"`
	MaxMessagePileUpNum     *int    `json:"maxMsgHeapNum"`
	PollingWaitSeconds      *int    `json:"pollingWaitSeconds"`
	VisibilityTimeout       *int    `json:"visibilityTimeout"`
	MaxMessageSize          *int    `json:"maxMsgSize"`
	MessageRetentionSeconds *int    `json:"msgRetentionSeconds"`
	RewindSeconds           *int    `json:"rewindSeconds"`
}

func (r *SetQueueAttributesRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["queueName"] = *r.QueueName
	if r.MaxMessagePileUpNum != nil {
		params["maxMsgHeapNum"] = strconv.Itoa(*r.MaxMessagePileUpNum)
	}
	if r.PollingWaitSeconds != nil {
		params["pollingWaitSeconds"] = strconv.Itoa(*r.PollingWaitSeconds)
	}
	if r.VisibilityTimeout != nil {
		params["visibilityTimeout"] = strconv.Itoa(*r.VisibilityTimeout)
	}
	if r.MaxMessageSize != nil {
		params["maxMsgSize"] = strconv.Itoa(*r.MaxMessageSize)
	}
	if r.MessageRetentionSeconds != nil {
		params["msgRetentionSeconds"] = strconv.Itoa(*r.MessageRetentionSeconds)
	}
	if r.RewindSeconds != nil {
		params["rewindSeconds"] = strconv.Itoa(*r.RewindSeconds)
	}
	return params
}

// SetQueueAttributesResponse represents a SetQueueAttributes response
type SetQueueAttributesResponse struct {
	Code                    *int    `json:"code"`
	Message                 *string `json:"message"`
	RequestID               *string `json:"requestId"`
	MaxMessagePileUpNum     *int    `json:"maxMsgHeapNum"`
	PollingWaitSeconds      *int    `json:"pollingWaitSeconds"`
	VisibilityTimeout       *int    `json:"visibilityTimeout"`
	MaxMessageSize          *int    `json:"maxMsgSize"`
	MessageRetentionSeconds *int    `json:"msgRetentionSeconds"`
	RewindSeconds           *int    `json:"rewindSeconds"`
}

// DeleteQueueRequest represents a DeleteQueue request
type DeleteQueueRequest struct {
	QueueName *string `json:"queueName"`
}

func (r *DeleteQueueRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["queueName"] = *r.QueueName
	return params
}

// DeleteQueueResponse represents a DeleteQueue response
type DeleteQueueResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
}

// SendMessageRequest represents a SendMessage request
type SendMessageRequest struct {
	QueueName    *string `json:"queueName"`
	MessageBody  *string `json:"msgBody"`
	DelaySeconds *int    `json:"delaySeconds"`
}

func (r *SendMessageRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["queueName"] = *r.QueueName
	params["msgBody"] = *r.MessageBody
	if r.DelaySeconds != nil {
		params["delaySeconds"] = strconv.Itoa(*r.DelaySeconds)
	}
	return params
}

// SendMessageResponse represents a SendMessage response
type SendMessageResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
	MessageID *string `json:"msgId"`
}

// BatchSendMessageRequest represents a BatchSendMessage request
type BatchSendMessageRequest struct {
	QueueName    *string   `json:"queueName"`
	Messages     *[]string `json:"msgBody.n"`
	DelaySeconds *int      `json:"delaySeconds"`
}

func (r *BatchSendMessageRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["queueName"] = *r.QueueName
	for i, msg := range *r.Messages {
		params["msgBody."+strconv.Itoa(i)] = msg
	}
	if r.DelaySeconds != nil {
		params["delaySeconds"] = strconv.Itoa(*r.DelaySeconds)
	}
	return params
}

// BatchSendMessageResponse represents a BatchSendMessage response
type BatchSendMessageResponse struct {
	Code        *int    `json:"code"`
	Message     *string `json:"message"`
	RequestID   *string `json:"requestId"`
	MessageList *[]struct {
		MessageID string `json:"msgId"`
	} `json:"msgList"`
}

// ReceiveMessageRequest represents a ReceiveMessage request
type ReceiveMessageRequest struct {
	QueueName          *string `json:"queueName"`
	PollingWaitSeconds *int    `json:"pollingWaitSeconds"`
}

func (r *ReceiveMessageRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["queueName"] = *r.QueueName
	if r.PollingWaitSeconds != nil {
		params["pollingWaitSeconds"] = strconv.Itoa(*r.PollingWaitSeconds)
	}
	return params
}

// ReceiveMessageResponse represents a ReceiveMessage response
type ReceiveMessageResponse struct {
	Code             *int    `json:"code"`
	Message          *string `json:"message"`
	RequestID        *string `json:"requestId"`
	MessageBody      *string `json:"msgBody"`
	MessageID        *string `json:"msgId"`
	ReceiptHandle    *string `json:"receiptHandle"`
	EnqueueTime      *int64  `json:"enqueueTime"`
	FirstDequeueTime *int64  `json:"firstDequeueTime"`
	NextVisibleTime  *int64  `json:"nextVisibleTime"`
	DequeueCount     *int    `json:"dequeueCount"`
}

// BatchReceiveMessageRequest represents a BatchReceiveMessage request
type BatchReceiveMessageRequest struct {
	QueueName          *string `json:"queueName"`
	Count              *int    `json:"numOfMsg"`
	PollingWaitSeconds *int    `json:"pollingWaitSeconds"`
}

func (r *BatchReceiveMessageRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["queueName"] = *r.QueueName
	params["numOfMsg"] = strconv.Itoa(*r.Count)
	if r.PollingWaitSeconds != nil {
		params["pollingWaitSeconds"] = strconv.Itoa(*r.PollingWaitSeconds)
	}
	return params
}

// BatchReceiveMessageResponse represents a BatchReceiveMessage response
type BatchReceiveMessageResponse struct {
	Code        *int    `json:"code"`
	Message     *string `json:"message"`
	RequestID   *string `json:"requestId"`
	MessageList *[]struct {
		MessageBody      string `json:"msgBody"`
		MessageID        string `json:"msgId"`
		ReceiptHandle    string `json:"receiptHandle"`
		EnqueueTime      int64  `json:"enqueueTime"`
		FirstDequeueTime int64  `json:"firstDequeueTime"`
		NextVisibleTime  int64  `json:"nextVisibleTime"`
		DequeueCount     int    `json:"dequeueCount"`
	} `json:"msgInfoList"`
}

// DeleteMessageRequest represents a DeleteMessage request
type DeleteMessageRequest struct {
	QueueName     *string `json:"queueName"`
	ReceiptHandle *string `json:"receiptHandle"`
}

func (r *DeleteMessageRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["queueName"] = *r.QueueName
	params["receiptHandle"] = *r.ReceiptHandle
	return params
}

// DeleteMessageResponse represents a DeleteMessage response
type DeleteMessageResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
}

// BatchDeleteMessageRequest represents a BatchDeleteMessage request
type BatchDeleteMessageRequest struct {
	QueueName      *string   `json:"queueName"`
	ReceiptHandles *[]string `json:"receiptHandle.n"`
}

func (r *BatchDeleteMessageRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["queueName"] = *r.QueueName
	for i, h := range *r.ReceiptHandles {
		params["receiptHandle."+strconv.Itoa(i)] = h
	}
	return params
}

// BatchDeleteMessageResponse represents a BatchDeleteMessage response
type BatchDeleteMessageResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
	ErrorList *[]struct {
		Code          int    `json:"code"`
		Message       string `json:"message"`
		ReceiptHandle string `json:"receiptHandle"`
	} `json:"errorList"`
}

// CreateTopicRequest represents a CreateTopic request
type CreateTopicRequest struct {
	TopicName      *string `json:"topicName"`
	MaxMessageSize *int    `json:"maxMsgSize"`
	FilterType     *int    `json:"filterType"`
}

func (r *CreateTopicRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["topicName"] = *r.TopicName
	if r.MaxMessageSize != nil {
		params["maxMsgSize"] = strconv.Itoa(*r.MaxMessageSize)
	}
	if r.FilterType != nil {
		params["filterType"] = strconv.Itoa(*r.FilterType)
	}
	return params
}

// CreateTopicResponse represents a CreateTopic response
type CreateTopicResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
	TopicID   *string `json:"topicId"`
}

// SetTopicAttributesRequest represents a SetTopicAttributes request
type SetTopicAttributesRequest struct {
	TopicName      *string `json:"topicName"`
	MaxMessageSize *int    `json:"maxMsgSize"`
}

func (r *SetTopicAttributesRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["topicName"] = *r.TopicName
	if r.MaxMessageSize != nil {
		params["maxMsgSize"] = strconv.Itoa(*r.MaxMessageSize)
	}
	return params
}

// SetTopicAttributesResponse represents a SetTopicAttributes response
type SetTopicAttributesResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
}

// ListTopicRequest represents a ListTopic request
type ListTopicRequest struct {
	Keyword *string `json:"searchWord"`
	Offset  *int    `json:"offset"`
	Limit   *int    `json:"limit"`
}

func (r *ListTopicRequest) getParams() map[string]string {
	params := make(map[string]string)
	if r.Keyword != nil {
		params["searchWord"] = *r.Keyword
	}
	if r.Offset != nil {
		params["offset"] = strconv.Itoa(*r.Offset)
	}
	if r.Limit != nil {
		params["limit"] = strconv.Itoa(*r.Limit)
	}
	return params
}

// ListTopicResponse represents a ListTopic response
type ListTopicResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
	Count     *int    `json:"totalCount"`
	TopicList *[]struct {
		TopicID   string `json:"topicId"`
		TopicName string `json:"topicName"`
	} `json:"topicList"`
}

// GetTopicAttributesRequest represents a GetTopicAttributes request
type GetTopicAttributesRequest struct {
	TopicName *string `json:"topicName"`
}

func (r *GetTopicAttributesRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["topicName"] = *r.TopicName
	return params
}

// GetTopicAttributesResponse represents a GetTopicAttributes response
type GetTopicAttributesResponse struct {
	Code                    *int    `json:"code"`
	Message                 *string `json:"message"`
	RequestID               *string `json:"requestId"`
	MessageCount            *int    `json:"msgCount"`
	MaxMessageSize          *int    `json:"maxMsgSize"`
	MessageRetentionSeconds *int    `json:"msgRetentionSeconds"`
	CreateTime              *int64  `json:"createTime"`
	LastModifyTime          *int64  `json:"lastModifyTime"`
	FilterType              *int    `json:"filterType"`
}

// DeleteTopicRequest represents a DeleteTopic request
type DeleteTopicRequest struct {
	TopicName *string `json:"topicName"`
}

func (r *DeleteTopicRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["topicName"] = *r.TopicName
	return params
}

// DeleteTopicResponse represents a DeleteTopic response
type DeleteTopicResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
}

// PublishMessageRequest represents a PublishMessage request
type PublishMessageRequest struct {
	TopicName   *string   `json:"topicName"`
	MessageBody *string   `json:"msgBody"`
	MessageTags *[]string `json:"msgTag.n"`
	RoutingKey  *string   `json:"routingKey"`
}

func (r *PublishMessageRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["topicName"] = *r.TopicName
	params["msgBody"] = *r.MessageBody
	if r.MessageTags != nil {
		for i, tag := range *r.MessageTags {
			params["msgTag."+strconv.Itoa(i)] = tag
		}
	}
	if r.RoutingKey != nil {
		params["routingKey"] = *r.RoutingKey
	}
	return params
}

// PublishMessageResponse represents a PublishMessage response
type PublishMessageResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
	MessageID *string `json:"msgId"`
}

// BatchPublishMessageRequest represents a BatchPublishMessage request
type BatchPublishMessageRequest struct {
	TopicName   *string   `json:"topicName"`
	Messages    *[]string `json:"msgBody.n"`
	MessageTags *[]string `json:"msgTag.n"`
	RoutingKey  *string   `json:"routingKey"`
}

func (r *BatchPublishMessageRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["topicName"] = *r.TopicName
	for i, msg := range *r.Messages {
		params["msgBody."+strconv.Itoa(i)] = msg
	}
	if r.MessageTags != nil {
		for i, tag := range *r.MessageTags {
			params["msgTag."+strconv.Itoa(i)] = tag
		}
	}
	if r.RoutingKey != nil {
		params["routingKey"] = *r.RoutingKey
	}
	return params
}

// BatchPublishMessageResponse represents a BatchPublishMessage response
type BatchPublishMessageResponse struct {
	Code        *int    `json:"code"`
	Message     *string `json:"message"`
	RequestID   *string `json:"requestId"`
	MessageList *[]struct {
		MessageID string `json:"msgId"`
	} `json:"msgList"`
}

// SubscribeRequest represents a Subscribe request
type SubscribeRequest struct {
	TopicName           *string   `json:"topicName"`
	SubscriptionName    *string   `json:"subscriptionName"`
	Protocol            *string   `json:"protocol"`
	Endpoint            *string   `json:"endpoint"`
	NotifyStrategy      *string   `json:"notifyStrategy"`
	NotifyContentFormat *string   `json:"notifyContentFormat"`
	FilterTags          *[]string `json:"filterTag.n"`
	BindingKeys         *[]string `json:"bindingKey.n"`
}

func (r *SubscribeRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["topicName"] = *r.TopicName
	params["subscriptionName"] = *r.SubscriptionName
	params["protocol"] = *r.Protocol
	params["endpoint"] = *r.Endpoint
	if r.NotifyStrategy != nil {
		params["notifyStrategy"] = *r.NotifyStrategy
	}
	if r.NotifyContentFormat != nil {
		params["notifyContentFormat"] = *r.NotifyContentFormat
	}
	if r.FilterTags != nil {
		for i, tag := range *r.FilterTags {
			params["filterTag."+strconv.Itoa(i)] = tag
		}
	}
	for i, key := range *r.BindingKeys {
		params["bindingKey."+strconv.Itoa(i)] = key
	}
	return params
}

// SubscribeResponse represents a Subscribe response
type SubscribeResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
}

// ListSubscriptionByTopicRequest represents a ListSubscriptionByTopic request
type ListSubscriptionByTopicRequest struct {
	TopicName *string `json:"topicName"`
	Keyword   *string `json:"searchWord"`
	Offset    *int    `json:"offset"`
	Limit     *int    `json:"limit"`
}

func (r *ListSubscriptionByTopicRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["topicName"] = *r.TopicName
	if r.Keyword != nil {
		params["searchWord"] = *r.Keyword
	}
	if r.Offset != nil {
		params["offset"] = strconv.Itoa(*r.Offset)
	}
	if r.Limit != nil {
		params["limit"] = strconv.Itoa(*r.Limit)
	}
	return params
}

// ListSubscriptionByTopicResponse represents a ListSubscriptionByTopic response
type ListSubscriptionByTopicResponse struct {
	Code             *int    `json:"code"`
	Message          *string `json:"message"`
	RequestID        *string `json:"requestId"`
	Count            *int    `json:"totalCount"`
	SubscriptionList *[]struct {
		SubscriptionID   string `json:"subscriptionId"`
		SubscriptionName string `json:"subscriptionName"`
		Protocol         string `json:"protocol"`
		Endpoint         string `json:"endpoint"`
	} `json:"subscriptionList"`
}

// SetSubscriptionAttributesRequest represents a SetSubscriptionAttributes request
type SetSubscriptionAttributesRequest struct {
	TopicName           *string   `json:"topicName"`
	SubscriptionName    *string   `json:"subscriptionName"`
	NotifyStrategy      *string   `json:"notifyStrategy"`
	NotifyContentFormat *string   `json:"notifyContentFormat"`
	FilterTags          *[]string `json:"filterTag.n"`
	BindingKeys         *[]string `json:"bindingKey.n"`
}

func (r *SetSubscriptionAttributesRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["topicName"] = *r.TopicName
	params["subscriptionName"] = *r.SubscriptionName
	if r.NotifyStrategy != nil {
		params["notifyStrategy"] = *r.NotifyStrategy
	}
	if r.NotifyContentFormat != nil {
		params["notifyContentFormat"] = *r.NotifyContentFormat
	}
	if r.FilterTags != nil {
		for i, tag := range *r.FilterTags {
			params["filterTag."+strconv.Itoa(i)] = tag
		}
	}
	for i, key := range *r.BindingKeys {
		params["bindingKey."+strconv.Itoa(i)] = key
	}
	return params
}

// SetSubscriptionAttributesResponse represents a SetSubscriptionAttributes response
type SetSubscriptionAttributesResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
}

// GetSubscriptionAttributesRequest represents a GetSubscriptionAttributes request
type GetSubscriptionAttributesRequest struct {
	TopicName        *string `json:"topicName"`
	SubscriptionName *string `json:"subscriptionName"`
}

func (r *GetSubscriptionAttributesRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["topicName"] = *r.TopicName
	params["subscriptionName"] = *r.SubscriptionName
	return params
}

// GetSubscriptionAttributesResponse represents a GetSubscriptionAttributes response
type GetSubscriptionAttributesResponse struct {
	Code                *int      `json:"code"`
	Message             *string   `json:"message"`
	RequestID           *string   `json:"requestId"`
	TopicOwner          *string   `json:"topicOwner"`
	MessageCount        *int      `json:"msgCount"`
	Protocol            *string   `json:"protocol"`
	Endpoint            *string   `json:"endpoint"`
	NotifyStrategy      *string   `json:"notifyStrategy"`
	NotifyContentFormat *string   `json:"notifyContentFormat"`
	CreateTime          *int64    `json:"createTime"`
	LastModifyTime      *int64    `json:"lastModifyTime"`
	BindingKeys         *[]string `json:"bindingKey"`
}

// UnsubscribeRequest represents a Unsubscribe request
type UnsubscribeRequest struct {
	TopicName        *string `json:"topicName"`
	SubscriptionName *string `json:"subscriptionName"`
}

func (r *UnsubscribeRequest) getParams() map[string]string {
	params := make(map[string]string)
	params["topicName"] = *r.TopicName
	params["subscriptionName"] = *r.SubscriptionName
	return params
}

// UnsubscribeResponse represents a Unsubscribe response
type UnsubscribeResponse struct {
	Code      *int    `json:"code"`
	Message   *string `json:"message"`
	RequestID *string `json:"requestId"`
}
