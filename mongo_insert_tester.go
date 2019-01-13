//MongoDBの性能チェックのためにだけに作ったツール
//MongoDBに「InsertCount」に設定した分だけデータを追加して、追加にかかった時間を計測
//さらに読み出し時間の計測

package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID      bson.ObjectId `bson:"_id"`
	Name    string        `bson:"name"`
	Adress  string        `bson:"adress"`
	Str2    string        `bson:"str2"`
	Age     int           `bson:"age"`
	Int1    int           `bson:"int1"`
	Int2    int           `bson:"int2"`
	Int3    int           `bson:"int3"`
	Int4    int           `bson:"int4"`
	Nowtime time.Time     `bson:"update_time"`
}

const InsertCount = 10000

func main() {

	//時間計測用に現在の時間を取得
	bf_t := time.Now()
	fmt.Println(bf_t)

	//mongoDBに接続する
	session, _ := mgo.Dial("mongodb://localhost/test")
	defer session.Close()
	db := session.DB("test")

	//Collectionの作成
	col := db.C("people")
	count := 0

	//データのインサート
	for i := 0; i < InsertCount; i++ {
		ritsu := &Person{
			ID:      bson.NewObjectId(),
			Name:    "テスター" + strconv.Itoa(i),
			Adress:  "東京都新宿区AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" + strconv.Itoa(i),
			Str2:    "文字列AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" + strconv.Itoa(i),
			Age:     i,
			Int1:    i + 1,
			Int2:    i + 2,
			Int3:    i + 3,
			Int4:    i + 4,
			Nowtime: time.Now(),
		}

		if err := col.Insert(ritsu); err != nil {
			log.Fatalln(err)
		}

		count++
	}

	//データの追加終わり
	//かかった時間を出力
	af_t := time.Now()
	fmt.Println(strconv.Itoa(count) + "件インサートにかかった時間")
	fmt.Println(af_t.Sub(bf_t))

	/**
	 * コレクションのデータをすべて取得する
	**/
	query := db.C("people").Find(bson.M{})
	var persons []Person
	query.All(&persons)

	bf_t = time.Now()
	fmt.Println(strconv.Itoa(len(persons)) + "件取得までにかかった時間")
	fmt.Println(bf_t.Sub(af_t))

	//ついでなのですべて出力してみる
	//for _, p := range persons {
	//	fmt.Printf("data:%s\n", p)
	//}

}
