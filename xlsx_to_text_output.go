package main

import (
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

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
		titles := [12]string{}

		print(sheet.MaxRow)
		//行に対する読み込み
		for i := 0; i < sheet.MaxRow; i++ {

			//内容
			infos := [12]string{}

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

			for _, word := range infos {
				if len(word) > 0 {
					fmt.Println(word)
				}

			}

			//ファイル出力？
			if i > 0 {
				//ファイル開く

				//タイトル出力

				//内容出力

				//ファイル閉じる
			}

			//列の読み込み
			/*
				for _, cell := range row.Cells {
					text := cell.String()
					if len(text) > 0 {
						fmt.Printf("%s\n", text)
					}

				}
			*/

			//内容の読み込み
			//テキストファイル作成

		}

		//BOMの書き出しサンプル
		//https://pinzolo.github.io/2017/03/29/utf8-csv-with-bom-on-golang.html
		/*
			for _, row := range sheet.Rows {
				for _, cell := range row.Cells {
					text := cell.String()
					if len(text) > 0 {
						//fmt.Printf("%s\n", text)
					}

				}
			}
		*/

	}

}
