package module

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var tpl *template.Template

func QueryRowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// 解析请求参数
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Invalid ID value", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		score, err := strconv.Atoi(r.FormValue("score"))
		if err != nil {
			http.Error(w, "Invalid score value", http.StatusBadRequest)
			return
		}
		insertRow(name, score)
		//http.Redirect(w, r, "/", http.StatusSeeOther)
		err = queryMultiRow()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error inserting row: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Student added successfully with ID %d", id)
	}
}

// 添加学生信息的Handler
func insertRowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		name := r.FormValue("name")
		score, err := strconv.Atoi(r.FormValue("score"))
		if err != nil {
			http.Error(w, "Invalid score value", http.StatusBadRequest)
		}
		insertRow(name, score)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		tpl.ExecuteTemplate(w, "add.html", nil)
	}
}
