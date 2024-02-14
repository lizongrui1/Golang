package module

import (
	"html/template"
	"os"
	"path/filepath"
	"testing"
)

func TestInitDB(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("无法获取当前工作目录: %v", err)
	}
	tplPath := filepath.Join(cwd, "templates", "*.html")

	_, err = template.ParseGlob(tplPath)
	if err != nil {
		t.Fatalf("解析模板时出错: %v", err)
	}

	db, err := InitDB()
	if err != nil {
		t.Fatalf("初始化数据库时出错: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("无法连接到数据库: %v", err)
	}
}
