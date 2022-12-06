package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

// UserHandler 实际用户接口
func UserHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	openid := r.Header.Get("X-Wx-Openid")
	if r.Method == http.MethodGet {
		var err error
		var data interface{}
		pageString := r.URL.Query().Get("page")
		pageSizeString := r.URL.Query().Get("pageSize")
		if pageString != "" && pageSizeString != "" {
			// 分页查询
			data, err = getUserByPages(pageString, pageSizeString)
		} else if openid != "" {
			// openid查询
			data, err = getUserByOpenId(openid)
		} else {
			// 全量查询
			data, err = getAllUser()
		}
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = data
		}
	} else if r.Method == http.MethodPost {
		err := deleteUser(r, openid)
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
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

// getUserByOpenId 通过openid获取用户信息
func getUserByOpenId(openid string) (model.UserModel, error) {
	user, err := dao.UserImp.FindUserByOpenId(openid)
	return user, err
}

// deleteUser 删除一条用户记录
func deleteUser(r *http.Request, openid string) error {
	action, _, err := GetAction(r)
	if err != nil {
		return err
	}
	if action == "delete" {
		err = dao.UserImp.DeleteUser(openid)
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}
	return err
}

// getAllUser 获取指定分页的记录
func getUserByPages(pageString string, pageSizeString string) (model.UserPage, error) {
	var err error
	var userPage model.UserPage
	page, err := strconv.Atoi(pageString)
	pageSize, err := strconv.Atoi(pageSizeString)
	if err != nil {
		return userPage, err
	}
	userPage, err = dao.UserImp.FindUserByPages(page, pageSize)
	return userPage, err
}

func getAllUser() ([]model.UserModel, error) {
	var err error
	var users = make([]model.UserModel, 0)
	users, err = dao.UserImp.FindAllUser()
	return users, err
}
