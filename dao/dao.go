package dao

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/martinroddam/lists/model"
)

type ListsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	users_collection = "users"
	lists_collection = "lists"
	tasks_collection = "tasks"
)

func (l *ListsDAO) Connect() {
	session, err := mgo.Dial(l.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(l.Database)
}

func (l *ListsDAO) InsertUser(user model.User) error {
	err := db.C(users_collection).Insert(&user)
	return err
}

func (l *ListsDAO) FindUserById(id string) (model.User, error) {
	var user model.User
	err := db.C(users_collection).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Find all users
func (l *ListsDAO) FindAllUsers() ([]model.User, error) {
	var users []model.User
	err := db.C(users_collection).Find(bson.M{}).All(&users)
	return users, err
}

func (l *ListsDAO) InsertList(list model.List) error {
	err := db.C(lists_collection).Insert(&list)
	return err
}

func (l *ListsDAO) FindListById(id string) (model.List, error) {
	var list model.List
	err := db.C(lists_collection).FindId(bson.ObjectIdHex(id)).One(&list)
	return list, err
}

// Find all lists
func (l *ListsDAO) FindAllLists() ([]model.List, error) {
	var lists []model.List
	err := db.C(lists_collection).Find(bson.M{}).All(&lists)
	return lists, err
}
