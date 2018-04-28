package lib

import "net/http"

func DoSynchronize(w http.ResponseWriter, r *http.Request) {
	// jsonのパース


	// パースした情報(cliEnrty)を元に、挿入、更新用のデータ(配列)を作成
	// cliEntryを挿入：cliEntryとidが一致するsrvEntryがない場合
	// cliEntryを更新：cliEntryとidが一致するsrvEntryのisSynchronizedがtrueの場合
	// cliEntryを無視：cliEntryとidが一致するsrvEntryのisSynchronizedがfalseの場合

	// 更新用データをDBに登録

	// クライアントに送信するjsonデータを作成

	// クライアントにjsonを送信

}
