package main

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
)

func list(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(dirwalk("./resources"))
	fmt.Fprintln(writer, dirwalk("./resources"))
	files := dirwalk("./resources")
	for _, file := range files {
		fmt.Fprintf(writer, fmt.Sprintf("./%s", file))
		fmt.Fprintf(writer, readfile("./"+file))
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "index", "navbar", "content")
}

func viewlist(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, dirwalk("./resources"), "index", "navbar", "list")
}
func upload(w http.ResponseWriter, r *http.Request) {
	// このハンドラ関数へのアクセスはPOSTメソッドのみ認める
	if r.Method != "POST" {
		fmt.Fprintln(w, "許可したメソッドとはことなります。")
		return
	}
	var file multipart.File
	var fileHeader *multipart.FileHeader
	var e error
	var uploadedFileName string
	var img []byte = make([]byte, 1024)
	// POSTされたファイルデータを取得する
	file, fileHeader, e = r.FormFile("image")
	if e != nil {
		fmt.Fprintln(w, "ファイルアップロードを確認できませんでした。")
		return
	}
	uploadedFileName = fileHeader.Filename
	// サーバー側に保存するために空ファイルを作成
	var saveImage *os.File
	saveImage, e = os.Create("./resources/" + uploadedFileName)
	if e != nil {
		fmt.Fprintln(w, "サーバ側でファイル確保できませんでした。")
		return
	}
	defer saveImage.Close()
	defer file.Close()
	var tempLength int64 = 0
	for {
		// 何byte読み込んだかを取得
		n, e := file.Read(img)
		// 読み混んだバイト数が0を返したらループを抜ける
		if n == 0 {
			fmt.Println(e)
			break
		}
		if e != nil {
			fmt.Println(e)
			fmt.Fprintln(w, "アップロードされたファイルデータのコピーに失敗。")
			return
		}
		saveImage.WriteAt(img, tempLength)
		tempLength = int64(n) + tempLength
	}
	fmt.Fprintf(w, "文字列HTTPとして出力させる")
}
