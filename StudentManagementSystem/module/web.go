package module

import (
	"fmt"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// 确保只有根路径 "/" 被这个处理器处理
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// 加载并发送 choose.html 文件
	http.ServeFile(w, r, "./module/templates/choose.html")
}

// 获取所有学生信息的Handler
func QueryRowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "解析表格失败", http.StatusBadRequest)
			return
		}
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "无效的学生ID", http.StatusBadRequest)
			return
		}
		// 查询学生信息
		stu, err := queryRow(id)
		if err != nil {
			http.Error(w, "查询失败", http.StatusInternalServerError)
			return
		}
		//在模板中使用查询结果
		err = tpl.ExecuteTemplate(w, "query.html", stu)
		if err != nil {
			return
		}
	} else {
		http.Error(w, "无效请求", http.StatusMethodNotAllowed)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// 添加学生信息的Handler
func InsertRowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "解析表格失败", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		score, err := strconv.Atoi(r.FormValue("score"))
		if err != nil {
			http.Error(w, "无效的分数值", http.StatusBadRequest)
		}
		err = insertRow(name, score)
		if err != nil {
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		err := tpl.ExecuteTemplate(w, "add.html", nil)
		if err != nil {
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// 修改学生信息的Handler
func UpdateRowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//POST请求（写）
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "解析表格失败", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		score, err := strconv.Atoi(r.FormValue("score"))
		if err != nil {
			http.Error(w, "无效的分数值", http.StatusBadRequest)
			return
		}
		err = updateRow(name, score)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} /*else {
		GET请求（读）
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "无效的学生ID值", http.StatusBadRequest)
			return
		}
		err = queryRow(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tpl.ExecuteTemplate(w, "edit.html", nil)
		if err != nil {
			return
		}
	}*/
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// 删除学生信息的Handler
func DeleteRowHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "表单解析错误", http.StatusBadRequest)
		return
	}
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "无效的学生ID", http.StatusBadRequest)
		return
	}
	// 调用deleteRow来删除学生
	if err := deleteRow(id); err != nil {
		// 如果删除过程中出现错误，返回内部服务器错误
		http.Error(w, fmt.Sprintf("删除失败: %v", err), http.StatusInternalServerError)
		return
	}
	// 删除成功后重定向到根路径
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
