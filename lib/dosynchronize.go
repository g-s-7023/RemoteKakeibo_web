package lib

import (
	"net/http"
	"encoding/json"
	"log"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine"
)

func DoSynchronize(w http.ResponseWriter, r *http.Request) {
	var err error
	// 必ず最後にbodyを閉じる
	defer r.Body.Close()
	// POSTメソッドでなければreturn
	if r.Method != "POST" {
		return
	}
	//===
	//=== クライアントから送られてきたjsonのパース
	//===
	var receivedKakeibo []Kakeibo
	if err = json.NewDecoder(r.Body).Decode(&receivedKakeibo); err != nil {
		return
	}
	//===
	//=== パースしたエントリから更新用のデータを作成
	//===
	ctx := appengine.NewContext(r)
	// パースしたエントリの情報を元に、挿入、更新用のデータ(配列)を作成
	// > エントリを挿入：クライアント側のエントリとidが一致するサーバ側のエントリがない場合
	// > エントリを更新：クライアント側のエントリとidが一致するサーバ側のエントリがあり、そのisSynchronizedがtrueの場合
	// > エントリを無視：クライアント側のエントリとidが一致するサーバ側のエントリがあり、そのisSynchronizedがfalseの場合
	var entriesForSync []ParamToSync
	var keysForDelete []*datastore.Key
	for _, cliVal := range receivedKakeibo {
		// 各エントリのIDがサーバ側のDBに既にあるかチェック
		key, year, month, isSync, err:= SearchIdFromClient(ctx, cliVal.Id)
		if err != nil {
			log.Fatal(err)
			continue
		}
		switch key {
		case nil:
			// 送信されたエントリとIDが一致するエントリがサーバ側にない場合、insert
			newKey := datastore.NewIncompleteKey(ctx, KAKEIBO_ENTRY, getMonthKey(appengine.NewContext(r), KakeiboId, cliVal.Year, cliVal.Month))
			entriesForSync = append(entriesForSync, ParamToSync{newKey, cliVal})
		default:
			// 送信されたエントリとIDが一致するエントリがサーバ側にある場合
			if isSync == TRUE {
				// isSynchronizedがTRUEならupdate、FALSEなら無視(=サーバ側優先)
				if cliVal.Year != year || cliVal.Month != month {
					// yearまたはmonthの値が変わるようであれば、そのkeyのエントリは消して新規に挿入
					keysForDelete = append(keysForDelete, key)
					newKey := datastore.NewIncompleteKey(ctx, KAKEIBO_ENTRY, getMonthKey(appengine.NewContext(r), KakeiboId, cliVal.Year, cliVal.Month))
					entriesForSync = append(entriesForSync, ParamToSync{newKey, cliVal})
				} else {
					// yearとmonthの値が変わらないようであれば、そのままエントリを更新
					entriesForSync = append(entriesForSync, ParamToSync{key, cliVal})
				}
			}
		}
	}
	//===
	//=== 更新用データをDBに登録
	//===
	Synchronize(ctx, entriesForSync ,keysForDelete)
	//===
	//=== クライアントで更新してもらうjsonデータを作成
	//===
	// クライアントに送信するエントリをDBから取得
	entriesForSend, err := getKakeiboToSend(ctx)
	if err != nil {
		return
	}
	// JSONへエンコード
	jsonBytes , err := json.Marshal(entriesForSend)
	if err != nil {
		return
	}
	//===
	//=== クライアントにjsonを送信
	//===
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
