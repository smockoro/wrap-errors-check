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
	r := regexp.MustCompile(`[A-Za-z0-9\_\.\/]*.go`)
	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		if !file.IsDir() && r.MatchString(path) {
			paths = append(paths, path)
			checkWrapped(path)
		} else {
			paths = append(paths, targetWalk(filepath.Join(dir, file.Name()))...)
		}
	}

	return paths
}

func checkWrapped(filepath string) {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, filepath, nil, parser.Mode(0))

	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.FuncDecl); ok {
			fmt.Printf("Function Name [ %s ]\n", v.Name.Name)
			ff := fset.File(v.Pos())
			fmt.Println(fset.File(v.Pos()))
			fmt.Println(ff.Name())        // ファイル名が取れる
			fmt.Println(ff.Line(v.Pos())) // これで行数が取れる
			if v.Type.Results != nil {
				for _, e := range v.Type.Results.List {
					switch rtype := e.Type.(type) {
					case *ast.Ident:
						if rtype.Name == "error" {
							returnCheck(v.Body)
						}
					}
				}
			}
		}
		return true
	})

	// For Debug
	//for _, d := range f.Decls {
	//ast.Print(fset, d)
	//}
}

func returnCheck(body *ast.BlockStmt) {
	ast.Inspect(body, func(n ast.Node) bool {
		if v, ok := n.(*ast.ReturnStmt); ok {
			for _, rlt := range v.Results {
				switch r := rlt.(type) {
				case *ast.Ident:
					fmt.Println(r.Name)
					if r.Name == "err" {
						// ファイル名、行数、関数名とか取れるとうれしい
						fmt.Println("err is not wrapped errors package")

					}
				}
			}
		}
		return true
	})
}

// ModuleCheck : a
func ModuleCheck(m ast.Expr) {

}

// FunctionCheck : a
func FunctionCheck(f *ast.Ident) {

}

//TODO
// テストコードでフォルダウォークするようにする ioutli系を使ったテストをする？
// 上記はどこかに参考資料があったので探して内容を理解してコードに残す。

// 最後の結果として何個のうち何個が通っているのか、
// 間違っている場所はどのソースの何行目何列目なのか出力レポート出力

// errorsだけでなく自分でラップ識別子を指定できるようにする。zapでくるんでるかとか
