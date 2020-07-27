package model

import (
	"database/sql"
	"fmt"
	"warmup/config"
	_ "github.com/go-sql-driver/mysql"
)

var MySQLClient *sql.DB

func ConnectMySQL() error {
	dbUsr := "yokel"
	dbPwd := "yokel"
	dbName := "project"
	var err error
	MySQLClient, err = sql.Open("mysql", dbUsr+":"+dbPwd+"@/"+dbName)
	if err != nil {
		return err
	}
	return nil
}

func ShowAllUser() ([]config.RegisterInfo, error) {
	selDB, err := MySQLClient.Query("SELECT * FROM user ")
	if err != nil {
		return []config.RegisterInfo{}, err
	}
	res := []config.RegisterInfo{}
	for selDB.Next() {
		var tmp config.RegisterInfo
		err = selDB.Scan(&tmp.Uid, &tmp.Pwd, &tmp.Nickname, &tmp.Phone, &tmp.Email, &tmp.Type)
		if err != nil {
			return []config.RegisterInfo{}, err
		}
		res = append(res, tmp)
	}
	return res, nil
}

func FindUser(userId string) (config.RegisterInfo, error) {
	selDB, err := MySQLClient.Query("SELECT * FROM user WHERE uid=?", userId)
	if err != nil {
		return config.RegisterInfo{}, err
	}
	var tmp config.RegisterInfo
	for selDB.Next() {
		err = selDB.Scan(&tmp.Uid, &tmp.Pwd, &tmp.Nickname, &tmp.Phone, &tmp.Email, &tmp.Type)
		if err != nil {
			return config.RegisterInfo{}, err
		}
	}
	return tmp, nil
}

func UpdateUser(UserInfo config.RegisterInfo) error {
	insForm, err := MySQLClient.Prepare("UPDATE user SET pwd=?, nickname=?, phone=?, email=? WHERE uid=?")
	if err != nil {
		return err
	}
	fmt.Println(UserInfo)
	insForm.Exec(UserInfo.Pwd, UserInfo.Nickname, UserInfo.Phone, UserInfo.Email,UserInfo.Uid)
	return nil
}

func DeleteUser(userId string) error {
	delForm, err := MySQLClient.Prepare("DELETE FROM user WHERE uid=?")
	if err != nil {
		return err
	}
	delForm.Exec(userId)
	return nil
}

func InsertNewUser(UserInfo config.RegisterInfo) error {
	insForm, err := MySQLClient.Prepare("INSERT INTO user(uid, pwd,nickname,phone,email,type) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	insForm.Exec(UserInfo.Uid, UserInfo.Pwd, UserInfo.Nickname, UserInfo.Phone, UserInfo.Email, UserInfo.Type)
	return nil
}

func init() {
	err := ConnectMySQL()
	if err != nil {
		panic(err)
	}
	err = ConnectRedisDataBase()
	if err != nil {
		panic(err)
	}
}
