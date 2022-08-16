package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.RedirectTrailingSlash = false
	r.GET("/test/name", func(gCtx *gin.Context) {
		gCtx.JSON(http.StatusOK, gin.H{
			"label":  "/test/name",
			"name":   gCtx.Param("name"),
			"active": gCtx.Param("active"),
		})
	})
	r.GET("/test/:name", func(gCtx *gin.Context) {

		gCtx.JSON(http.StatusOK, gin.H{
			"label":  "/test/:name",
			"name":   gCtx.Param("name"),
			"active": gCtx.Param("active"),
		})
	})
	r.GET("/test/:name/", func(gCtx *gin.Context) {
		gCtx.JSON(http.StatusOK, gin.H{
			"label":  "/test/:name/",
			"name":   gCtx.Param("name"),
			"active": gCtx.Param("active"),
		})
	})
	if err := r.Run(); nil != err {
		log.Fatal(err)
	}
}
