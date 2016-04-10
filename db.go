package main

import r "github.com/dancannon/gorethink"

var dbSession *r.Session

type dbParams map[string]interface{}

type dbConfig struct {
	Address string
	DBName  string
}

func (c dbConfig) dial() error {
	var err error
	dbSession, err = r.Connect(r.ConnectOpts{
		Address:  c.Address,
		Database: c.DBName,
		MaxIdle:  10,
		MaxOpen:  10,
	})
	if err != nil {
		return err
	}
	c.createDB()
	c.createTables()
	return nil
}

func (c dbConfig) createDB() {
	r.DBCreate(c.DBName).Exec(dbSession)
	dbSession.Use(c.DBName)
}

func (c dbConfig) createTables() {
	r.TableCreate("users").Exec(dbSession)
	r.TableCreate("tokens").Exec(dbSession)
}
