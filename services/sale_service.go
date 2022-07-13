package services

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/repo"
)

type SaleService interface {
	List(m map[string]interface{}) (response datamodels.Response)
	Save(sale datamodels.Sale) (response datamodels.Response)
	GetByID(id string) (response datamodels.Response)
	GetByName(Usename string) (response datamodels.Response)
	DeleteByID(id string) (response datamodels.Response)
}

type saleService struct {
	repo repo.SalesRepository
}

var saleRepo = repo.NewSalesRepository()

func NewSaleService() SaleService {
	return &saleService{
		repo: saleRepo,
	}
}

func (g *saleService) List(m map[string]interface{}) (response datamodels.Response) {
	sales, _ := g.repo.List(nil)
	response.Code = 20000
	response.Msg = "success"
	response.Data = sales
	return
}

func (g *saleService) Save(sale datamodels.Sale) (response datamodels.Response) {
	err := g.repo.Save(sale)
	if err != nil {
		response.Code = 30001
		response.Msg = fmt.Sprintf("保存数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"

	return
}

func (g *saleService) GetByID(id string) (response datamodels.Response) {
	sale, err := g.repo.GetByID(id)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = sale
	return
}

func (g *saleService) GetByName(Usename string) (response datamodels.Response) {
	sale, err := g.repo.GetByName(Usename)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = sale
	return
}

func (g *saleService) DeleteByID(id string) (response datamodels.Response) {
	err := g.repo.DeleteByID(id)
	if err != nil {
		response.Code = 30003
		response.Msg = fmt.Sprintf("删除数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	return
}
