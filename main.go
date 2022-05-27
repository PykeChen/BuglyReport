package main

import (
	"base.bugly/api"
	log "base.bugly/pkg/plog"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {

	log.InitFileLogger("./", "report")

	// 1.创建路由
	r := gin.Default()
	// 2. 绑定路由规则、执行的函数
	group := r.Group("/bugly")
	// gin.Context, 封装了request和response
	group.POST("/report", reportCallback)

	// 3. 监听端口，不指定默认8080
	err := r.Run(":8000")
	if err != nil {
		panic(err)
		return
	}
}



func reportCallback(c *gin.Context) {

	//var data models.MvpFeedbackList
	//err := c.ShouldBindJSON(&data)
	//tools.HasError(err, "", 500)
	//result, err := data.Create()
	//tools.HasError(err, "", -1)
	//app.OK(c, result, "")
	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("Body:\n%+v \n", string(data))
	log.Printf("Body:\n%v", string(data))
	c.String(api.StatusOK, "success")
}
