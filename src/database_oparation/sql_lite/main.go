package main

import (
	"GoSamples/src/database_oparation/base/common"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// DBに接続
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	// 最後にDBをクローズ
	defer db.Close()

	// DB操作部分については以下の関数内を参照
	common.Operation(db)
}
