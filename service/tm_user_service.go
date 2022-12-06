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

// TmUserHandler 暂存用户接口
func TmUserHandler(w http.ResponseWriter, r *http.Request) {

	res := &JsonResult{}
	pageString := r.URL.Query().Get("page")
	pageSizeString := r.URL.Query().Get("pageSize")
	if r.Method == http.MethodGet {
		// get API 分页查询数据
		var err error
		var data interface{}
		if pageString != "" && pageSizeString != "" {
			data, err = getTmUserByPages(pageString, pageSizeString)
		}
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = data
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
func getTmUserByPages(pageString string, pageSizeString string) (model.TmUserPage, error) {
	page, err := strconv.Atoi(pageString)
	pageSize, err := strconv.Atoi(pageSizeString)
	var tmUserPage model.TmUserPage
	if err != nil {
		return tmUserPage, err
	}
	tmUserPage, err = dao.TmUserImp.GetTmUserByPages(page, pageSize)
	return tmUserPage, err

}

// modifyTmUser() 修改暂存用户信息操作
func modifyTmUser(r *http.Request, openid string) error {
	action, data, err := GetAction(r)
	if err != nil {
		return err
	}
	if action == "add" {
		err = addOneTmUser(openid, data)
	} else if action == "update" {
		err = addOneTmUserToUser(data)
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

// deleteOneTmUser 添加一条暂存用户信息到实际用户信息表中
func addOneTmUserToUser(data string) error {
	user := model.UserModel{}
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return err
	}
	err := dao.TmUserImp.UpdateTmUser(&user)
	return err
}
