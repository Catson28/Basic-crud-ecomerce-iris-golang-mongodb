package controllers

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/services"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

type ProductController struct {
	Ctx     iris.Context
	Service services.ProductService
}

func NewProductController() *ProductController {
	return &ProductController{
		Service: services.NewProductService(),
	}
}

func (g *ProductController) Get() (response datamodels.Response) {
	return g.Service.List(nil)
}

func (g *ProductController) GetBy(id string) (response datamodels.Response) {
	return g.Service.GetByID(id)
}

func (g *ProductController) GetName(name string) (response datamodels.Response) {
	return g.Service.GetByName(name)
}

func (g *ProductController) DeleteBy(id string) (response datamodels.Response) {
	return g.Service.DeleteByID(id)
}

func (g *ProductController) PutBy(id string) (response datamodels.Response) {
	product := datamodels.Product{}
	err := g.Ctx.ReadJSON(&product)
	if err != nil {
		response.Code = 40001
		response.Msg = fmt.Sprintf("参数解析失败：%v", err)
	}
	product.ID = bson.ObjectIdHex(id)
	return g.Service.Save(product)
}

func (g *ProductController) Post() (response datamodels.Response) {
	product := datamodels.Product{}
	err := g.Ctx.ReadJSON(&product)

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
			response.Data = iris.NewProblem().Title("Validation error").Detail("One or more fields failed to be validated").Type("/api/v1/products/validation-errors").Key("errors", validationErrors)

			return
		}

		// It's probably an internal JSON error, let's dont give more info here.
		fmt.Println("errrooooooo")
		g.Ctx.StopWithStatus(iris.StatusInternalServerError)
		//return
	}

	return g.Service.Save(product)
}
