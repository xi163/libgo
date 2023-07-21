package dbwraper

import (
	"github.com/xi123/libgo/utils/dbwraper/Gorm"
	"github.com/xi123/libgo/utils/dbwraper/Mongo"
	"github.com/xi123/libgo/utils/dbwraper/Redis"
	"github.com/xi123/libgo/utils/dbwraper/Sql"
	"github.com/xi123/libgo/utils/json"
)

var Wrap = &Wraper{}

type Wraper struct {
	Redis Redis.DB
	Mongo Mongo.DB
	Sql   Sql.DB
	Gorm  Gorm.DB
}

func Init(RedisConf, MongoConf, SqlConf, GormConf any) {
	json.Parse(json.Bytes(RedisConf), &Redis.Conf)
	json.Parse(json.Bytes(MongoConf), &Mongo.Conf)
	json.Parse(json.Bytes(SqlConf), &Sql.Conf)
	json.Parse(json.Bytes(GormConf), &Gorm.Conf)
	Wrap.Redis.Init(Redis.Conf)
	Wrap.Mongo.Init(Mongo.Conf)
	Wrap.Sql.Init(Sql.Conf)
	Wrap.Gorm.Init(Gorm.Conf)
}
