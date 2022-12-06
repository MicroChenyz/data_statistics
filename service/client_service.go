package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

func ClientHandler(w http.ResponseWriter, r *http.Request) {

	res := &JsonResult{}
	if r.Method == http.MethodGet {
		// Get方法用户获取客户信息
		userIdString := r.URL.Query().Get("userid")
		clientIdString := r.URL.Query().Get("clientid")
		var err error
		var client interface{}
		if userIdString != "" {
			client, err = getClientByUserId(userIdString)
		} else if clientIdString != "" {
			client, err = getClientByClientId(clientIdString)
		} else {
			err = fmt.Errorf("参数不匹配")
		}
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = client
		}
	} else if r.Method == http.MethodPost {
		// Post方法用于存储客户信息
		err := modifyClient(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		}
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)

}

func modifyClient(r *http.Request) error {
	action, data, err := GetAction(r)
	if err != nil {
		return err
	}
	if action == "add" {
		err = addOneClient(data)
	} else if action == "delete" {
		err = deleteOneClient(data)
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}

	return err

}

func addOneClient(data string) error {
	client := model.Client{}
	if err := json.Unmarshal([]byte(data), &client); err != nil {
		return err
	}
	client.CreateTime = time.Now()
	err := dao.ClientImp.SaveClient(&client)
	return err
}

func deleteOneClient(data string) error {
	fmt.Println(data)
	type clientId struct {
		Id int32 `json:"id"`
	}
	var id clientId
	if err := json.Unmarshal([]byte(data), &id); err != nil {
		return err
	}
	err := dao.ClientImp.ClearClient(id.Id)
	return err
}

func getClientByClientId(clientIdString string) (model.Client, error) {
	var client model.Client
	clientId, err := strconv.Atoi(clientIdString)
	if err != nil {
		return client, err
	}
	client, err = dao.ClientImp.GetClientById(int32(clientId))
	return client, err

}

func getClientByUserId(userIdString string) ([]model.ClientResponse, error) {
	clientResponse := make([]model.ClientResponse, 0)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return nil, err
	}
	clientResponse, err = dao.ClientImp.GetClientByUserid(int32(userId))
	return clientResponse, err
}
