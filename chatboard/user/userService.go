package user

import "xorm.io/xorm"

var engine *xorm.Engine

func OpenService(dbEngine *xorm.Engine) {
	engine = dbEngine
}
