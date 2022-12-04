package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

// TmUserHandler 暂存用户接口
func TmUserHandler(w http.ResponseWriter, r *http.Request) {

	res := &JsonResult{}
	if r.Method == http.MethodGet {
		tmUsers, err := getAllTmUser()
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = tmUsers
		}
	} else if r.Method == http.MethodPost {
		openid := r.Header.Get("X-Wx-Openid")
		err := modifyTmUser(r, openid)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		}
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	// 序列化操作，并写入w中
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)

}

// getAllTmUser() 返回所有的暂存用户
func getAllTmUser() ([]model.TmUserModel, error) {
	tmUsers, err := dao.TmUserImp.GetTmUser()
	if err != nil {
		return nil, err
	}
	return tmUsers, err

}

// modifyTmUser() 修改暂存用户信息操作
func modifyTmUser(r *http.Request, openid string) error {
	action, data, err := GetAction(r)
	if err != nil {
		return err
	}
	if action == "add" {
		err = addOneTmUser(openid, data)
	} else if action == "delete" {
		err = deleteOneTmUser(data)
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}

	return err

}

// addOneTmUser() 添加一条暂存用户信息
func addOneTmUser(openid string, data string) error {
	tmUser := model.TmUserModel{}
	if err := json.Unmarshal([]byte(data), &tmUser); err != nil {
		return err
	}
	tmUser.CreateAt = time.Now()
	tmUser.OpenId = openid
	err := dao.TmUserImp.SaveTmUser(&tmUser)
	return err

}

// deleteOneTmUser 删除一条暂存用户信息
func deleteOneTmUser(data string) error {
	user := model.UserModel{}
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return err
	}
	err := dao.TmUserImp.ClearTmUser(&user)
	return err
}
