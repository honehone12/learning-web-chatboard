package common

import (
	"chatboard/thread"
	"fmt"
	"os"
	"os/user"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func OpenDB(dbName string, callSync bool) (dbEngine *xorm.Engine, err error) {
	driver := "postgres"
	para := "dbname=%s user=%s password=%s host=localhost port=5432 sslmode=disable"
	dbEngine, err = xorm.NewEngine(driver,
		fmt.Sprintf(para, dbName, os.Getenv("DBUSER"), os.Getenv("DBPASS")))
	if err != nil {
		return
	}
	dbEngine.ShowSQL(true)
	dbEngine.SetMaxOpenConns(10)
	if callSync {
		sync(dbEngine)
	}
	return
}

func sync(dbEngine *xorm.Engine) {
	dbEngine.Sync2(new(thread.Thread))
	dbEngine.Sync2(new(user.User))
}
