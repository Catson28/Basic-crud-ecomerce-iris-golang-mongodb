package controllers

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/services"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

type PurchaseController struct {
	Ctx     iris.Context
	Service services.PurchaseService
}

func NewPurchaseController() *PurchaseController {
	return &PurchaseController{
		Service: services.NewPurchaseService(),
	}
}

func (g *PurchaseController) Get() (response datamodels.Response) {
	return g.Service.List(nil)
}

func (g *PurchaseController) GetBy(id string) (response datamodels.Response) {
	return g.Service.GetByID(id)
}

func (g *PurchaseController) GetName(name string) (response datamodels.Response) {
	return g.Service.GetByName(name)
}

func (g *PurchaseController) DeleteBy(id string) (response datamodels.Response) {
	return g.Service.DeleteByID(id)
}

func (g *PurchaseController) PutBy(id string) (response datamodels.Response) {
	purchase := datamodels.Purchase{}
	err := g.Ctx.ReadJSON(&purchase)
	if err != nil {
		response.Code = 40001
		response.Msg = fmt.Sprintf("参数解析失败：%v", err)
	}
	purchase.ID = bson.ObjectIdHex(id)
	return g.Service.Save(purchase)
}

func (g *PurchaseController) Post() (response datamodels.Response) {
	purchase := datamodels.Purchase{}
	err := g.Ctx.ReadJSON(&purchase)

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
			response.Data = iris.NewProblem().Title("Validation error").Detail("One or more fields failed to be validated").Type("/api/v1/purchases/validation-errors").Key("errors", validationErrors)

			return
		}

		// It's probably an internal JSON error, let's dont give more info here.
		fmt.Println("errrooooooo")
		g.Ctx.StopWithStatus(iris.StatusInternalServerError)
		//return
	}

	return g.Service.Save(purchase)
}
