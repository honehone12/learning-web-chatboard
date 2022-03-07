package user

import "xorm.io/xorm"

var engine *xorm.Engine

// engin can be separated from thread db
func OpenService(dbEngine *xorm.Engine) {
	engine = dbEngine
}
