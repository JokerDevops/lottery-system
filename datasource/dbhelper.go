package datasource

import (
	"fmt"
	"log"
	"lottery-system/conf"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var masterInstence *xorm.Engine
var dbLock sync.Mutex

// 创建单例
func InstanceDbMaster() *xorm.Engine {
	if masterInstence != nil {
		return masterInstence
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	// 再次检查
	if masterInstence != nil {
		return masterInstence
	}
	return NewDbMaster()

}

// 实例化 xorm 引擎
func NewDbMaster() *xorm.Engine {
	sourcename := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Pwd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database,
	)
	instance, err := xorm.NewEngine(conf.DriverName, sourcename)
	if err != nil {
		log.Fatal("dbhelper.NewDbMaster: NewEngine error", err)
		return nil
	}
	instance.ShowSQL(true)
	masterInstence = instance
	return instance
}
