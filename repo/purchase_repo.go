package repo

import (
	"tentativa/datamodels"
	"tentativa/datasource"

	"gopkg.in/mgo.v2/bson"
)

type PurchasesRepository interface {
	List(query map[string]interface{}) (purchases []datamodels.Purchase, err error)
	Save(purchase datamodels.Purchase) (err error)
	GetByID(id string) (purchase datamodels.Purchase, err error)
	GetByName(Usename string) (purchase datamodels.Purchase, err error)
	DeleteByID(id string) (err error)
}

type purchasesRepository struct {
	collection string
}

func NewPurchasesRepository() PurchasesRepository {
	return &purchasesRepository{
		collection: "purchases",
	}
}

func (g *purchasesRepository) List(query map[string]interface{}) (purchases []datamodels.Purchase, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if err = col.Find(nil).All(&purchases); err != nil {
		if err.Error() != datasource.GetErrNotFound().Error() {
			return
		}
	}
	return
}

func (g *purchasesRepository) Save(purchase datamodels.Purchase) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if purchase.ID.Hex() == "" {
		purchase.ID = bson.NewObjectId()
		err = col.Insert(purchase)
	} else {
		err = col.Update(bson.M{"_id": purchase.ID}, purchase)
	}

	return
}

func (g *purchasesRepository) GetByID(id string) (purchase datamodels.Purchase, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.FindId(bson.ObjectIdHex(id)).One(&purchase)
	return
}

func (g *purchasesRepository) GetByName(Usename string) (purchase datamodels.Purchase, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)
	err = col.Find(bson.M{"name": "ASDFGH"}).One(&purchase) //Funciona
	err = col.Find(bson.M{"name": Usename}).One(&purchase)

	return
}

func (g *purchasesRepository) DeleteByID(id string) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.RemoveId(bson.ObjectIdHex(id))
	return
}
