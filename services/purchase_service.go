package services

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/repo"
)

type PurchaseService interface {
	List(m map[string]interface{}) (response datamodels.Response)
	Save(purchase datamodels.Purchase) (response datamodels.Response)
	GetByID(id string) (response datamodels.Response)
	GetByName(Usename string) (response datamodels.Response)
	DeleteByID(id string) (response datamodels.Response)
}

type purchaseService struct {
	repo repo.PurchasesRepository
}

var purchaseRepo = repo.NewPurchasesRepository()

func NewPurchaseService() PurchaseService {
	return &purchaseService{
		repo: purchaseRepo,
	}
}

func (g *purchaseService) List(m map[string]interface{}) (response datamodels.Response) {
	purchases, _ := g.repo.List(nil)
	response.Code = 20000
	response.Msg = "success"
	response.Data = purchases
	return
}

func (g *purchaseService) Save(purchase datamodels.Purchase) (response datamodels.Response) {
	err := g.repo.Save(purchase)
	if err != nil {
		response.Code = 30001
		response.Msg = fmt.Sprintf("保存数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"

	return
}

func (g *purchaseService) GetByID(id string) (response datamodels.Response) {
	purchase, err := g.repo.GetByID(id)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = purchase
	return
}

func (g *purchaseService) GetByName(Usename string) (response datamodels.Response) {
	purchase, err := g.repo.GetByName(Usename)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = purchase
	return
}

func (g *purchaseService) DeleteByID(id string) (response datamodels.Response) {
	err := g.repo.DeleteByID(id)
	if err != nil {
		response.Code = 30003
		response.Msg = fmt.Sprintf("删除数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	return
}
