package services

import (
	"fmt"
	"tentativa/datamodels"
	"tentativa/repo"
)

type GameService interface {
	List(m map[string]interface{}) (response datamodels.Response)
	Save(game datamodels.Game) (response datamodels.Response)
	GetByID(id string) (response datamodels.Response)
	GetByName(Usename string) (response datamodels.Response)
	DeleteByID(id string) (response datamodels.Response)
}

type gameService struct {
	repo repo.GamesRepository
}

var gameRepo = repo.NewGamesRepository()

func NewGameService() GameService {
	return &gameService{
		repo: gameRepo,
	}
}

func (g *gameService) List(m map[string]interface{}) (response datamodels.Response) {
	games, _ := g.repo.List(nil)
	response.Code = 20000
	response.Msg = "success"
	response.Data = games
	return
}

func (g *gameService) Save(game datamodels.Game) (response datamodels.Response) {
	err := g.repo.Save(game)
	if err != nil {
		response.Code = 30001
		response.Msg = fmt.Sprintf("保存数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"

	return
}

func (g *gameService) GetByID(id string) (response datamodels.Response) {
	game, err := g.repo.GetByID(id)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = game
	return
}

func (g *gameService) GetByName(Usename string) (response datamodels.Response) {
	game, err := g.repo.GetByName(Usename)
	if err != nil {
		response.Code = 30002
		response.Msg = fmt.Sprintf("查询数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	response.Data = game
	return
}

func (g *gameService) DeleteByID(id string) (response datamodels.Response) {
	err := g.repo.DeleteByID(id)
	if err != nil {
		response.Code = 30003
		response.Msg = fmt.Sprintf("删除数据失败：%v", err)
	}
	response.Code = 20000
	response.Msg = "success"
	return
}
