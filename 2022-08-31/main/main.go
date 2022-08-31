package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CheckboxForm struct {
	// checked or unchecked > true or false
	CheckboxStatus string `form:"ex-checkbox"`
}

func main() {

	g := gin.Default()

	g.POST("/p1", func(c *gin.Context) {

		var form CheckboxForm
		err := c.ShouldBind(&form)
		if err != nil {
			panic(err)
		}
		fmt.Println(form)

	})

	g.Run()
}
