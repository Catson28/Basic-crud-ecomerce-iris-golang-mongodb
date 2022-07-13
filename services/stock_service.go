package services

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/repo"
)

type StockService interface {
	List(m map[string]interface{}) (response datamodels.Response)
	Save(stock datamodels.Stock) (response datamodels.Response)
	GetByID(id string) (response datamodels.Response)
	GetByName(Usename string) (response datamodels.Response)
	DeleteByID(id string) (response datamodels.Response)
}

type stockService struct {
	repo repo.StocksRepository
}

var stockRepo = repo.NewStocksRepository()

func NewStockService() StockService {
	return &stockService{
		repo: stockRepo,
	}
}

func (g *stockService) List(m map[string]interface{}) (response datamodels.Response) {
	stocks, _ := g.repo.List(nil)
	response.Code = 20000
	response.Msg = "success"
	response.Data = stocks
	return
}

func (g *stockService) Save(stock datamodels.Stock) (response datamodels.Response) {
	err := g.repo.Save(stock)
	if err != nil {
		response.Code = 30001
		response.Msg = fmt.Sprintf("保存数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"

	return
}

func (g *stockService) GetByID(id string) (response datamodels.Response) {
	stock, err := g.repo.GetByID(id)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = stock
	return
}

func (g *stockService) GetByName(Usename string) (response datamodels.Response) {
	stock, err := g.repo.GetByName(Usename)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = stock
	return
}

func (g *stockService) DeleteByID(id string) (response datamodels.Response) {
	err := g.repo.DeleteByID(id)
	if err != nil {
		response.Code = 30003
		response.Msg = fmt.Sprintf("删除数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	return
}
