package repo

import (
	"tentativa/datamodels"
	"tentativa/datasource"

	"gopkg.in/mgo.v2/bson"
)

type SalesRepository interface {
	List(query map[string]interface{}) (sales []datamodels.Sale, err error)
	Save(sale datamodels.Sale) (err error)
	GetByID(id string) (sale datamodels.Sale, err error)
	GetByName(Usename string) (sale datamodels.Sale, err error)
	DeleteByID(id string) (err error)
}

type salesRepository struct {
	collection string
}

func NewSalesRepository() SalesRepository {
	return &salesRepository{
		collection: "sales",
	}
}

func (g *salesRepository) List(query map[string]interface{}) (sales []datamodels.Sale, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if err = col.Find(nil).All(&sales); err != nil {
		if err.Error() != datasource.GetErrNotFound().Error() {
			return
		}
	}
	return
}

func (g *salesRepository) Save(sale datamodels.Sale) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if sale.ID.Hex() == "" {
		sale.ID = bson.NewObjectId()
		err = col.Insert(sale)
	} else {
		err = col.Update(bson.M{"_id": sale.ID}, sale)
	}

	return
}

func (g *salesRepository) GetByID(id string) (sale datamodels.Sale, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.FindId(bson.ObjectIdHex(id)).One(&sale)
	return
}

func (g *salesRepository) GetByName(Usename string) (sale datamodels.Sale, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)
	err = col.Find(bson.M{"name": "ASDFGH"}).One(&sale) //Funciona
	err = col.Find(bson.M{"name": Usename}).One(&sale)

	return
}

func (g *salesRepository) DeleteByID(id string) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.RemoveId(bson.ObjectIdHex(id))
	return
}
