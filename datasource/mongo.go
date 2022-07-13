package datasource

import (
	"log"
	"tentativa/config"
	"time"

	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
)

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:          []string{config.GConfig.Mongo.Host},
		Direct:         false,
		Timeout:        time.Second * 60,
		FailFast:       false,
		Database:       config.GConfig.Mongo.Name,
		ReplicaSetName: "",
		Source:         "",
		Service:        "",
		ServiceHost:    "",
		Mechanism:      "",
		Username:       "",
		Password:       "",
		PoolLimit:      4096,
		DialServer:     nil,
		Dial:           nil,
	}
	sess, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	session = sess
	session.SetMode(mgo.Monotonic, true)
}

type SessionStore struct {
	session *mgo.Session
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		session: session.Copy(),
	}
}

func (s *SessionStore) C(name string) *mgo.Collection {
	return s.session.DB(config.GConfig.Mongo.Name).C(name)
}

func (s *SessionStore) Close() {
	s.session.Close()
}

func GetErrNotFound() error {
	return mgo.ErrNotFound
}
