//MySqlの性能チェックのためにだけに作ったツール
//MySqlに「InsertCount」に設定した分だけデータを追加して、追加にかかった時間を計測
//さらに読み出し時間の計測

package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

//定義
const InsertCount = 10000

type Person struct {
	ID         int64     `db:"id, primarykey, autoincrement"`
	Name       string    `db:"name,size:45"`   // Column size set to 45
	Adress     string    `db:"adress,size:45"` // Column size set to 45
	Str2       string    `db:"str2,size:1024"` // Column size set to 1024
	Age        int       `db:"age"`
	Int1       int       `db:"int1"`
	Int2       int       `db:"int2"`
	Int3       int       `db:"int3"`
	Int4       int       `db:"int4"`
	UpdateTime time.Time `db:"update_time"`
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func main() {

	//時間計測用に現在の時間を取得
	bf_t := time.Now()
	fmt.Println(bf_t)

	//mysqlに接続する
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err.Error())
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Person{}, "t_persons").SetKeys(true, "ID")
	defer dbmap.Db.Close()

	//データの追加
	count := 0
	for i := 0; i < InsertCount; i++ {

		//データ追加
		ritsu := &Person{
			Name:       "テスター" + strconv.Itoa(i),
			Adress:     "東京都新宿区AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" + strconv.Itoa(i),
			Str2:       "文字列AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" + strconv.Itoa(i),
			Age:        i,
			Int1:       i + 1,
			Int2:       i + 2,
			Int3:       i + 3,
			Int4:       i + 4,
			UpdateTime: time.Now(),
		}

		//p1 := newData("title_"+time.Now().String(), bodyDummyData, time.Now().AddDate(yearCount, 0, 0))
		err = dbmap.Insert(ritsu)
		checkErr(err, "Insert failed")

		count++
	}

	//データの追加終わり
	//かかった時間を出力
	af_t := time.Now()
	fmt.Println(strconv.Itoa(count) + "件インサートにかかった時間")
	fmt.Println(af_t.Sub(bf_t))

	//以下のように*でカラムをしてもなぜかデータが取れないので注意が必要
	//rows, err := dbmap.Select(Person{}, `SELECT * FROM t_persons;`)
	rows, err := dbmap.Select(Person{}, `SELECT id,name,age,adress,str2 FROM t_persons where age < 10000`)

	/*
		for i, r := range rows {
			row := r.(*Person)
			fmt.Printf("[%d] id: %d, name: %s\n", i, row.ID, row.Name)
		}
	*/
	bf_t = time.Now()
	fmt.Println(strconv.Itoa(len(rows)) + "件取得までにかかった時間")
	fmt.Println(bf_t.Sub(af_t))

}
