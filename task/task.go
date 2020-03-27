package task

import (
	"github.com/gin-gonic/gin"
	"github.com/nishikawa-ssc/todo_test/task/constant"
	"github.com/nishikawa-ssc/todo_test/task/controller"
	"github.com/nishikawa-ssc/todo_test/task/db"
)

// Run 実行
func Run() {
	// データアクセス初期化
	db.Init()
	defer db.Close()
	// ルータ設定
	r := gin.Default()

	// ハンドラ設定
	h := controller.HdlrTask{Db: db.Data()}
	g := r.Group("/task")
	{
		g.Static("/style", "task/style")
		g.Static("/js", "task/js")
		g.StaticFile("/", constant.BaseHTML)

		g.GET("/fetchAll", h.FetchAll)
		g.GET("/fetchOne", h.FetchOne)
		g.POST("/add", h.Create)
		g.POST("/update/", h.Update)
		g.POST("/delete", h.Delete)
	}

	r.Run(constant.Port)
}
