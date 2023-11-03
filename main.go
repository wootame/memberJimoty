package main

import (
    "github.com/gin-gonic/gin"
    "memberProject/routers" // ルーティングの設定があるパッケージをインポート
)

func main() {
    r := gin.Default()
    routers.SetupRouter(r)

	// ポート番号は必要に応じて変更
    r.Run(":3000")
}
