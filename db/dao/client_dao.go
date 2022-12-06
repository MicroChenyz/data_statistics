package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const clientTableName = "client"

func (c ClientInterfaceImp) SaveClient(client *model.Client) error {
	//TODO implement me
	var err error
	cli := db.Get()
	err = cli.Table(clientTableName).Save(client).Error
	return err
}

func (c ClientInterfaceImp) DeleteClientByClientId(id int) error {
	//TODO implement me
	var err error
	cli := db.Get()
	err = cli.Table(clientTableName).Delete(&model.Client{Id: id}).Error
	return err
}

func (c ClientInterfaceImp) DeleteClientByOpenId(openid string) error {
	var err error
	cli := db.Get()
	err = cli.Table(clientTableName).Where("openid=?", openid).Delete(&model.Client{}).Error
	return err
}

func (c ClientInterfaceImp) GetClientByOpenId(openid string) ([]model.ClientResponse, error) {
	//TODO implement me
	var err error
	var clientResponse = make([]model.ClientResponse, 0)
	var clients = make([]model.Client, 0)
	cli := db.Get()
	err = cli.Table(clientTableName).Where("openid=?", openid).Find(&clients).Error
	if err != nil {
		return nil, err
	}
	for i := range clients {
		name := clients[i].ClientName + clients[i].StoveNum
		var clientRes = model.ClientResponse{Id: clients[i].Id, Name: name}
		clientResponse = append(clientResponse, clientRes)
	}
	return clientResponse, err
}

func (c ClientInterfaceImp) GetClientByClientId(id int) (model.Client, error) {
	//TODO implement me
	var err error
	var client = model.Client{}
	cli := db.Get()
	err = cli.Table(clientTableName).First(&client, id).Error
	return client, err
}

func (c ClientInterfaceImp) GetAllClientByPages(page int, pageSize int) (model.ClientPage, error) {
	var err error
	var clientPage = model.ClientPage{Page: page, PageSize: pageSize}
	offset := (page - 1) * pageSize
	cli := db.Get()
	err = cli.Table(clientTableName).Order("id asc").
		Limit(pageSize).Offset(offset).Scan(&clientPage.Data).
		Limit(-1).Offset(-1).Count(&clientPage.Total).
		Error
	return clientPage, err
}

func (c ClientInterfaceImp) GetAllClient() ([]model.Client, error) {
	var err error
	var client = make([]model.Client, 0)
	cli := db.Get()
	err = cli.Table(clientTableName).Find(&client).Error
	return client, err
}
