package tool

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"webpro/model"
)

//x-orm 数据库连接
type Orm struct {
	*xorm.Engine
}

var DbEngine *Orm

//读取json配置,连接数据库,返回orm引擎
func OrmEngine(cfg *Config) (*Orm, error) {
	//读取数据源配置
	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" +
		database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		log.Println("xorm.NewEngine() err: ", err)
		return nil, err
	}
	//sql展示
	engine.ShowSQL(database.ShowSql)

	//数据库表映射实体类, 可映射多张表...
	err = engine.Sync2(new(model.SmsCode),
		new(model.Member))
	if err != nil {
		log.Println("engine.Sync2() err: ", err)
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm, nil
}
