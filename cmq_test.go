package cmq_test

import (
	"encoding/json"

	cmq "github.com/zzzhr1990/qcmq"
)

const (
	secretID  = "AKIDNeonaDg6IjBXjrmZBg0GpGWfdPQMRF7r"
	secretKey = "jHd2FJCtZ7VhhVAdjgYSGsKaeVCh4Gk6"
)

var client *cmq.Client

func init() {
	client = cmq.NewClient(secretID, secretKey, cmq.RegionGuangzhou, cmq.InternetAccess)
}

func serialize(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
