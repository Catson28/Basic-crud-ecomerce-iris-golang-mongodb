package repo

import (
	"tentativa/datamodels"
	"tentativa/datasource"

	"gopkg.in/mgo.v2/bson"
)

type ClientsRepository interface {
	List(query map[string]interface{}) (clients []datamodels.Client, err error)
	Save(client datamodels.Client) (err error)
	GetByID(id string) (client datamodels.Client, err error)
	GetByName(Usename string) (client datamodels.Client, err error)
	DeleteByID(id string) (err error)
}

type clientsRepository struct {
	collection string
}

func NewClientsRepository() ClientsRepository {
	return &clientsRepository{
		collection: "clients",
	}
}

func (g *clientsRepository) List(query map[string]interface{}) (clients []datamodels.Client, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if err = col.Find(nil).All(&clients); err != nil {
		if err.Error() != datasource.GetErrNotFound().Error() {
			return
		}
	}
	return
}

func (g *clientsRepository) Save(client datamodels.Client) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if client.ID.Hex() == "" {
		client.ID = bson.NewObjectId()
		err = col.Insert(client)
	} else {
		err = col.Update(bson.M{"_id": client.ID}, client)
	}

	return
}

func (g *clientsRepository) GetByID(id string) (client datamodels.Client, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.FindId(bson.ObjectIdHex(id)).One(&client)
	return
}

func (g *clientsRepository) GetByName(Usename string) (client datamodels.Client, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)
	err = col.Find(bson.M{"name": "ASDFGH"}).One(&client) //Funciona
	err = col.Find(bson.M{"name": Usename}).One(&client)

	return
}

func (g *clientsRepository) DeleteByID(id string) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.RemoveId(bson.ObjectIdHex(id))
	return
}
