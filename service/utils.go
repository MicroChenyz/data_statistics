package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

// GetAction 获取action和data
func GetAction(r *http.Request) (string, string, error) {
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{})
	if err := decoder.Decode(&body); err != nil {
		return "", "", err
	}
	defer r.Body.Close()
	action, ok := body["action"]
	data := body["data"]
	if !ok {
		return "", "", fmt.Errorf("缺少 action 参数")
	}
	actionMsg := fmt.Sprintf("%v", action)
	if body == nil {
		return actionMsg, "", nil
	}
	dataType, _ := json.Marshal(data)
	dataString := string(dataType)

	return actionMsg, dataString, nil

}
