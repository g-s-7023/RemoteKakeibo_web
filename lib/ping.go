package lib

import (
	"net/http"
	"encoding/json"
	"log"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	log.Printf("func Ping start")
	// メッセージ
	type msg struct{
		Name string
		Code int
	}
	msgArray := make([]msg, 2)
	//===
	//=== 受信
	//===
	// 必ず最後にbodyを閉じる
	defer r.Body.Close()
	// POSTメソッドでなければreturn
	if r.Method != "POST" {
		return
	}
	var msgReceived []msg
	if err := json.NewDecoder(r.Body).Decode(&msgReceived); err != nil {
		return
	}
	msgArray[0] = msgReceived[0]
	//===
	//=== 送信
	//===
	// 送信用の構造体
	msgToSend := struct {
		Name string
		Code int
	}{
		Name: "test from server",
		Code: 1,
	}
	msgArray[1] = msgToSend
	// JSONへエンコード
	jsonBytes, err := json.Marshal(msgArray)
	if err != nil {
		return
	}
	//===
	//=== クライアントにjsonを送信
	//===
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	log.Printf("message sent 1:%s, 2:%s", msgArray[0].Name, msgArray[1].Name)
}
