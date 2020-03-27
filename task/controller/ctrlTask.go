package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "github.com/nishikawa-ssc/todo_test/task/models"
)

// HdlrTask Taskのハンドラ
type HdlrTask struct {
	Db *gorm.DB
}

// FetchAll 全データ取得
func (h *HdlrTask) FetchAll(c *gin.Context) {
	// 全データ取得
	var dat []m.Task
	h.Db.Order("id asc").Find(&dat)

	c.JSON(http.StatusOK, dat)
}

// get １レコード取得
func get(id int, db *gorm.DB) m.Task {
	var dat m.Task
	db.First(&dat, id)

	return dat
}

// FetchOne 1レコード取得
func (h *HdlrTask) FetchOne(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	dat := get(id, h.Db)
	c.JSON(http.StatusOK, dat)
}

// Create 新規追加
func (h *HdlrTask) Create(c *gin.Context) {
	// 入力情報収集
	t := c.PostForm("text")
	stat := c.PostForm("stat")
	s, err := strconv.Atoi(stat)
	if err != nil {
		panic(err)
	}

	// DB登録
	h.Db.Create(&m.Task{Text: t, Stat: s})
}

// Update 更新
func (h *HdlrTask) Update(c *gin.Context) {
	n := c.PostForm("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	t := c.PostForm("text")
	status := c.PostForm("stat")
	s, err := strconv.Atoi(status)
	if err != nil {
		panic(err)
	}

	// 対象IDのレコード取得
	dat := get(id, h.Db)
	// 値の更新
	dat.Text = t
	dat.Stat = s
	// 保存
	h.Db.Save(&dat)
}

// Delete 削除
func (h *HdlrTask) Delete(c *gin.Context) {
	n := c.PostForm("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	// 対象のレコード取得
	dat := get(id, h.Db)
	// 削除
	h.Db.Delete(&dat)
}
