package repo

import (
	"tentativa/datamodels"
	"tentativa/datasource"

	"gopkg.in/mgo.v2/bson"
)

type MoresRepository interface {
	List(query map[string]interface{}) (mores []datamodels.More, err error)
	Save(more datamodels.More) (err error)
	GetByID(id string) (more datamodels.More, err error)
	GetByName(Usename string) (more datamodels.More, err error)
	DeleteByID(id string) (err error)
}

type moresRepository struct {
	collection string
}

func NewMoresRepository() MoresRepository {
	return &moresRepository{
		collection: "mores",
	}
}

func (m *moresRepository) List(query map[string]interface{}) (mores []datamodels.More, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(m.collection)

	if err = col.Find(nil).All(&mores); err != nil {
		if err.Error() != datasource.GetErrNotFound().Error() {
			return
		}
	}
	return
}

func (m *moresRepository) Save(more datamodels.More) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(m.collection)

	if more.ID.Hex() == "" {
		more.ID = bson.NewObjectId()
		err = col.Insert(more)
	} else {
		err = col.Update(bson.M{"_id": more.ID}, more)
	}

	return
}

func (m *moresRepository) GetByID(id string) (more datamodels.More, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(m.collection)

	err = col.FindId(bson.ObjectIdHex(id)).One(&more)
	return
}

func (m *moresRepository) GetByName(Usename string) (more datamodels.More, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(m.collection)
	err = col.Find(bson.M{"name": "ASDFGH"}).One(&more) //Funciona
	err = col.Find(bson.M{"name": Usename}).One(&more)

	return
}

func (m *moresRepository) DeleteByID(id string) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(m.collection)

	err = col.RemoveId(bson.ObjectIdHex(id))
	return
}
