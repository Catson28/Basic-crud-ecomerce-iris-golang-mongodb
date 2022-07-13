package controllers

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/services"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

type TodoController struct {
	Ctx     iris.Context
	Service services.TodoService
}

func NewTodoController() *TodoController {
	return &TodoController{
		Service: services.NewTodoService(),
	}
}

func (g *TodoController) Get() (response datamodels.Response) {
	return g.Service.List(nil)
}

func (g *TodoController) GetBy(id string) (response datamodels.Response) {
	return g.Service.GetByID(id)
}

func (g *TodoController) DeleteBy(id string) (response datamodels.Response) {
	return g.Service.DeleteByID(id)
}

func (g *TodoController) PutBy(id string) (response datamodels.Response) {
	todo := datamodels.Todo{}
	err := g.Ctx.ReadJSON(&todo)
	if err != nil {
		response.Code = 40001
		response.Msg = fmt.Sprintf("参数解析失败：%v", err)
	}
	todo.ID = bson.ObjectIdHex(id)
	return g.Service.Save(todo)
}

//func (g *TodoController) Post() (response datamodels.Response) {
//todo := datamodels.Todo{}
//err := g.Ctx.ReadJSON(&todo)
//if err != nil {
//	response.Code = 40001
//	response.Msg = fmt.Sprintf("参数解析失败：%v", err)
//}
//return g.Service.Save(todo)
//}
func (g *TodoController) Post() (response datamodels.Response) {
	todo := datamodels.Todo{} //pegamos estrutura do modelo que vamos validar ou trabalhar
	err := g.Ctx.ReadJSON(&todo) // filtramos os dados requisitados e enviamos Para o Servico

	if err != nil { // varificamos se os dados filtrados estao corretos
		// Handle the error, below you will find the right way to do that...

		if errs, ok := err.(validator.ValidationErrors); ok {// verificamos novamente se realmente existe um erro
			//fmt.Println("errrooooooo")
			// Wrap the errors with JSON format, the underline library returns the errors as interface.
			validationErrors := wrapValidationErrors(errs)//Atribuimos o erro especifico em uma variavel

			// Fire an application/json+problem response and stop the handlers chain.

			//g.Ctx.StopWithProblem(iris.StatusBadRequest, )

			response.Code = 40001// apresentamos o valor do erro
			response.Msg = fmt.Sprintf("参数解析失败：%v", err) // apresentamos a mensagem de erro
			response.Data = iris.NewProblem().Title("Validation error").Detail("One or more fields failed to be validated").Type("/api/v1/todos/validation-errors").Key("errors", validationErrors)// Apresentamos o erro especifico

			return
		}

		// It's probably an internal JSON error, let's dont give more info here.
		fmt.Println("errrooooooo")
		g.Ctx.StopWithStatus(iris.StatusInternalServerError)// Caso Nem uma linha responda apresentamos  erro interno do servidor
		//return
	}

	return g.Service.Save(todo)
}
