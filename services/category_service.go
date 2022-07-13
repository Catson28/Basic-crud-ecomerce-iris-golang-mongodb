package services

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/repo"
)

type CategoryService interface {
	List(m map[string]interface{}) (response datamodels.Response)
	Save(category datamodels.Category) (response datamodels.Response)
	GetByID(id string) (response datamodels.Response)
	GetByName(Usename string) (response datamodels.Response)
	DeleteByID(id string) (response datamodels.Response)
}

type categoryService struct {
	repo repo.CategoriesRepository
}

var categoryRepo = repo.NewCategoriesRepository()

func NewCategoryService() CategoryService {
	return &categoryService{
		repo: categoryRepo,
	}
}

func (g *categoryService) List(m map[string]interface{}) (response datamodels.Response) {
	categorys, _ := g.repo.List(nil)
	response.Code = 20000
	response.Msg = "success"
	response.Data = categorys
	return
}

func (g *categoryService) Save(category datamodels.Category) (response datamodels.Response) {
	err := g.repo.Save(category)
	if err != nil {
		response.Code = 30001
		response.Msg = fmt.Sprintf("保存数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"

	return
}

func (g *categoryService) GetByID(id string) (response datamodels.Response) {
	category, err := g.repo.GetByID(id)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = category
	return
}

func (g *categoryService) GetByName(Usename string) (response datamodels.Response) {
	category, err := g.repo.GetByName(Usename)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = category
	return
}

func (g *categoryService) DeleteByID(id string) (response datamodels.Response) {
	err := g.repo.DeleteByID(id)
	if err != nil {
		response.Code = 30003
		response.Msg = fmt.Sprintf("删除数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	return
}
