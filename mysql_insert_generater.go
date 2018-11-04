//------------------------------------------------
//データ作成ツール
//データは、以下の条件で作成する
//テキストデータ(body)は、１万文字
//2014,2015,2016,2017,2018でそれぞれ１万件のデータを作成(４年前までのデータ作成)
//------------------------------------------------

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

//定義
const MakeDataCount = 10000 //作成するデータ数
const BodySize = 10000

//DBのテーブルに対応するデータ型の定義
type DiaryLog struct {
	Id       int64     `db:"id, primarykey, autoincrement"`
	Title    string    `db:"title,size:45"` // Column size set to 45
	Body     string    `db:"body,size:1024"`
	Created  time.Time `db:"created"`
	Modified time.Time `db:"modified"`
}

//insert用に作成
func newData(title string, body string, created time.Time) DiaryLog {
	return DiaryLog{
		Title:    title,
		Created:  created,
		Modified: time.Now(),
		Body:     body,
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func main() {

	//mysqlに接続する
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err.Error())
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(DiaryLog{}, "t_diary_logs").SetKeys(true, "Id")
	defer dbmap.Db.Close()

	//文字列を作成する
	var bodyDummyData string
	bodyDummyData = "body_"
	for i := 0; i < BodySize; i++ {
		bodyDummyData += "a"
	}
	bodyDummyData += "_end"

	//４年分の繰り返し
	//テストデータを生成する
	var yearCount = 0
	for i := 0; i <= 4; i++ {

		fmt.Println("%d年目のデータ生成開始", i)

		for j := 0; j < MakeDataCount; j++ {
			p1 := newData("title_"+time.Now().String(), bodyDummyData, time.Now().AddDate(yearCount, 0, 0))
			err = dbmap.Insert(&p1)
			checkErr(err, "Insert failed")
		}

		yearCount -= 1
	}

}
