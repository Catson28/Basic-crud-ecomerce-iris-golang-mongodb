package services

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/repo"
)

type TodoService interface {
	List(m map[string]interface{}) (response datamodels.Response)
	Save(todo datamodels.Todo) (response datamodels.Response)
	GetByID(id string) (response datamodels.Response)
	DeleteByID(id string) (response datamodels.Response)
}

type todoService struct {
	repo repo.TodosRepository
}

var todoRepo = repo.NewTodosRepository()

func NewTodoService() TodoService {
	return &todoService{
		repo: todoRepo,
	}
}

func (g *todoService) List(m map[string]interface{}) (response datamodels.Response) {
	todos, _ := g.repo.List(nil)
	response.Code = 20000
	response.Msg = "success"
	response.Data = todos
	return
}

func (g *todoService) Save(todo datamodels.Todo) (response datamodels.Response) {
	err := g.repo.Save(todo)
	if err != nil {
		response.Code = 30001
		response.Msg = fmt.Sprintf("保存数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"

	return
}

func (g *todoService) GetByID(id string) (response datamodels.Response) {
	todo, err := g.repo.GetByID(id)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = todo
	return
}

func (g *todoService) DeleteByID(id string) (response datamodels.Response) {
	err := g.repo.DeleteByID(id)
	if err != nil {
		response.Code = 30003
		response.Msg = fmt.Sprintf("删除数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	return
}
