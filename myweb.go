package main

//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//type myForm struct {
//	Colors []string `form:"colors[]"`
//}
//
//func main() {
//	r := gin.Default()
//
//	r.LoadHTMLGlob("templates/**/*")
//	r.GET("users/pokerlive", indexHandler)
//	r.GET("users/index2", indexHandler)
//	r.GET("invite", Invite)
//
//	r.Run(":8080")
//}
//
//func indexHandler(c *gin.Context) {
//	/*
//		获取参数列表
//	*/
//	req := c.Request.URL.Query()
//	fmt.Printf("indexHandler req:%+v\n", req)
//
//	/*
//		构造参数
//	*/
//	var obj = gin.H{}
//	for k, v := range req {
//		obj[k] = v[0]
//	}
//
//	/*
//		构造html
//	*/
//	fmt.Printf("indexHandler path:%+v\n", c.FullPath())
//	//pathList := strings.Split(c.FullPath(), "/")
//	if len(c.FullPath()) == 1 {
//		c.String(http.StatusOK, "接口路径不符合规则", nil)
//		return
//	}
//
//	var htmlFile = c.FullPath()[1:len(c.FullPath())]
//	htmlFile += ".html"
//
//	/*
//		返回对应html
//	*/
//	c.HTML(http.StatusOK, htmlFile, obj)
//	fmt.Printf("path:[%s] params:[%+v] file:[%s] obj:[%+v]\n", c.FullPath(), req, htmlFile, obj)
//}
//
//func Invite(c *gin.Context) {
//	/*
//		获取参数列表
//	*/
//	req := c.Request.URL.Query()
//	fmt.Printf("indexHandler req:%+v\n", req)
//
//
//	fpid := req.Get("fpid")
//	msg := req.Get("msg")
//
//	fmt.Printf("fpid:%d msg:%s", fpid, msg)
//}
