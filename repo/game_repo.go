package repo

import (
	"tentativa/datamodels"
	"tentativa/datasource"

	"gopkg.in/mgo.v2/bson"
)

type GamesRepository interface {
	List(query map[string]interface{}) (games []datamodels.Game, err error)
	Save(game datamodels.Game) (err error)
	GetByID(id string) (game datamodels.Game, err error)
	GetByName(Usename string) (game datamodels.Game, err error)
	DeleteByID(id string) (err error)
}

type gamesRepository struct {
	collection string
}

func NewGamesRepository() GamesRepository {
	return &gamesRepository{
		collection: "games",
	}
}

func (g *gamesRepository) List(query map[string]interface{}) (games []datamodels.Game, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if err = col.Find(nil).All(&games); err != nil {
		if err.Error() != datasource.GetErrNotFound().Error() {
			return
		}
	}
	return
}

func (g *gamesRepository) Save(game datamodels.Game) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if game.ID.Hex() == "" {
		game.ID = bson.NewObjectId()
		err = col.Insert(game)
	} else {
		err = col.Update(bson.M{"_id": game.ID}, game)
	}

	return
}

func (g *gamesRepository) GetByID(id string) (game datamodels.Game, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.FindId(bson.ObjectIdHex(id)).One(&game)
	return
}

func (g *gamesRepository) GetByName(Usename string) (game datamodels.Game, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)
	err = col.Find(bson.M{"name": "ASDFGH"}).One(&game) //Funciona
	err = col.Find(bson.M{"name": Usename}).One(&game)

	return
}

func (g *gamesRepository) DeleteByID(id string) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.RemoveId(bson.ObjectIdHex(id))
	return
}
