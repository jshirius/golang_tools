package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type DiaryLog struct {
	// db tag lets you specify the column name if it differs from the struct field
	Id       int64     `db:"id, primarykey, autoincrement"`
	Title    string    `db:"title,size:45"` // Column size set to 45
	Body     string    `db:"body,size:1024"`
	Created  time.Time `db:"created"`
	Modified time.Time `db:"modified"`
}

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

	// Select
	/*
		rows, err := db.Query("select * from t_diary_logs")
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()

		// 構造体へマッピング

			for rows.Next() {

				datas, _ := rows.Columns()
				for i, col := range datas {
					fmt.Printf("index: %d, name: %s\n", i, col)
				}

				//user := User{}
				//err := rows.Scan(&user.Id, &user.Name)
			}
	*/

	//insert
	p1 := newData("title", "body", time.Now().AddDate(-1, 0, 0))
	err = dbmap.Insert(&p1)
	checkErr(err, "Insert failed")
}
