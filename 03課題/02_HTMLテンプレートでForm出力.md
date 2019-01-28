HTMLテンプレート追加
POSTすると先程のhelloに遷移するシンプルなフォーム

```
<!DOCTYPE html>
<html>
<body>
    <form action="/hello" method="get">
    <input type="text" name="input_1">
    <input type="submit" value="送信">
    </form>
</body>
</html>
```

プログラム追記

```
func handleForm(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("./form.html.tpl"))
	// テンプレートを描画
	s := "Hello, world."
	if err := t.ExecuteTemplate(w, "form.html.tpl", s); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/form", handleForm)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
```