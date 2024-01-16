package module

import (
	"fmt"
	"net/http"
	"strconv"
)

<<<<<<< HEAD
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
	if r.Method != "POST" {
		http.Error(w, "无效请求", http.StatusMethodNotAllowed)
		return
	}
	// 解析请求参数
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "解析表格失败", http.StatusBadRequest)
		return
	}
	// 查询学生信息
	students, err := queryMultiRow()
	if err != nil {
		http.Error(w, "查询失败", http.StatusInternalServerError)
		return
	}
	// 在模板中使用查询结果
	err = tpl.ExecuteTemplate(w, "query.html", students)
	if err != nil {
		return
=======
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
>>>>>>> 6f45ca077e45fe2b265c28107e07681450dc02a6
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
<<<<<<< HEAD
		}
		err = insertRow(name, score)
		if err != nil {
			return
=======
>>>>>>> 6f45ca077e45fe2b265c28107e07681450dc02a6
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		err := tpl.ExecuteTemplate(w, "add.html", nil)
		if err != nil {
			return
		}
	}
}

// 修改学生信息的Handler
func UpdateRowHandler(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	if r.Method == "POST" {
		//POST请求（写）
=======
	if r.Method == "POST" { //POST请求
>>>>>>> 6f45ca077e45fe2b265c28107e07681450dc02a6
		err := r.ParseForm()
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "无效的学生ID值", http.StatusBadRequest)
			return
		}
<<<<<<< HEAD
=======
		name := r.FormValue("name")
>>>>>>> 6f45ca077e45fe2b265c28107e07681450dc02a6
		score, err := strconv.Atoi(r.FormValue("score"))
		if err != nil {
			http.Error(w, "无效的分数值", http.StatusBadRequest)
			return
		}
<<<<<<< HEAD
		err = updateRow(id, score)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} /*else {
		GET请求（读）
=======
		updateRow(id, name, score)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else { //GET请求
		/*studentID := r.FormValue("id")
		if studentID == "" {
			http.Error(w, "未提供学生ID", http.StatusBadRequest)
			return
		}*/
>>>>>>> 6f45ca077e45fe2b265c28107e07681450dc02a6
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
<<<<<<< HEAD
		err = tpl.ExecuteTemplate(w, "edit.html", nil)
		if err != nil {
			return
		}
	}*/
=======
		tpl.ExecuteTemplate(w, "edit.html", nil)
	}
>>>>>>> 6f45ca077e45fe2b265c28107e07681450dc02a6
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
