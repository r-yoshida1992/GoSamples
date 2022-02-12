package main

import "github.com/jinzhu/gorm"

func main() {
	// DBに接続
	db := dbConnect()

	// 

	// 最後にDBをクローズ
	defer db.Close()

}

func dbConnect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
