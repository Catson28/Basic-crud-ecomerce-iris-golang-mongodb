package services

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/repo"
)

type ClientService interface {
	List(m map[string]interface{}) (response datamodels.Response)
	Save(client datamodels.Client) (response datamodels.Response)
	GetByID(id string) (response datamodels.Response)
	GetByName(Usename string) (response datamodels.Response)
	DeleteByID(id string) (response datamodels.Response)
}

type clientService struct {
	repo repo.ClientsRepository
}

var clientRepo = repo.NewClientsRepository()

func NewClientService() ClientService {
	return &clientService{
		repo: clientRepo,
	}
}

func (g *clientService) List(m map[string]interface{}) (response datamodels.Response) {
	clients, _ := g.repo.List(nil)
	response.Code = 20000
	response.Msg = "success"
	response.Data = clients
	return
}

func (g *clientService) Save(client datamodels.Client) (response datamodels.Response) {
	err := g.repo.Save(client)
	if err != nil {
		response.Code = 30001
		response.Msg = fmt.Sprintf("保存数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"

	return
}

func (g *clientService) GetByID(id string) (response datamodels.Response) {
	client, err := g.repo.GetByID(id)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = client
	return
}

func (g *clientService) GetByName(Usename string) (response datamodels.Response) {
	client, err := g.repo.GetByName(Usename)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = client
	return
}

func (g *clientService) DeleteByID(id string) (response datamodels.Response) {
	err := g.repo.DeleteByID(id)
	if err != nil {
		response.Code = 30003
		response.Msg = fmt.Sprintf("删除数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	return
}
