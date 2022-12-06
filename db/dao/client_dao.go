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

func (c ClientInterfaceImp) ClearClient(id int32) error {
	//TODO implement me
	var err error
	cli := db.Get()
	err = cli.Table(clientTableName).Delete(&model.Client{Id: id}).Error
	return err
}

func (c ClientInterfaceImp) GetClientByUserid(userId int32) ([]model.ClientResponse, error) {
	//TODO implement me
	var err error
	var clientResponse = make([]model.ClientResponse, 0)
	var clients = make([]model.Client, 0)
	cli := db.Get()
	err = cli.Table(clientTableName).Where("user_id=?", userId).Find(&clients).Error
	for i := range clients {
		name := clients[i].ClientName + clients[i].StoveNum
		var clientRes = model.ClientResponse{Id: clients[i].Id, Name: name}
		clientResponse = append(clientResponse, clientRes)
	}
	return clientResponse, err
}

func (c ClientInterfaceImp) GetClientById(id int32) (model.Client, error) {
	//TODO implement me
	var err error
	var client = model.Client{}
	cli := db.Get()
	err = cli.Table(clientTableName).Where("id=?", id).Find(&client).Error
	return client, err
}
