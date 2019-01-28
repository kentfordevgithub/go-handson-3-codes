
テンプレートファイルを作成

```
<!DOCTYPE html>
<html>
<body>
    msg: {{.}} 
</body>
</html>
```

プログラム

```
package main

import (
	"html/template"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("./hello.html.tpl"))
	// テンプレートを描画
	s := "Hello, world."
	if err := t.ExecuteTemplate(w, "hello.html.tpl", s); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", handleHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
```

サーバ実行

```
$ cd server
$ go run server.go
```

アクセス先：http://localhost:8080/hello

