package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ???
// 在后端使用的时候还是用int64
// 进行json序列化与反序列化的时候就使用字符串

type Data struct {
	ID int64 `json:"id"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/data", func(c *gin.Context) {
		// int64
		// math.MaxInt64  --> 1<<63 -1
		// math.MinInt64  --> -1 << 63
		//						       1234567891234568000
		c.JSON(http.StatusOK, Data{1234567891234567889})
	})

	r.Run(":8990")
}
