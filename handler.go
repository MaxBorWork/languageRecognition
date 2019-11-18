package main

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func NgrammMethod(c *gin.Context) {

}

func AlphabetMethod(c *gin.Context)  {

}
