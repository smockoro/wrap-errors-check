package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	fmt.Println(targetWalk("./target"))
}

func targetWalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var paths []string
	for _, file := range files {
		if !file.IsDir() {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
		paths = append(paths, targetWalk(filepath.Join(dir, file.Name()))...)
	}

	return paths
}

//TODO
// テストコードでフォルダウォークするようにする ioutli系を使ったテストをする？
// 上記はどこかに参考資料があったので探して内容を理解してコードに残す。

// ウォークしてファイル *.goのファイルだけを取得する。
// 取得したファイルに対して構文解析をかけて errをlogに履いている部分を切り出す
// 構文木的にerrorsでラップされていないければ問題なのでチェックしておく
// 最後の結果として何個のうち何個が通っているのか、
// 間違っている場所はどのソースの何行目何列目なのか出力レポート出力

// errorsだけでなく自分でラップ識別子を指定できるようにする。zapでくるんでるかとか
