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

func StemmingHandler(w http.ResponseWriter, r *http.Request) {

	res := &JsonResult{}
	openid := r.Header.Get("X-Wx-Openid")
	if r.Method == http.MethodGet {
		var err error
		var data interface{}
		idString := r.URL.Query().Get("id")
		if idString != "" {
			// 根据id获取表单信息
			data, err = getStemmingById(idString)
		} else if openid != "" {
			data, err = getStemmingByOpenid(openid)
		}
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = data
		}
	} else if r.Method == http.MethodPost {
		err := modifyStemming(r, openid)
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

func modifyStemming(r *http.Request, openid string) error {
	action, data, err := GetAction(r)
	if err != nil {
		return err
	}
	if action == "add" {
		err = addOneStemming(data, openid)
	} else if action == "update" {
		err = updateOneStemming(data, openid)
	} else {
		err = fmt.Errorf("参数 action : %s 错误", action)
	}
	return err
}

func addOneStemming(data string, openid string) error {
	stemming := model.Stemming{}
	if err := json.Unmarshal([]byte(data), &stemming); err != nil {
		return err
	}
	stemming.OpenId = openid
	stemming.CreateTime = time.Now()
	stemming.ModifyTime = time.Now()
	// 需要计算的内容
	stemming.PeriodOfConsumption = stemming.PreviousPeriodSurplus + stemming.CurrentPeriodAog - stemming.CurrentResidue
	days := stemming.EndTime.Sub(stemming.StartTime).Hours()/24 + 1
	stemming.DailyOfConsumption = stemming.PeriodOfConsumption / days
	stemming.PeriodIronMudConsume = stemming.PeriodOfConsumption * 1000 / stemming.PeriodSumOfIron
	stemming.ConversionPrice = stemming.ContractPrice * stemming.PeriodIronMudConsume / 1000
	err := dao.StemmingImp.SaveStemming(&stemming)
	return err
}

func updateOneStemming(data string, openid string) error {
	stemming := model.Stemming{}
	if err := json.Unmarshal([]byte(data), &stemming); err != nil {
		return err
	}

	stemming.ModifyTime = time.Now()
	stemming.OpenId = openid
	// 需要计算的内容
	stemming.PeriodOfConsumption = stemming.PreviousPeriodSurplus + stemming.CurrentPeriodAog - stemming.CurrentResidue
	days := stemming.EndTime.Sub(stemming.StartTime).Hours() / 24
	stemming.DailyOfConsumption = stemming.PeriodOfConsumption / days
	stemming.PeriodIronMudConsume = stemming.PeriodOfConsumption * 1000 / stemming.PeriodSumOfIron
	stemming.ConversionPrice = stemming.ContractPrice * stemming.PeriodIronMudConsume / 1000
	fmt.Println(stemming)
	err := dao.StemmingImp.UpdateStemmingById(&stemming)
	return err
}

func getStemmingById(idString string) (model.Stemming, error) {
	var err error
	var stemming model.Stemming
	id, err := strconv.Atoi(idString)
	if err != nil {
		return stemming, err
	}
	stemming, err = dao.StemmingImp.GetStemmingById(id)
	return stemming, err
}

func getStemmingByOpenid(openid string) ([]model.StemmingResponse, error) {
	stemmingResponse, err := dao.StemmingImp.GetStemmingByOpenid(openid)
	return stemmingResponse, err
}
