package config

import (
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/dto"
	databaseInterface "github.com/fajarardiyanto/flt-go-database/interfaces"
	databaseLib "github.com/fajarardiyanto/flt-go-database/lib"
	"github.com/fajarardiyanto/flt-go-logger/interfaces"
	"github.com/fajarardiyanto/flt-go-logger/lib"
)

var (
	logger   interfaces.Logger
	database databaseInterface.SQL
	rdb      databaseInterface.Redis
	mongoDB  databaseInterface.Mongo
)

func init() {
	logger = lib.NewLib()
	logger.Init(dto.GetConfig().Name)
}

func Config() {
	db := databaseLib.NewLib()
	db.Init(GetLogger())

	InitMysql(db)
	InitRedis(db)
	InitMongo(db)
}

func InitMysql(db databaseInterface.Database) {
	database = db.LoadSQLDatabase(dto.GetConfig().Database.Mysql)

	if err := database.LoadSQL(); err != nil {
		logger.Error(err).Quit()
	}
}

func InitRedis(db databaseInterface.Database) {
	rdb = db.LoadRedisDatabase(dto.GetConfig().Database.Redis)

	if err := rdb.Init(); err != nil {
		logger.Error(err)
		return
	}
}

func InitMongo(db databaseInterface.Database) {
	mongoDB = db.LoadMongoDatabase(dto.GetConfig().Database.Mongo)

	if err := mongoDB.Init(); err != nil {
		logger.Error(err)
		return
	}
}

func GetLogger() interfaces.Logger {
	return logger
}

func GetDBConn() databaseInterface.SQL {
	return database
}

func GetRdbConn() databaseInterface.Redis {
	return rdb
}

func GetMongoConn() databaseInterface.Mongo {
	return mongoDB
}
