package repo

import (
	"tentativa/datamodels"
	"tentativa/datasource"

	"gopkg.in/mgo.v2/bson"
)

type TodosRepository interface {
	List(query map[string]interface{}) (todos []datamodels.Todo, err error)
	Save(todo datamodels.Todo) (err error)
	GetByID(id string) (todo datamodels.Todo, err error)
	DeleteByID(id string) (err error)
}

type todosRepository struct {
	collection string
}

func NewTodosRepository() TodosRepository {
	return &todosRepository{
		collection: "project_todos",
	}
}

func (g *todosRepository) List(query map[string]interface{}) (todos []datamodels.Todo, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if err = col.Find(nil).All(&todos); err != nil {
		if err.Error() != datasource.GetErrNotFound().Error() {
			return
		}
	}
	return
}

func (g *todosRepository) Save(todo datamodels.Todo) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	if todo.ID.Hex() == "" {
		todo.ID = bson.NewObjectId()
		err = col.Insert(todo)
	} else {
		err = col.Update(bson.M{"_id": todo.ID}, todo)
	}

	return
}

func (g *todosRepository) GetByID(id string) (todo datamodels.Todo, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.FindId(bson.ObjectIdHex(id)).One(&todo)
	return
}

func (g *todosRepository) DeleteByID(id string) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(g.collection)

	err = col.RemoveId(bson.ObjectIdHex(id))
	return
}
