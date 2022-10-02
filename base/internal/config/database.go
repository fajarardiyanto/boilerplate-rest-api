package config

import (
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/dto"
	databaseinterface "github.com/fajarardiyanto/flt-go-database/interfaces"
	databaseLib "github.com/fajarardiyanto/flt-go-database/lib"
)

var (
	database databaseinterface.SQL
	rdb      databaseinterface.Redis
)

func Init() {
	db := databaseLib.NewLib()
	db.Init(GetLogger())

	InitMysql(db)
	InitRedis(db)
}

func InitMysql(db databaseinterface.Database) {
	database = db.LoadSQLDatabase(dto.GetConfig().Database.Mysql)

	if err := database.LoadSQL(); err != nil {
		logger.Error(err).Quit()
	}
}

func InitRedis(db databaseinterface.Database) {
	rdb = db.LoadRedisDatabase(dto.GetConfig().Database.Redis)

	if err := rdb.Init(); err != nil {
		logger.Error(err)
		return
	}
}

func GetDBConn() databaseinterface.SQL {
	return database
}

func GetRdbConn() databaseinterface.Redis {
	return rdb
}
