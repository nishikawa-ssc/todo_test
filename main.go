package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nishikawa-ssc/todo_test/controller"
	task "github.com/nishikawa-ssc/todo_test/db"
)

func main() {
	// データアクセス初期化
	task.Init()
	defer task.Close()

	// ルータ設定
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	h := controller.HdlrTask{Db: task.DaTask()}
	r.GET("/", h.GetAll)
	r.POST("/", h.Create)
	r.GET("/:id", h.Edit)
	r.POST("/update/:id", h.Update)
	r.POST("/delete/:id", h.Delete)

	r.Run()
}
