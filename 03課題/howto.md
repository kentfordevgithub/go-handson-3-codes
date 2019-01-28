https://qiita.com/spiegel-im-spiegel/items/5cb1587cb55d6f6a34d7


ルートにソース作る

```
package main

import (
    "fmt"

    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
}
```

go mod initする

```
$ go mod init hello
```

こんなgo.modファイルがルートに出来上がる

```
module hello
```

プログラムを実行できる

```
$ go run hello.go
```

依存パッケージ(rsc.io)があるのでrunした瞬間に対象パッケージが取得される
結果、go.modが下記のように追記される

```
module hello

require rsc.io/quote v1.5.2 // indirect
```

依存関係を確認するコマンド

```
$ go mod graph
```

一旦ここでVSCodeを再起動、そうするとVSCode上のパッケージが無いエラーもなくなる

