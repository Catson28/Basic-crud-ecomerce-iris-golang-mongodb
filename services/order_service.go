package services

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/repo"
)

type OrderService interface {
	List(m map[string]interface{}) (response datamodels.Response)
	Save(order datamodels.Order) (response datamodels.Response)
	GetByID(id string) (response datamodels.Response)
	GetByName(Usename string) (response datamodels.Response)
	DeleteByID(id string) (response datamodels.Response)
}

type orderService struct {
	repo repo.OrdersRepository
}

var orderRepo = repo.NewOrdersRepository()

func NewOrderService() OrderService {
	return &orderService{
		repo: orderRepo,
	}
}

func (g *orderService) List(m map[string]interface{}) (response datamodels.Response) {
	orders, _ := g.repo.List(nil)
	response.Code = 20000
	response.Msg = "success"
	response.Data = orders
	return
}

func (g *orderService) Save(order datamodels.Order) (response datamodels.Response) {
	err := g.repo.Save(order)
	if err != nil {
		response.Code = 30001
		response.Msg = fmt.Sprintf("保存数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"

	return
}

func (g *orderService) GetByID(id string) (response datamodels.Response) {
	order, err := g.repo.GetByID(id)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = order
	return
}

func (g *orderService) GetByName(Usename string) (response datamodels.Response) {
	order, err := g.repo.GetByName(Usename)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = order
	return
}

func (g *orderService) DeleteByID(id string) (response datamodels.Response) {
	err := g.repo.DeleteByID(id)
	if err != nil {
		response.Code = 30003
		response.Msg = fmt.Sprintf("删除数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	return
}
