package repo

import (
	"tentativa/datamodels"
	"tentativa/datasource"

	"gopkg.in/mgo.v2/bson"
)

type OrdersRepository interface {
	List(query map[string]interface{}) (orders []datamodels.Order, err error)
	Save(order datamodels.Order) (err error)
	GetByID(id string) (order datamodels.Order, err error)
	GetByName(Usename string) (order datamodels.Order, err error)
	DeleteByID(id string) (err error)
}

type ordersRepository struct {
	collection string
}

func NewOrdersRepository() OrdersRepository {
	return &ordersRepository{
		collection: "orders",
	}
}

func (g *ordersRepository) List(query map[string]interface{}) (orders []datamodels.Order, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if err = col.Find(nil).All(&orders); err != nil {
		if err.Error() != datasource.GetErrNotFound().Error() {
			return
		}
	}
	return
}

func (g *ordersRepository) Save(order datamodels.Order) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if order.ID.Hex() == "" {
		order.ID = bson.NewObjectId()
		err = col.Insert(order)
	} else {
		err = col.Update(bson.M{"_id": order.ID}, order)
	}

	return
}

func (g *ordersRepository) GetByID(id string) (order datamodels.Order, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.FindId(bson.ObjectIdHex(id)).One(&order)
	return
}

func (g *ordersRepository) GetByName(Usename string) (order datamodels.Order, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)
	err = col.Find(bson.M{"name": "ASDFGH"}).One(&order) //Funciona
	err = col.Find(bson.M{"name": Usename}).One(&order)

	return
}

func (g *ordersRepository) DeleteByID(id string) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.RemoveId(bson.ObjectIdHex(id))
	return
}
