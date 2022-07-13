package services

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/repo"
)

type StoredService interface {
	List(m map[string]interface{}) (response datamodels.Response)
	Save(stored datamodels.Stored) (response datamodels.Response)
	GetByID(id string) (response datamodels.Response)
	GetByName(Usename string) (response datamodels.Response)
	DeleteByID(id string) (response datamodels.Response)
}

type storedService struct {
	repo repo.StoredsRepository
}

var storedRepo = repo.NewStoredsRepository()

func NewStoredService() StoredService {
	return &storedService{
		repo: storedRepo,
	}
}

func (g *storedService) List(m map[string]interface{}) (response datamodels.Response) {
	storeds, _ := g.repo.List(nil)
	response.Code = 20000
	response.Msg = "success"
	response.Data = storeds
	return
}

func (g *storedService) Save(stored datamodels.Stored) (response datamodels.Response) {
	err := g.repo.Save(stored)
	if err != nil {
		response.Code = 30001
		response.Msg = fmt.Sprintf("保存数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"

	return
}

func (g *storedService) GetByID(id string) (response datamodels.Response) {
	stored, err := g.repo.GetByID(id)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = stored
	return
}

func (g *storedService) GetByName(Usename string) (response datamodels.Response) {
	stored, err := g.repo.GetByName(Usename)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = stored
	return
}

func (g *storedService) DeleteByID(id string) (response datamodels.Response) {
	err := g.repo.DeleteByID(id)
	if err != nil {
		response.Code = 30003
		response.Msg = fmt.Sprintf("删除数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	return
}
