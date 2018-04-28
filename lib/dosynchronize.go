package lib

import (
	"net/http"
	"encoding/json"
)

func DoSynchronize(w http.ResponseWriter, r *http.Request) {
	var err error
	// 必ず最後にbodyを閉じる
	defer r.Body.Close()
	// jsonのパース
	var receivedKakeibo []Kakeibo
	if r.Method == "POST"{
		if err = json.NewDecoder(r.Body).Decode(&receivedKakeibo); err != nil{




		}
	}
	// パースした情報(cliEnrty)を元に、挿入、更新用のデータ(配列)を作成
	// cliEntryを挿入：cliEntryとidが一致するsrvEntryがない場合
	// cliEntryを更新：cliEntryとidが一致するsrvEntryのisSynchronizedがtrueの場合
	// cliEntryを無視：cliEntryとidが一致するsrvEntryのisSynchronizedがfalseの場合

	// 更新用データをDBに登録

	// クライアントに送信するjsonデータを作成

	// クライアントにjsonを送信

}
