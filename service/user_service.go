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
	if r.Method == http.MethodGet {
		openid := r.URL.Query().Get("openid")
		var err error
		var user interface{}
		if openid == "" {
			pageString := r.URL.Query().Get("page")
			pageSizeString := r.URL.Query().Get("pageSize")
			page, err := strconv.Atoi(pageString)
			pageSize, err := strconv.Atoi(pageSizeString)
			if err == nil {
				user, err = getAllUser(page, pageSize)
			}
		} else {
			user, err = getUserByOpenId(openid)
		}
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = user
		}
	} else if r.Method == http.MethodPost {
		err := deleteUser(r)
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
func deleteUser(r *http.Request) error {
	action, data, err := GetAction(r)
	if err != nil {
		return err
	}
	if action == "delete" {
		user := model.UserModel{}
		if err := json.Unmarshal([]byte(data), &user); err != nil {
			return err
		}
		err = dao.UserImp.ClearUser(user.Id)
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}
	return err
}

// getAllUser 获取指定分页的记录
func getAllUser(page int, pageSize int) (model.Pages, error) {
	pages, err := dao.UserImp.FindAllUser(page, pageSize)
	return pages, err
}
