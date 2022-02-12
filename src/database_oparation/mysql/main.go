package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //直接的な記述が無いが、インポートしたいものに対しては"_"を頭につける決まり
	"github.com/jinzhu/gorm"
	"io/ioutil"
)

func main() {
	// DBに接続
	db := gormConnect()
	// ModelとDB名を一致させて取得できるようにする
	db.SingularTable(true)
	// メソッド終了時にDBをクローズする
	defer db.Close()
	// 結果格納用の変数
	var records []User
	// 全件取得
	db.Find(&records)
	fmt.Println(records)
}

// dbsetting.jsonから接続設定を読み込む
func readConfig() DbConfig {
	raw, err := ioutil.ReadFile("dbConfig.json")
	if err != nil {
		panic(err)
	}

	var conf DbConfig
	json.Unmarshal(raw, &conf)
	return conf
}

/**
 * DBに接続
 * DBの各種設定はdbConfig.jsonを適宜変更する
 */
func gormConnect() *gorm.DB {
	conf := readConfig()
	DBMS := "mysql"
	USER := conf.User
	PASS := conf.Password
	DBNAME := conf.DbName
	HOSTNAME := conf.HostName
	PORT := conf.Port
	// MySQLだと文字コードの問題で"?parseTime=true"を末尾につける必要がある
	CONNECT := USER + ":" + PASS + "@tcp(" + HOSTNAME + ":" + PORT + ")/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

// User テーブルのmodel
type User struct {
	UserName string
	Password string
}

// DbConfig dbConfig.jsonのmodel
type DbConfig struct {
	User     string
	Password string
	DbName   string
	HostName string
	Port     string
}
