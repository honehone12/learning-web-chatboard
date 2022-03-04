package common

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func OpenDB(dbName string) (dbEngine *xorm.Engine, err error) {
	driver := "postgres"
	para := "dbname=%s user=%s password=%s host=localhost port=5432 sslmode=disable"
	dbEngine, err = xorm.NewEngine(driver,
		fmt.Sprintf(para, dbName, os.Getenv("DBUSER"), os.Getenv("DBPASS")))
	if err != nil {
		return
	}
	dbEngine.ShowSQL(true)
	dbEngine.SetMaxOpenConns(10)
	return
}
