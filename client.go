package cmq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func init() {
	// for nonce
	rand.Seed(time.Now().Unix())
}

// Client represents a CMQ client
type Client struct {
	secretID   string
	secretKey  string
	region     string
	net        string
	httpClient *http.Client
}

// NewClient creates a new CMQ client
func NewClient(secretID, secretKey, region, net string) *Client {
	return &Client{
		secretID:   secretID,
		secretKey:  secretKey,
		region:     region,
		net:        net,
		httpClient: http.DefaultClient,
	}
}

// WithHTTPClient sets a user-defined HTTP client value returning a Client pointer for chaining
func (c *Client) WithHTTPClient(client *http.Client) *Client {
	c.httpClient = client
	return c
}

// CMQ Public API
// @see https://cloud.tencent.com/document/product/406/5852

// CreateQueue creates a CMQ queue
func (c *Client) CreateQueue(req *CreateQueueRequest) (*CreateQueueResponse, error) {
	body, err := c.do(createQueue, req)
	if err != nil {
		return nil, err
	}

	resp := &CreateQueueResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// ListQueue lists all queue created
func (c *Client) ListQueue(req *ListQueueRequest) (*ListQueueResponse, error) {
	body, err := c.do(listQueue, req)
	if err != nil {
		return nil, err
	}

	resp := &ListQueueResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// GetQueueAttributes gets queue attributes
func (c *Client) GetQueueAttributes(req *GetQueueAttributesRequest) (*GetQueueAttributesResponse, error) {
	body, err := c.do(getQueueAttributes, req)
	if err != nil {
		return nil, err
	}

	resp := &GetQueueAttributesResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// SetQueueAttributes sets queue attributes
func (c *Client) SetQueueAttributes(req *SetQueueAttributesRequest) (*SetQueueAttributesResponse, error) {
	body, err := c.do(setQueueAttributes, req)
	if err != nil {
		return nil, err
	}

	resp := &SetQueueAttributesResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// DeleteQueue deletes a queue
func (c *Client) DeleteQueue(req *DeleteQueueRequest) (*DeleteQueueResponse, error) {
	body, err := c.do(deleteQueue, req)
	if err != nil {
		return nil, err
	}

	resp := &DeleteQueueResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// SendMessage sends a message to the specified queue
func (c *Client) SendMessage(req *SendMessageRequest) (*SendMessageResponse, error) {
	body, err := c.do(sendMessage, req)
	if err != nil {
		return nil, err
	}

	resp := &SendMessageResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// BatchSendMessage sends multiple messages to the specified queue
func (c *Client) BatchSendMessage(req *BatchSendMessageRequest) (*BatchSendMessageResponse, error) {
	body, err := c.do(batchSendMessage, req)
	if err != nil {
		return nil, err
	}

	resp := &BatchSendMessageResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// ReceiveMessage consumes a message from the specified queue
func (c *Client) ReceiveMessage(req *ReceiveMessageRequest) (*ReceiveMessageResponse, error) {
	body, err := c.do(receiveMessage, req)
	if err != nil {
		return nil, err
	}

	resp := &ReceiveMessageResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// BatchReceiveMessage consumes multiple messages from the specified queue
func (c *Client) BatchReceiveMessage(req *BatchReceiveMessageRequest) (*BatchReceiveMessageResponse, error) {
	body, err := c.do(batchReceiveMessage, req)
	if err != nil {
		return nil, err
	}

	resp := &BatchReceiveMessageResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// DeleteMessage delete a message from the specified queue
func (c *Client) DeleteMessage(req *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	body, err := c.do(deleteMessage, req)
	if err != nil {
		return nil, err
	}

	resp := &DeleteMessageResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// BatchDeleteMessage delete multiple messages from the specified queue
func (c *Client) BatchDeleteMessage(req *BatchDeleteMessageRequest) (*BatchDeleteMessageResponse, error) {
	body, err := c.do(batchDeleteMessage, req)
	if err != nil {
		return nil, err
	}

	resp := &BatchDeleteMessageResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// CreateTopic creates a new topic
func (c *Client) CreateTopic(req *CreateTopicRequest) (*CreateTopicResponse, error) {
	body, err := c.do(createTopic, req)
	if err != nil {
		return nil, err
	}

	resp := &CreateTopicResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// SetTopicAttributes updates a topic's attributes
func (c *Client) SetTopicAttributes(req *SetTopicAttributesRequest) (*SetTopicAttributesResponse, error) {
	body, err := c.do(setTopicAttributes, req)
	if err != nil {
		return nil, err
	}

	resp := &SetTopicAttributesResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// ListTopic lists all topics created
func (c *Client) ListTopic(req *ListTopicRequest) (*ListTopicResponse, error) {
	body, err := c.do(listTopic, req)
	if err != nil {
		return nil, err
	}

	resp := &ListTopicResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// GetTopicAttributes gets attributes of the specified topic
func (c *Client) GetTopicAttributes(req *GetTopicAttributesRequest) (*GetTopicAttributesResponse, error) {
	body, err := c.do(getTopicAttributes, req)
	if err != nil {
		return nil, err
	}

	resp := &GetTopicAttributesResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// DeleteTopic delete a specified topic
func (c *Client) DeleteTopic(req *DeleteTopicRequest) (*DeleteTopicResponse, error) {
	body, err := c.do(deleteTopic, req)
	if err != nil {
		return nil, err
	}

	resp := &DeleteTopicResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// PublishMessage publish a message to the specified topic
func (c *Client) PublishMessage(req *PublishMessageRequest) (*PublishMessageResponse, error) {
	body, err := c.do(publishMessage, req)
	if err != nil {
		return nil, err
	}

	resp := &PublishMessageResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// BatchPublishMessage publish multiple messages to the specified topic
func (c *Client) BatchPublishMessage(req *BatchPublishMessageRequest) (*BatchPublishMessageResponse, error) {
	body, err := c.do(batchPublishMessage, req)
	if err != nil {
		return nil, err
	}

	resp := &BatchPublishMessageResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// Subscribe create a new subscription of the specified topic
func (c *Client) Subscribe(req *SubscribeRequest) (*SubscribeResponse, error) {
	body, err := c.do(subscribe, req)
	if err != nil {
		return nil, err
	}

	resp := &SubscribeResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// ListSubscriptionByTopic lists all subscriptions of the specified topic
func (c *Client) ListSubscriptionByTopic(req *ListSubscriptionByTopicRequest) (*ListSubscriptionByTopicResponse, error) {
	body, err := c.do(listSubscriptionByTopic, req)
	if err != nil {
		return nil, err
	}

	resp := &ListSubscriptionByTopicResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// SetSubscriptionAttributes update attributes of the specified subscription
func (c *Client) SetSubscriptionAttributes(req *SetSubscriptionAttributesRequest) (*SetSubscriptionAttributesResponse, error) {
	body, err := c.do(setSubscriptionAttributes, req)
	if err != nil {
		return nil, err
	}

	resp := &SetSubscriptionAttributesResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// GetSubscriptionAttributes get attributes of the specified subscription
func (c *Client) GetSubscriptionAttributes(req *GetSubscriptionAttributesRequest) (*GetSubscriptionAttributesResponse, error) {
	body, err := c.do(getSubscriptionAttributes, req)
	if err != nil {
		return nil, err
	}

	resp := &GetSubscriptionAttributesResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

// Unsubscribe delete a subscription of the specified topic
func (c *Client) Unsubscribe(req *UnsubscribeRequest) (*UnsubscribeResponse, error) {
	body, err := c.do(unsubscribe, req)
	if err != nil {
		return nil, err
	}

	resp := &UnsubscribeResponse{}
	err = json.Unmarshal(body, resp)
	return resp, err
}

func (c *Client) do(action string, req cmqRequest) ([]byte, error) {
	params := req.getParams()
	params["Action"] = action
	params["Region"] = c.region
	params["Timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	params["Nonce"] = strconv.Itoa(rand.Int())
	params["SecretId"] = c.secretID
	params["SignatureMethod"] = hmacSHA256

	scheme := "https"
	if c.net == IntranetAccess {
		scheme = "http"
	}
	host := fmt.Sprintf(getHostPatternByAction(action), params["Region"], c.net)
	src := genPlaintext(params, host)
	params["Signature"] = sign(src, c.secretKey, hmacSHA256)

	api := fmt.Sprintf("%s://%s%s", scheme, host, path)
	data := url.Values{}
	for key, value := range params {
		data.Add(key, value)
	}

	resp, err := c.httpClient.PostForm(api, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
