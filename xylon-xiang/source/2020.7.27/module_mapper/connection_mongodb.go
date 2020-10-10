package module_mapper

import (
	"2020.7.27/config"
	"database/sql"
	"log"
)

var UserSysDB *sql.DB

func init() {
	var err error

	UserSysDB, err = sql.Open("mysql", config.Config.Mysql.UserName+":"+
		config.Config.Mysql.UserPassword+"@tcp("+config.Config.Mysql.Host+")/"+config.Config.Mysql.DatabaseName)
	if err != nil {
		log.Fatal(err)
	}
}
