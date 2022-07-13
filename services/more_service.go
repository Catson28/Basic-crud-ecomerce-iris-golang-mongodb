package services

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/repo"
)

type MoreService interface {
	List(m map[string]interface{}) (response datamodels.Response)
	Save(more datamodels.More) (response datamodels.Response)
	GetByID(id string) (response datamodels.Response)
	GetByName(Usename string) (response datamodels.Response)
	DeleteByID(id string) (response datamodels.Response)
}

type moreService struct {
	repo repo.MoresRepository
}

var moreRepo = repo.NewMoresRepository()

func NewMoreService() MoreService {
	return &moreService{
		repo: moreRepo,
	}
}

func (g *moreService) List(m map[string]interface{}) (response datamodels.Response) {
	mores, _ := g.repo.List(nil)
	response.Code = 20000
	response.Msg = "success"
	response.Data = mores
	return
}

func (g *moreService) Save(more datamodels.More) (response datamodels.Response) {
	err := g.repo.Save(more)
	if err != nil {
		response.Code = 30001
		response.Msg = fmt.Sprintf("保存数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"

	return
}

func (g *moreService) GetByID(id string) (response datamodels.Response) {
	more, err := g.repo.GetByID(id)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = more
	return
}

func (g *moreService) GetByName(Usename string) (response datamodels.Response) {
	more, err := g.repo.GetByName(Usename)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = more
	return
}

func (g *moreService) DeleteByID(id string) (response datamodels.Response) {
	err := g.repo.DeleteByID(id)
	if err != nil {
		response.Code = 30003
		response.Msg = fmt.Sprintf("删除数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	return
}
