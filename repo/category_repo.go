package repo

import (
	"tentativa/datamodels"
	"tentativa/datasource"

	"gopkg.in/mgo.v2/bson"
)

type CategoriesRepository interface {
	List(query map[string]interface{}) (categories []datamodels.Category, err error)
	Save(category datamodels.Category) (err error)
	GetByID(id string) (category datamodels.Category, err error)
	GetByName(Usename string) (category datamodels.Category, err error)
	DeleteByID(id string) (err error)
}

type categoriesRepository struct {
	collection string
}

func NewCategoriesRepository() CategoriesRepository {
	return &categoriesRepository{
		collection: "categories",
	}
}

func (g *categoriesRepository) List(query map[string]interface{}) (categories []datamodels.Category, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if err = col.Find(nil).All(&categories); err != nil {
		if err.Error() != datasource.GetErrNotFound().Error() {
			return
		}
	}
	return
}

func (g *categoriesRepository) Save(category datamodels.Category) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if category.ID.Hex() == "" {
		category.ID = bson.NewObjectId()
		err = col.Insert(category)
	} else {
		err = col.Update(bson.M{"_id": category.ID}, category)
	}

	return
}

func (g *categoriesRepository) GetByID(id string) (category datamodels.Category, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.FindId(bson.ObjectIdHex(id)).One(&category)
	return
}

func (g *categoriesRepository) GetByName(Usename string) (category datamodels.Category, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)
	err = col.Find(bson.M{"name": "ASDFGH"}).One(&category) //Funciona
	err = col.Find(bson.M{"name": Usename}).One(&category)

	return
}

func (g *categoriesRepository) DeleteByID(id string) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.RemoveId(bson.ObjectIdHex(id))
	return
}
