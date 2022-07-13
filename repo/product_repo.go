package repo

import (
	"tentativa/datamodels"
	"tentativa/datasource"

	"gopkg.in/mgo.v2/bson"
)

type ProductsRepository interface {
	List(query map[string]interface{}) (products []datamodels.Product, err error)
	Save(product datamodels.Product) (err error)
	GetByID(id string) (product datamodels.Product, err error)
	GetByName(Usename string) (product datamodels.Product, err error)
	DeleteByID(id string) (err error)
}

type productsRepository struct {
	collection string
}

func NewProductsRepository() ProductsRepository {
	return &productsRepository{
		collection: "products",
	}
}

func (g *productsRepository) List(query map[string]interface{}) (products []datamodels.Product, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if err = col.Find(nil).All(&products); err != nil {
		if err.Error() != datasource.GetErrNotFound().Error() {
			return
		}
	}
	return
}

func (g *productsRepository) Save(product datamodels.Product) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if product.ID.Hex() == "" {
		product.ID = bson.NewObjectId()
		err = col.Insert(product)
	} else {
		err = col.Update(bson.M{"_id": product.ID}, product)
	}

	return
}

func (g *productsRepository) GetByID(id string) (product datamodels.Product, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.FindId(bson.ObjectIdHex(id)).One(&product)
	return
}

func (g *productsRepository) GetByName(Usename string) (product datamodels.Product, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)
	err = col.Find(bson.M{"name": "ASDFGH"}).One(&product) //Funciona
	err = col.Find(bson.M{"name": Usename}).One(&product)

	return
}

func (g *productsRepository) DeleteByID(id string) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.RemoveId(bson.ObjectIdHex(id))
	return
}
