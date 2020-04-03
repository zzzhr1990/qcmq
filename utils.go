package cmq

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"sort"
	"strings"
)

const (
	hmacSHA256       = "HmacSHA256"
	hmacSHA1         = "HmacSHA1"
	method           = "POST"
	queueHostPattern = "cmq-queue-%s.api.%s.com"
	topicHostPattern = "cmq-topic-%s.api.%s.com"
	path             = "/v2/index.php"
)

var m map[string]string

func init() {
	m = make(map[string]string)
	m[createQueue] = queueHostPattern
	m[listQueue] = queueHostPattern
	m[getQueueAttributes] = queueHostPattern
	m[setQueueAttributes] = queueHostPattern
	m[deleteQueue] = queueHostPattern
	m[sendMessage] = queueHostPattern
	m[batchSendMessage] = queueHostPattern
	m[receiveMessage] = queueHostPattern
	m[batchReceiveMessage] = queueHostPattern
	m[deleteMessage] = queueHostPattern
	m[batchDeleteMessage] = queueHostPattern
	m[createTopic] = topicHostPattern
	m[setTopicAttributes] = topicHostPattern
	m[listTopic] = topicHostPattern
	m[getTopicAttributes] = topicHostPattern
	m[deleteTopic] = topicHostPattern
	m[publishMessage] = topicHostPattern
	m[batchPublishMessage] = topicHostPattern
	m[subscribe] = topicHostPattern
	m[listSubscriptionByTopic] = topicHostPattern
	m[setSubscriptionAttributes] = topicHostPattern
	m[getSubscriptionAttributes] = topicHostPattern
	m[unsubscribe] = topicHostPattern
}

// String returns a pointer to the string value passed in
func String(s string) *string {
	return &s
}

// StringSlice returns a pointer to the string slice value passed in
func StringSlice(s []string) *[]string {
	return &s
}

// Int returns a pointer to the int value passed in
func Int(i int) *int {
	return &i
}

// Int64 returns a pointer to the int64 value passed in
func Int64(i int64) *int64 {
	return &i
}

func genPlaintext(params map[string]string, host string) string {
	query := sortParams(params)
	return fmt.Sprintf("%s%s%s?%s", method, host, path, query)
}

func sortParams(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	values := make([]string, 0, len(keys))
	for _, key := range keys {
		values = append(values, fmt.Sprintf("%s=%s", key, params[key]))
	}

	return strings.Join(values, "&")
}

func sign(src string, key string, method string) string {
	var h hash.Hash
	switch method {
	case hmacSHA256:
		h = hmac.New(sha256.New, []byte(key))
	default:
		h = hmac.New(sha1.New, []byte(key))
	}

	h.Write([]byte(src))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func getHostPatternByAction(action string) string {
	return m[action]
}
