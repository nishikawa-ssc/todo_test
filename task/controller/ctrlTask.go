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

// Get １レコード取得
func Get(id int, db *gorm.DB) m.Task {
	var dat m.Task
	db.First(&dat, id)

	return dat
}

// GetAll 一覧表示
func (h *HdlrTask) GetAll(c *gin.Context) {
	// 全データ取得
	var dat []m.Task
	h.Db.Order("created_at desc").Find(&dat)
	// 表示
	c.HTML(http.StatusOK, "index.html", gin.H{"dat": dat})
}

// Create 新規追加
func (h *HdlrTask) Create(c *gin.Context) {
	// 入力情報収集
	t := c.PostForm("text")
	status := c.PostForm("status")
	s, err := strconv.Atoi(status)
	if err != nil {
		panic(err)
	}

	// DB登録
	h.Db.Create(&m.Task{Text: t, Stat: s})
	// 初期表示（一覧表示）をリダイレクト
	c.Redirect(http.StatusMovedPermanently, "/task")
}

// Edit 編集
func (h *HdlrTask) Edit(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	dat := Get(id, h.Db)
	c.HTML(http.StatusOK, "edit.html", gin.H{"dat": dat})
}

// Update 更新
func (h *HdlrTask) Update(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	t := c.PostForm("text")
	status := c.PostForm("status")
	s, err := strconv.Atoi(status)
	if err != nil {
		panic(err)
	}

	// 対象IDのレコード取得
	dat := Get(id, h.Db)
	// 値の更新
	dat.Text = t
	dat.Stat = s
	// 保存
	h.Db.Save(&dat)

	c.Redirect(http.StatusMovedPermanently, "/task")
}

// Confirm 削除確認
func (h *HdlrTask) Confirm(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	dat := Get(id, h.Db)
	c.HTML(http.StatusOK, "confirm.html", gin.H{"dat": dat})
}

// Delete 削除
func (h *HdlrTask) Delete(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	// 対象のレコード取得
	dat := Get(id, h.Db)
	// 削除
	h.Db.Delete(&dat)
	c.Redirect(http.StatusMovedPermanently, "/task")
}
