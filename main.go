package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
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
			path := filepath.Join(dir, file.Name())
			paths = append(paths, path)
			checkGofile(path)
		} else {
			paths = append(paths, targetWalk(filepath.Join(dir, file.Name()))...)
		}
	}

	return paths
}

func checkGofile(filepath string) {
	r := regexp.MustCompile(`[A-Za-z0-9\_\.\/]*.go`)
	if r.MatchString(filepath) {
		fmt.Println(filepath)
		// *.goにマッチすれば構文解析
		checkWrapped(filepath)
	}
}

func checkWrapped(filepath string) {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, filepath, nil, parser.Mode(0))

	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.CallExpr); ok {
			fmt.Println(v.Fun.Pos())
		}
		return true
	})

	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.CallExpr); ok {
			var f *ast.Ident
			var m ast.Expr
			switch fun := v.Fun.(type) {
			case *ast.Ident:
				f = fun
			case *ast.SelectorExpr:
				m, f = fun.X, fun.Sel
			}
			fmt.Println(f, m)
		}
		return true
	})

	for _, d := range f.Decls {
		ast.Print(fset, d)
	}
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
