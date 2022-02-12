package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// Operation DBの操作部分
func Operation(db *gorm.DB) {

	// テーブルの作成
	db.AutoMigrate(&TestTable{})

	// データのインサート(3件)
	db.Create(&TestTable{UserName: "taro", Password: "password"})
	db.Create(&TestTable{UserName: "jiro", Password: "123456"})
	db.Create(&TestTable{UserName: "saburo", Password: "qwerty"})

	// レコードの取得(全件)
	var allRecords []TestTable
	db.Find(&allRecords)            // SELECT * FROM test_table;
	fmt.Println("1 : ", allRecords) // [{taro password} {jiro 123456} {saburo qwerty}]

	// レコードの取得(最初の1件)
	var oneRecord TestTable
	db.First(&oneRecord)           // SELECT * FROM test_table LIMIT 1;
	fmt.Println("2 : ", oneRecord) // {taro password}

	// レコードの取得(条件指定)
	var whereRecords []TestTable
	db.Find(&whereRecords, "user_name = ?", "jiro") // SELECT * FROM test_table WHERE user_name = 'jiro';
	fmt.Println("3 : ", whereRecords)               // [{jiro 123456}]

	// レコードの更新
	db.Model(&allRecords).Where("user_name = ?", "jiro").Update("password", "abcdefg")
	db.Find(&allRecords)            // 全レコード再取得
	fmt.Println("4 : ", allRecords) // [{taro password} {jiro abcdefg} {saburo qwerty}]

	// レコードの削除
	db.Delete(allRecords)

}

// TestTable テーブルの構造体。
// type名がテーブル名となり、その中にカラムを定義する。
type TestTable struct {
	UserName string
	Password string
}
