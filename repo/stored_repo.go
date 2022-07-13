package repo

import (
	"tentativa/datamodels"
	"tentativa/datasource"

	"gopkg.in/mgo.v2/bson"
)

type StoredsRepository interface {
	List(query map[string]interface{}) (storeds []datamodels.Stored, err error)
	Save(stored datamodels.Stored) (err error)
	GetByID(id string) (stored datamodels.Stored, err error)
	GetByName(Usename string) (stored datamodels.Stored, err error)
	DeleteByID(id string) (err error)
}

type storedsRepository struct {
	collection string
}

func NewStoredsRepository() StoredsRepository {
	return &storedsRepository{
		collection: "storeds",
	}
}

func (g *storedsRepository) List(query map[string]interface{}) (storeds []datamodels.Stored, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if err = col.Find(nil).All(&storeds); err != nil {
		if err.Error() != datasource.GetErrNotFound().Error() {
			return
		}
	}
	return
}

func (g *storedsRepository) Save(stored datamodels.Stored) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if stored.ID.Hex() == "" {
		stored.ID = bson.NewObjectId()
		err = col.Insert(stored)
	} else {
		err = col.Update(bson.M{"_id": stored.ID}, stored)
	}

	return
}

func (g *storedsRepository) GetByID(id string) (stored datamodels.Stored, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.FindId(bson.ObjectIdHex(id)).One(&stored)
	return
}

func (g *storedsRepository) GetByName(Usename string) (stored datamodels.Stored, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)
	err = col.Find(bson.M{"name": "ASDFGH"}).One(&stored) //Funciona
	err = col.Find(bson.M{"name": Usename}).One(&stored)

	return
}

func (g *storedsRepository) DeleteByID(id string) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.RemoveId(bson.ObjectIdHex(id))
	return
}
