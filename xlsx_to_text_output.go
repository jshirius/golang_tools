package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/tealeg/xlsx"
)

const MaxCol = 14 //列の最大値

func main() {

	var excelFile string
	excelFile = ""

	//コマンドライン取得
	if len(os.Args) >= 2 {
		//合成対象の画像を入っているディレクトリRoot
		excelFile = os.Args[1]

	} else {
		fmt.Println("コマンドライン引数に元になるExcelファイルを指定してください")
		return
	}

	//ファイルを開く
	excelFileName := excelFile
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		//...
		fmt.Print("存在しないExcelファイルです。確認してください")
		return
	}

	//----------------------------------
	//読み込み
	//----------------------------------

	for _, sheet := range xlFile.Sheets {

		//タイトル
		titles := [MaxCol]string{}

		//print(sheet.MaxRow)
		//行に対する読み込み
		for i := 0; i < sheet.MaxRow; i++ {

			//内容
			infos := [MaxCol]string{}

			//列に対する読み込み
			for j := 0; j < sheet.MaxCol; j++ {
				var cell *xlsx.Cell
				cell = sheet.Cell(i, j)
				//print(cell.Value)

				//タイトルを入れるか？
				if i == 0 {
					titles[j] = cell.Value

				} else {
					//２行目以降
					infos[j] = cell.Value
				}

			}

			//ファイル出力？
			//１列目が0以上であること
			if i > 0 {
				if len(infos[0]) > 0 {
					//ファイル開く
					fileName := infos[1] + ".txt"
					fmt.Println(fileName)
					outPutTextFile(fileName, titles, infos)
				}

			}

		}

		break
	}

}

//配列のコピーになっているから、どこかのタイミングでスライスに変更しよう
func outPutTextFile(fileName string, titles [MaxCol]string, infos [MaxCol]string) {

	//ファイルオープン
	fp := newFile(fileName)
	defer fp.Close()
	writer := bufio.NewWriter(fp)

	//BOMを追加する
	writer.Write([]byte{0xEF, 0xBB, 0xBF})
	writeString := ""
	//for j := 0; j < len(infos); j++ {

	//ファイル名
	writeString += infos[1] + "\n\n"

	//キーワード出力
	writeString += "■キーワード \n"
	writeString += infos[2] + " " + infos[3] + " " + infos[4] + " " + infos[5] + "\n\n"
	//fmt.Printf("%s\n", text)
	//ターゲット読者出力
	writeString += "■ターゲット読者(検索意図) \n"
	writeString += infos[6] + "\n\n"

	//タイトル
	writeString += ".■タイトル\n"
	writeString += infos[7] + "\n\n"

	//Div
	writeString += "[div]\n\n"

	//見出し
	writeString += ".■見出し\n"
	writeString += infos[8] + "\n\n"

	//センテンス1
	writeString += ".#第1段落\n"
	writeString += infos[9] + "\n\n"

	//センテンス2
	writeString += ".#第2段落\n"
	writeString += infos[10] + "\n\n"

	//センテンス3
	writeString += ".#第3段落\n"
	writeString += infos[11] + "\n\n"

	//まとめ
	writeString += ".#まとめ\n"
	writeString += infos[12] + "\n\n"

	//div
	writeString += "[/div]\n"

	//備考
	writeString += "--------備考----------- \n"
	writeString += infos[13] + "\n\n"

	//}
	_, err := writer.WriteString(writeString)
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush()
}

func newFile(fn string) *os.File {
	fp, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return fp
}
