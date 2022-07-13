package controllers

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/datamodels/request"
	"tentativa/services"

	"tentativa/util"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

type MoreController struct {
	Ctx     iris.Context
	Service services.MoreService
}

func NewMoreController() *MoreController {
	return &MoreController{
		Service: services.NewMoreService(),
	}
}

func (g *MoreController) Get() (response datamodels.Response) {
	return g.Service.List(nil)
}

func (g *MoreController) GetBy(id string) (response datamodels.Response) {
	return g.Service.GetByID(id)
}

func (g *MoreController) GetName(name string) (response datamodels.Response) {
	return g.Service.GetByName(name)
}

func (g *MoreController) DeleteBy(id string) (response datamodels.Response) {
	return g.Service.DeleteByID(id)
}

func (g *MoreController) PutBy(id string) (response datamodels.Response) {
	more := datamodels.More{}
	err := g.Ctx.ReadJSON(&more)
	if err != nil {
		response.Code = 40001
		response.Msg = fmt.Sprintf("参数解析失败：%v", err)
	}
	more.ID = bson.ObjectIdHex(id)
	return g.Service.Save(more)
}

//func (g *MoreController) Post() (response datamodels.Response) {
//more := datamodels.More{}
//err := g.Ctx.ReadJSON(&more)
//if err != nil {
//	response.Code = 40001
//	response.Msg = fmt.Sprintf("参数解析失败：%v", err)
//}
//return g.Service.Save(more)
//}
func (g *MoreController) Post() (response datamodels.Response) {
	moreRequest := request.MoreRequest{}

	more := datamodels.More{}
	//pssMore := more
	err := g.Ctx.ReadJSON(&moreRequest)

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
			response.Data = iris.NewProblem().Title("Validation error").Detail("One or more fields failed to be validated").Type("/api/v1/mores/validation-errors").Key("errors", validationErrors)

			return
		}

		// It's probably an internal JSON error, let's dont give more info here.
		fmt.Println("errrooooooo")
		g.Ctx.StopWithStatus(iris.StatusInternalServerError)
		//return
	}
	/*

	 */
	//
	hashedPassword, err := util.GeneratePassword(moreRequest.Password)
	if err != nil {
		//return datamodels.More{}, err
	}
	more.Password = hashedPassword
	more.Games = moreRequest.Games
	more.Name = moreRequest.Name

	return g.Service.Save(more)
}
