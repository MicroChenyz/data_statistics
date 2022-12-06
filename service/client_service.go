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
	openid := r.Header.Get("X-Wx-Openid")
	if r.Method == http.MethodGet {
		var err error
		var data interface{}
		pageString := r.URL.Query().Get("page")
		pageSizeString := r.URL.Query().Get("pageSize")
		idString := r.URL.Query().Get("id")
		// 分页获取所有客户信息
		if pageString != "" && pageSizeString != "" {
			data, err = getClientByPages(pageString, pageSizeString)
		} else if openid != "" {
			data, err = getClientByOpenId(openid)
		} else if idString != "" {
			data, err = getClientByClientId(idString)
		} else {
			data, err = getAllClient()
		}
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = data
		}
	} else if r.Method == http.MethodPost {
		// Post方法用于存储客户信息
		err := modifyClient(r, openid)
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

func modifyClient(r *http.Request, openid string) error {
	action, data, err := GetAction(r)
	if err != nil {
		return err
	}
	if action == "add" {
		err = addOneClient(data, openid)
	} else if action == "delete" {
		err = deleteOneClient(data, openid)
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}

	return err

}

func addOneClient(data string, openid string) error {
	client := model.Client{}
	if err := json.Unmarshal([]byte(data), &client); err != nil {
		return err
	}
	client.OpenId = openid
	client.CreateTime = time.Now()
	err := dao.ClientImp.SaveClient(&client)
	return err
}

func deleteOneClient(data string, openid string) error {
	type clientId struct {
		Id int `json:"id"`
	}
	var err error
	if data == "null" {
		err = dao.ClientImp.DeleteClientByOpenId(openid)
	} else {
		var clientid clientId
		if err := json.Unmarshal([]byte(data), &clientid); err != nil {
			return err
		}
		err = dao.ClientImp.DeleteClientByClientId(clientid.Id)
	}
	return err
}

func getClientByClientId(idString string) (model.Client, error) {
	var client model.Client
	id, err := strconv.Atoi(idString)
	if err != nil {
		return client, err
	}
	client, err = dao.ClientImp.GetClientByClientId(id)
	return client, err
}

func getClientByOpenId(openid string) ([]model.ClientResponse, error) {
	clientResponse, err := dao.ClientImp.GetClientByOpenId(openid)
	return clientResponse, err
}

func getClientByPages(pageString string, pageSizeString string) (model.ClientPage, error) {
	var err error
	var clientPage model.ClientPage
	page, err := strconv.Atoi(pageString)
	pageSize, err := strconv.Atoi(pageSizeString)
	if err != nil {
		return clientPage, err
	}
	clientPage, err = dao.ClientImp.GetAllClientByPages(page, pageSize)
	return clientPage, err
}

func getAllClient() ([]model.Client, error) {
	client, err := dao.ClientImp.GetAllClient()
	return client, err
}
