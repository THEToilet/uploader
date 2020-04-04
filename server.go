package main

import (
	format "fmt"
	"html/template"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var mux *http.ServeMux
	mux = http.NewServeMux()

	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/", index)

	// http.Server構造体のポインタを宣言
	var server *http.Server
	// http.Serverのオブジェクトを確保
	// &をつけること構造体ではなくポインタを返却
	server = &http.Server{} // or new (http.Server);
	server.Addr = ":11180"
	server.Handler = mux
	server.ListenAndServe()
}

func index(writer http.ResponseWriter, request *http.Request) {
	var t *template.Template
	// テンプレートをロード
	t, _ = template.ParseFiles("template/index.html")
	t.Execute(writer, struct{}{})
}

func upload(w http.ResponseWriter, r *http.Request) {
	// このハンドラ関数へのアクセスはPOSTメソッドのみ認める
	if r.Method != "POST" {
		format.Fprintln(w, "許可したメソッドとはことなります。")
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
		format.Fprintln(w, "ファイルアップロードを確認できませんでした。")
		return
	}
	uploadedFileName = fileHeader.Filename
	// サーバー側に保存するために空ファイルを作成
	var saveImage *os.File
	saveImage, e = os.Create("./resources/" + uploadedFileName)
	if e != nil {
		format.Fprintln(w, "サーバ側でファイル確保できませんでした。")
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
			format.Println(e)
			break
		}
		if e != nil {
			format.Println(e)
			format.Fprintln(w, "アップロードされたファイルデータのコピーに失敗。")
			return
		}
		saveImage.WriteAt(img, tempLength)
		tempLength = int64(n) + tempLength
	}
	format.Fprintf(w, "文字列HTTPとして出力させる")
}
