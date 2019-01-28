プログラム作る

```
package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type CalcResult struct {
	Input  int `json:"input"`
	Output int `json:"output"`
}

func main() {
	e := echo.New()

	// Routing
	e.GET("/twice/:input", func(c echo.Context) error {
		input, err := strconv.Atoi(c.Param("input"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "input error.")
		}
		return c.JSON(http.StatusOK, twice(input))
	})

	e.Logger.Fatal(e.Start(":9090"))
}

func twice(in int) *CalcResult {
	out := in * 2
	r := &CalcResult{
		Input:  in,
		Output: out,
	}
	return r
}
```

freshを通して実行

```
# カレントディレクトリを実行したいパッケージの配下へ移動
$ cd ./api
# freshを実行
$ $HOME/go/bin/fresh
```
