package repo

import (
	"tentativa/datamodels"
	"tentativa/datasource"

	"gopkg.in/mgo.v2/bson"
)

type AddressesRepository interface {
	List(query map[string]interface{}) (addresses []datamodels.Address, err error)
	Save(address datamodels.Address) (err error)
	GetByID(id string) (address datamodels.Address, err error)
	GetByName(Usename string) (address datamodels.Address, err error)
	DeleteByID(id string) (err error)
}

type addressesRepository struct {
	collection string
}

func NewAddressesRepository() AddressesRepository {
	return &addressesRepository{
		collection: "addresses",
	}
}

func (g *addressesRepository) List(query map[string]interface{}) (addresses []datamodels.Address, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if err = col.Find(nil).All(&addresses); err != nil {
		if err.Error() != datasource.GetErrNotFound().Error() {
			return
		}
	}
	return
}

func (g *addressesRepository) Save(address datamodels.Address) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if address.ID.Hex() == "" {
		address.ID = bson.NewObjectId()
		err = col.Insert(address)
	} else {
		err = col.Update(bson.M{"_id": address.ID}, address)
	}

	return
}

func (g *addressesRepository) GetByID(id string) (address datamodels.Address, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.FindId(bson.ObjectIdHex(id)).One(&address)
	return
}

func (g *addressesRepository) GetByName(Usename string) (address datamodels.Address, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)
	err = col.Find(bson.M{"name": "ASDFGH"}).One(&address) //Funciona
	err = col.Find(bson.M{"name": Usename}).One(&address)

	return
}

func (g *addressesRepository) DeleteByID(id string) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.RemoveId(bson.ObjectIdHex(id))
	return
}
