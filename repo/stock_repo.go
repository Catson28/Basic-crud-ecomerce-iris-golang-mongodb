package repo

import (
	"tentativa/datamodels"
	"tentativa/datasource"

	"gopkg.in/mgo.v2/bson"
)

type StocksRepository interface {
	List(query map[string]interface{}) (stocks []datamodels.Stock, err error)
	Save(stock datamodels.Stock) (err error)
	GetByID(id string) (stock datamodels.Stock, err error)
	GetByName(Usename string) (stock datamodels.Stock, err error)
	DeleteByID(id string) (err error)
}

type stocksRepository struct {
	collection string
}

func NewStocksRepository() StocksRepository {
	return &stocksRepository{
		collection: "stocks",
	}
}

func (g *stocksRepository) List(query map[string]interface{}) (stocks []datamodels.Stock, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if err = col.Find(nil).All(&stocks); err != nil {
		if err.Error() != datasource.GetErrNotFound().Error() {
			return
		}
	}
	return
}

func (g *stocksRepository) Save(stock datamodels.Stock) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if stock.ID.Hex() == "" {
		stock.ID = bson.NewObjectId()
		err = col.Insert(stock)
	} else {
		err = col.Update(bson.M{"_id": stock.ID}, stock)
	}

	return
}

func (g *stocksRepository) GetByID(id string) (stock datamodels.Stock, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.FindId(bson.ObjectIdHex(id)).One(&stock)
	return
}

func (g *stocksRepository) GetByName(Usename string) (stock datamodels.Stock, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)
	err = col.Find(bson.M{"name": "ASDFGH"}).One(&stock) //Funciona
	err = col.Find(bson.M{"name": Usename}).One(&stock)

	return
}

func (g *stocksRepository) DeleteByID(id string) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.RemoveId(bson.ObjectIdHex(id))
	return
}
