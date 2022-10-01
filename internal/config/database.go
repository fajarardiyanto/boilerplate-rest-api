package config

import (
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/dto"
	databaseinterface "github.com/fajarardiyanto/flt-go-database/interfaces"
	databaseLib "github.com/fajarardiyanto/flt-go-database/lib"
)

var (
	database databaseinterface.SQL
)

func Init() {
	db := databaseLib.NewLib()
	db.Init(GetLogger())

	database = db.LoadSQLDatabase(dto.GetConfig().Mysql)

	if err := database.LoadSQL(); err != nil {
		logger.Error(err).Quit()
	}
}

func GetDBConn() databaseinterface.SQL {
	return database
}
