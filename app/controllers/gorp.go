package controllers

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/laser/revel-todos/app/models"
	_ "github.com/mattn/go-sqlite3"
	r "github.com/revel/revel"
	"github.com/revel/revel/modules/db/app"
	"log"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	t := Dbm.AddTableWithName(models.Todo{}, "todos")
	t.SetKeys(true, "Id")

	Dbm.TraceOn("[gorp]", r.INFO)
	err := Dbm.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	todos := []*models.Todo{
		&models.Todo{"Call Dad", true, 0},
		&models.Todo{"Do dishes", false, 0},
		&models.Todo{"Call Mom", false, 0},
	}

	for _, todo := range todos {
		if err := Dbm.Insert(todo); err != nil {
			panic(err)
		}
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

type Gorp struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *Gorp) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *Gorp) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *Gorp) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
