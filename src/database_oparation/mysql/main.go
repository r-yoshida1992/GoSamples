package main

import (
	"GoSamples/src/database_oparation/base/common"
	_ "github.com/go-sql-driver/mysql" //直接的な記述が無いが、インポートしたいものに対しては"_"を頭につける決まり
	"github.com/jinzhu/gorm"
)

func main() {
	// 接続文字列
	// {ユーザ名} + ":" + {パスワード} + "@tcp(" + {ホスト名} + ":" + {ポート} + ")/" + {データベース} + "?parseTime=true"
	connect := "test:test@tcp(localhost:3306)/test?parseTime=true"
	// DBに接続
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		panic(err.Error())
	}
	// 最後にDBをクローズ
	defer db.Close()

	// DB操作部分については以下の関数内を参照
	common.Operation(db)
}
