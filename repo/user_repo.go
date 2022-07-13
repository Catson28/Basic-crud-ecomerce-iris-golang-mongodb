package repo

import (
	"tentativa/datamodels"
	"tentativa/datasource"

	"gopkg.in/mgo.v2/bson"
)

type UsersRepository interface {
	List(query map[string]interface{}) (users []datamodels.User, err error)
	Save(user datamodels.User) (err error)
	GetByID(id string) (user datamodels.User, err error)
	GetByName(Usename string) (user datamodels.User, err error)
	DeleteByID(id string) (err error)
	GetSignerByName(username string) (user datamodels.User, err error)
}

type usersRepository struct {
	collection string
}

func NewUsersRepository() UsersRepository {
	return &usersRepository{
		collection: "users",
	}
}

func (m *usersRepository) List(query map[string]interface{}) (users []datamodels.User, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(m.collection)

	if err = col.Find(nil).All(&users); err != nil {
		if err.Error() != datasource.GetErrNotFound().Error() {
			return
		}
	}
	return
}

func (m *usersRepository) Save(user datamodels.User) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(m.collection)

	if user.ID.Hex() == "" {
		user.ID = bson.NewObjectId()
		err = col.Insert(user)
	} else {
		err = col.Update(bson.M{"_id": user.ID}, user)
	}

	return
}

func (m *usersRepository) GetByID(id string) (user datamodels.User, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(m.collection)

	err = col.FindId(bson.ObjectIdHex(id)).One(&user)
	return
}

func (m *usersRepository) GetByName(Usename string) (user datamodels.User, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(m.collection)
	err = col.Find(bson.M{"name": "ASDFGH"}).One(&user) //Funciona
	err = col.Find(bson.M{"name": Usename}).One(&user)

	return
}

func (m *usersRepository) DeleteByID(id string) (err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(m.collection)

	err = col.RemoveId(bson.ObjectIdHex(id))
	return
}

func (m *usersRepository) GetSignerByName(username string) (user datamodels.User, err error) {
	db := datasource.NewSessionStore()
	defer db.Close()
	col := db.C(m.collection)
	//err = col.Find(bson.M{"name": "ASDFGH"}).One(&user) //Funciona
	err = col.Find(bson.M{"username": username}).One(&user)
	return
}
