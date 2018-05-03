package doa

import (
	. "github.com/jroyals2/firstFullStackGo/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type ItemDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "items"
)

func (i *ItemDAO) Connect() {
	session, err := mgo.Dial(i.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(i.Database)
}

func (i *ItemDAO) FindAll() ([]Item, error) {
	var items []Item
	err := db.C(COLLECTION).Find(bson.M{}).All(&items)
	return items, err
}

func (i *ItemDAO) FindById(id string) (Item, error) {
	var item Item
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex).One(&item)
	return item, err
}

func (i *ItemDAO) Insert(item Item) error {
	err := db.C(COLLECTION).Insert(&item)
	return err
}

func (i *ItemDAO) Delete(item Item) error {
	err := db.C(COLLECTION).Remove(&item)
	return err
}

func (i *ItemDAO) Update(item Item) error {
	err := db.C(COLLECTION).UpdateId(item.ID, &item)
	return err
}
