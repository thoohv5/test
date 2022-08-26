package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Params struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
}

func main() {

	g := gin.Default()
	// 开启JSON绑定支持params和query
	gin.EnableParamsAndQueryBinding()

	g.GET("/test/:b", func(gCtx *gin.Context) {

		p := new(Params)

		gCtx.ShouldBindJSON(p)

		gCtx.JSON(http.StatusOK, p)
	})

	g.Run()

	/**
	GET http://127.0.0.1:8080/test/item?a=99
	Accept: application/json

	{"c": "4"}
	**/
	/**
	{
		"a": "99",
		"b": "item",
		"c": "4"
	}
	*/

}
