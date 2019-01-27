package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tealeg/xlsx"
)

const MaxCol = 14 //列の最大値

const Name = "一行目に出力する文字列"
const FileName = "ファイル名"

/*
一列目 関連キーワード1,2,3,4
３行目 知りたいこと・疑問

■買う人
■タイトル

*/

func main() {

	var excelFile string
	excelFile = ""

	//コマンドライン取得
	if len(os.Args) >= 2 {
		//エクセルファイル
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

		//行に対する読み込み
		for i := 0; i < sheet.MaxRow; i++ {

			//内容
			infos := [MaxCol]string{}

			print(sheet.MaxCol)
			//列に対する読み込み
			for j := 0; j < sheet.MaxCol; j++ {

				if MaxCol <= j {
					break
				}
				var cell *xlsx.Cell
				cell = sheet.Cell(i, j)
				//print(cell.Value)

				infos[j] = cell.Value

			}

			//ファイル出力？
			//１列目が0以上であること
			if i > 0 {
				if len(infos[0]) > 0 {
					//ファイル開く
					fileName := FileName + strconv.Itoa(i) + ".txt"
					fmt.Println(fileName)
					outPutTextFile(fileName, infos)
				}

			}

		}

		break
	}

}

//配列のコピーになっているから、どこかのタイミングでスライスに変更しよう
func outPutTextFile(fileName string, infos [MaxCol]string) {

	//ファイルオープン
	fp := newFile(fileName)
	defer fp.Close()
	writer := bufio.NewWriter(fp)

	//BOMを追加する
	writer.Write([]byte{0xEF, 0xBB, 0xBF})
	writeString := ""
	//for j := 0; j < len(infos); j++ {

	//キーワード出力
	writeString += "■キーワード \n"
	writeString += infos[0] + " " + infos[1] + " " + infos[2] + " " + infos[3] + "\n\n"
	//fmt.Printf("%s\n", text)

	//タイトル
	writeString += ".■タイトル\n"
	writeString += infos[5] + "\n\n"

	//Div
	writeString += "[div]\n\n"

	//見出し
	writeString += ".■見出し\n"
	writeString += infos[4] + "\n\n"

	//センテンス1
	writeString += ".#第1段落\n"

	//センテンス2
	writeString += ".#第2段落\n"

	//センテンス3
	writeString += ".#第3段落\n"

	//まとめ
	writeString += ".#まとめ\n"

	//div
	writeString += "[/div]\n"

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
