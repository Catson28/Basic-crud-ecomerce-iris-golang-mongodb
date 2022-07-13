package controllers

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/services"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

type StoredController struct {
	Ctx     iris.Context
	Service services.StoredService
}

func NewStoredController() *StoredController {
	return &StoredController{
		Service: services.NewStoredService(),
	}
}

func (g *StoredController) Get() (response datamodels.Response) {
	return g.Service.List(nil)
}

func (g *StoredController) GetBy(id string) (response datamodels.Response) {
	return g.Service.GetByID(id)
}

func (g *StoredController) GetName(name string) (response datamodels.Response) {
	return g.Service.GetByName(name)
}

func (g *StoredController) DeleteBy(id string) (response datamodels.Response) {
	return g.Service.DeleteByID(id)
}

func (g *StoredController) PutBy(id string) (response datamodels.Response) {
	stored := datamodels.Stored{}
	err := g.Ctx.ReadJSON(&stored)
	if err != nil {
		response.Code = 40001
		response.Msg = fmt.Sprintf("参数解析失败：%v", err)
	}
	stored.ID = bson.ObjectIdHex(id)
	return g.Service.Save(stored)
}

func (g *StoredController) Post() (response datamodels.Response) {
	stored := datamodels.Stored{}
	err := g.Ctx.ReadJSON(&stored)

	if err != nil {
		// Handle the error, below you will find the right way to do that...

		if errs, ok := err.(validator.ValidationErrors); ok {
			//fmt.Println("errrooooooo")
			// Wrap the errors with JSON format, the underline library returns the errors as interface.
			validationErrors := wrapValidationErrors(errs)

			// Fire an application/json+problem response and stop the handlers chain.

			//g.Ctx.StopWithProblem(iris.StatusBadRequest, )

			response.Code = 40001
			response.Msg = fmt.Sprintf("参数解析失败：%v", err)
			response.Data = iris.NewProblem().Title("Validation error").Detail("One or more fields failed to be validated").Type("/api/v1/storeds/validation-errors").Key("errors", validationErrors)

			return
		}

		// It's probably an internal JSON error, let's dont give more info here.
		fmt.Println("errrooooooo")
		g.Ctx.StopWithStatus(iris.StatusInternalServerError)
		//return
	}

	return g.Service.Save(stored)
}
