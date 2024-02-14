package module

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestUpdateRowHandler(t *testing.T) {
	// 创建一个表单数据
	formData := url.Values{}
	formData.Set("name", "Test Student")
	formData.Set("score", "90")

	// 创建一个模拟的请求
	req, err := http.NewRequest("POST", "/update", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// 创建一个 ResponseRecorder (模拟的 ResponseWriter) 来记录响应
	rr := httptest.NewRecorder()

	// 实例化一个具体的 http.Handler，并调用它的 ServeHTTP 方法来直接使用我们的模拟请求和 ResponseRecorder
	handler := http.HandlerFunc(UpdateRowHandler)
	handler.ServeHTTP(rr, req)

	// 检查响应状态码
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}
}

func TestLoginHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoginHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestDeleteRowHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteRowHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestHomeHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HomeHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestInsertRowHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertRowHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestLoginHandler1(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoginHandler(tt.args.w, tt.args.r)
		})
	}
}

//func TestQueryRowHandler(t *testing.T) {
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			QueryRowHandler(tt.args.w, tt.args.r)
//		})
//	}
//}

func TestUpdateRowHandler1(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateRowHandler(tt.args.w, tt.args.r)
		})
	}
}

//func TestQueryRowHandler(t *testing.T) {
//	form := url.Values{}
//	form.Add("1999", "李四")
//	req, err := http.NewRequest("POST", "/query", strings.NewReader(form.Encode()))
//	if err != nil {
//		t.Fatal(err)
//	}
//	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//	// 创建一个 ResponseWriter 来接收 HTTP 响应
//	rr := httptest.NewRecorder()
//
//	// 处理 HTTP 请求
//	handler := http.HandlerFunc(QueryRowHandler)
//	handler.ServeHTTP(rr, req)
//
//	// 检查响应状态码是否为http.StatusOK，因为这是一个成功的查询响应
//	// 注意：实际测试中，你需要根据 queryRow 函数的实现和测试环境调整期望的状态码和响应体
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
//	}
//
//	// 检查响应体的具体内容可能会比较困难，因为它依赖于模板渲染的结果
//	// 可能需要检查响应体是否包含某些关键字，如学生的名字或ID
//	// expectedBody := "expected content based on mock data"
//	// if !strings.Contains(rr.Body.String(), expectedBody) {
//	//     t.Errorf("handler returned unexpected body: got %v want to include %v", rr.Body.String(), expectedBody)
//	// }
//}

func TestQueryRowHandler(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("无法获取当前工作目录: %v", err)
	}
	fmt.Println("当前工作目录:", wd)
	db, _ := InitDB()

	defer db.Close()
	testStudentNumber := "1999"
	formData := url.Values{}
	formData.Set("id", testStudentNumber)
	req, err := http.NewRequest("POST", "/query", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(QueryRowHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expectedName := "李四"
	expectedScore := "100"
	body := rr.Body.String()
	if !strings.Contains(body, expectedName) || !strings.Contains(body, expectedScore) {
		t.Errorf("获得： %v ，想要包含姓名： %v 和分数： %v\n", body, expectedName, expectedScore)
	}
}
//模板解析错误：open ./module/templates/querySuccess.html: no such file or directory???