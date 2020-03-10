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
	r.LoadHTMLGlob("task/templates/*.html")

	// ハンドラ設定
	h := controller.HdlrTask{Db: d.Task()}
	g := r.Group("/task")
	{
		g.GET("", h.GetAll)
		g.POST("", h.Create)
		g.GET("/:id", h.Edit)
		g.POST("/update/:id", h.Update)
		g.POST("/delete/:id", h.Delete)
	}

	r.Run()
}
