package main

import (
	"gopkg.in/mgo.v2"
)

// Database is a database connection structure
type Database struct {
	Session  *mgo.Session
	Database *mgo.Database
}

// DatabaseConnect is function for establish database connection
func DatabaseConnect(config *Config) (*Database, error) {
	session, err := mgo.Dial(config.Database.Host)
	if err != nil {
		return &Database{}, err
	}
	db := Database{}
	db.Session = session
	db.Database = session.DB(config.Database.Name)
	return &db, nil
}

// Disconnect is function for close the database connection
func (db Database) Disconnect() {
	db.Session.Close()
}
