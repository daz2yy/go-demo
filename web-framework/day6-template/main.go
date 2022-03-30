package main

// 前端模版
// 1. 支持返回静态文件
// 2. 支持替换模版文件的变量（渲染）

import (
	"fmt"
	"html/template"
	"mango"
	"net/http"
	"time"
)

type Student struct {
	Name string
	Age  int
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	engine := mango.New()
	engine.Use(mango.Logger())

	engine.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/assets", "./static")

	// demo
	stu1 := &Student{Name: "Jack", Age: 18}
	stu2 := &Student{Name: "LiLei", Age: 21}
	engine.GET("/", func(c *mango.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	engine.GET("/student", func(c *mango.Context) {
		c.HTML(http.StatusOK, "student.tmpl", mango.H{
			"title":  "Mango",
			"stuArr": [2]*Student{stu1, stu2},
		})
	})
	engine.GET("/date", func(c *mango.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", mango.H{
			"title": "Mango",
			"now":   time.Date(2022, 3, 30, 0, 0, 0, 0, time.UTC),
		})
	})

	engine.Run(":8899")
}
