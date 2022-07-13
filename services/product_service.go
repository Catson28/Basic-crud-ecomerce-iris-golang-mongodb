package services

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/repo"
)

type ProductService interface {
	List(m map[string]interface{}) (response datamodels.Response)
	Save(product datamodels.Product) (response datamodels.Response)
	GetByID(id string) (response datamodels.Response)
	GetByName(Usename string) (response datamodels.Response)
	DeleteByID(id string) (response datamodels.Response)
}

type productService struct {
	repo repo.ProductsRepository
}

var productRepo = repo.NewProductsRepository()

func NewProductService() ProductService {
	return &productService{
		repo: productRepo,
	}
}

func (g *productService) List(m map[string]interface{}) (response datamodels.Response) {
	products, _ := g.repo.List(nil)
	response.Code = 20000
	response.Msg = "success"
	response.Data = products
	return
}

func (g *productService) Save(product datamodels.Product) (response datamodels.Response) {
	err := g.repo.Save(product)
	if err != nil {
		response.Code = 30001
		response.Msg = fmt.Sprintf("保存数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"

	return
}

func (g *productService) GetByID(id string) (response datamodels.Response) {
	product, err := g.repo.GetByID(id)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = product
	return
}

func (g *productService) GetByName(Usename string) (response datamodels.Response) {
	product, err := g.repo.GetByName(Usename)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = product
	return
}

func (g *productService) DeleteByID(id string) (response datamodels.Response) {
	err := g.repo.DeleteByID(id)
	if err != nil {
		response.Code = 30003
		response.Msg = fmt.Sprintf("删除数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	return
}
