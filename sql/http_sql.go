/*package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func check(account, password string) bool {
	sqlStr := `select * from login where account = ? and password =?`
	rows, err := db.Query(sqlStr, account, password)
	if err != nil {
		return false
	}
	defer rows.Close()
	if rows.Next() {
		return true
	}
	return false
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
<html>
<head>
<meta charset="utf-8">
<title>登入系统</title>
</head>
<body>

<h1>登入</h1>
<p>
	<form name="input" method="post">
	<table>
		<tr>
			<td>账号：</td>
			<td><input type="text" name="account"></td>
		</tr>
		<tr>
			<td>密码：</td>
			<td><input type="password" name="password"></td>
		</tr>
	</table>
		<input type="submit" value="Submit">
	</form>
</p> 
</body>
</html>`)

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	account := r.Form.Get("account")
	password := r.Form.Get("password")
	flag := check(account, password)
	if flag == true {
		fmt.Fprintln(w, "yes")
	} else {
		fmt.Fprintln(w, "no")
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	http.HandleFunc("/", Login)
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("err", err)
		return
	}
}
*/