package config

import(

)


type Word struct {
	ID    int    `json:"id"`
	Word  string `json:"word"`
}

var wrodList = []Word {
	Word{ID: 1, Word: "おっ！今日早いじゃん！あんま無理しちゃダメだよ"},
	Word{ID: 2, Word: "座右の銘は『一球入魂』"},
	Word{ID: 3, Word: "だって考えてみてくださいよ！偏差値30の人が東大受かりますか？って話ですよ"},
	Word{ID: 4, Word: "運命の赤い糸はそこにあったんですね"},
	Word{ID: 5, Word: "エンジニアはデスマッチ"},
	Word{ID: 6, Word: "セブンイレブンのシャーペンだ！"},
	Word{ID: 7, Word: "現金触ったの久しぶりだよ"},
	Word{ID: 8, Word: "忙しくて、トイレ行くっていう選択肢がなかったんすよ"},
	Word{ID: 9, Word: "ファミリーってやつっすよ"},
	Word{ID: 10, Word: "【今日のラッキー☆ミカイ占い】グリーンのジャケットを着ているそこのアナタ！今日のアナタは絶好調！！！！どんなエラーにも立ち向かっていけるはず。"},
}

func main() {
	// ここでワードのランダム生成する
}