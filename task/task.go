import (
	"github.com/gin-gonic/gin"
	"github.com/nishikawa-ssc/todo_test/task/controller"
	"github.com/nishikawa-ssc/todo_test/task/db"
)

func run() {
	// データアクセス初期化
	db.Init()
	defer db.Close()

	// ルータ設定
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	h := controller.HdlrTask{Db: db.Task()}
	r.GET("/", h.GetAll)
	r.POST("/", h.Create)
	r.GET("/:id", h.Edit)
	r.POST("/update/:id", h.Update)
	r.POST("/delete/:id", h.Delete)

	r.Run()
}
