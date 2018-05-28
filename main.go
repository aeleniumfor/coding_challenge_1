package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Hello struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil) //ポート8080で待機
}

//ハンドラ関数定義
func handler(w http.ResponseWriter, r *http.Request) {
	//内部処理
	msg := Hello{"Hello World!!"}
	json_s, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json") //ヘッダ情報付加
	fmt.Fprint(w, string(json_s))
}
