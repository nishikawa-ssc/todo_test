package task

import (
	"github.com/gin-gonic/gin"
	"github.com/nishikawa-ssc/todo_test/task/controller"
	d "github.com/nishikawa-ssc/todo_test/task/db"
)

// Run 実行
func Run() {
	// データアクセス初期化
	d.Init()
	defer d.Close()

	// ルータ設定
	r := gin.Default()
	r.Static("/style", "./task/style")
	r.Static("/js", "./task/js")
	r.StaticFile("/task", "./task/templates/task.html")

	// ハンドラ設定
	h := controller.HdlrTask{Db: d.Task()}
	g := r.Group("/taskapp")
	{
		g.GET("/fetchAll", h.FetchAll)
		g.GET("/fetchOne", h.FetchOne)
		g.POST("/add", h.Create)
		g.POST("/edit", h.Edit)
		g.POST("/update/:id", h.Update)
		g.POST("/delete", h.Delete)
	}

	r.Run()
}
