package module

import (
	"fmt"
	"net/http"
	"strconv"
)

// 获取所有学生信息的Handler
func QueryRowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// 解析请求参数
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "解析表单错误", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "无效的ID值", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		score, err := strconv.Atoi(r.FormValue("score"))
		if err != nil {
			http.Error(w, "无效的分数值", http.StatusBadRequest)
			return
		}
		insertRow(name, score)
		// http.Redirect(w, r, "/", http.StatusSeeOther)
		err = queryRow()
		if err != nil {
			http.Error(w, fmt.Sprintf("插入行时出错: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "成功添加学生，ID为 %d", id)
	}

	tpl.ExecuteTemplate(w, "index.html", student)
}

// 添加学生信息的Handler
func InsertRowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		name := r.FormValue("name")
		score, err := strconv.Atoi(r.FormValue("score"))
		if err != nil {
			http.Error(w, "无效的分数值", http.StatusBadRequest)
		}
		insertRow(name, score)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		tpl.ExecuteTemplate(w, "add.html", nil)
	}
}

// 修改学生信息的Handler
func UpdateRowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { //POST请求
		err := r.ParseForm()
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "无效的学生ID值", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		score, err := strconv.Atoi(r.FormValue("score"))
		if err != nil {
			http.Error(w, "无效的分数值", http.StatusBadRequest)
			return
		}
		updateRow(id, name, score)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else { //GET请求
		/*studentID := r.FormValue("id")
		if studentID == "" {
			http.Error(w, "未提供学生ID", http.StatusBadRequest)
			return
		}*/
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
		tpl.ExecuteTemplate(w, "edit.html", nil)
	}
}

// 删除学生信息的Handler
func DeleteRowHandler(w http.ResponseWriter, r *http.Request) {
	// 从URL查询参数中获取id
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "学生ID未提供", http.StatusBadRequest)
		return
	}

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
