package mongo

import (
	"github.com/kildevaeld/go-acl"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoStore struct {
	c *mgo.Collection
}

func (self *MongoStore) GetRole(name string) *acl.Role {
	var role *acl.Role
	self.c.Find(bson.M{"name": name}).One(role)
	return role
}

func (self *MongoStore) AddRole(role *acl.Role) error {
	return self.c.Insert(role)
}

func (self *MongoStore) UpdateRole(role *acl.Role) error {
	self.c.Update(bson.M{"name":role.Name}, role)
	return nil
}

func (self *MongoStore) RemoveRole(role *acl.Role) error {
	self.c.Remove(bson.M{"name":role.Name})
	return nil
}

func NewMongoStore(c *mgo.Collection) *MongoStore {
	
	index := mgo.Index{
		Key: []string{"name"},
		Unique: true,
		DropDups: true,
		Background: true,
	}
	
	c.EnsureIndex(index)
	
	return &MongoStore{c}
}
