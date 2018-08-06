// 需要先安装: go get github.com/go-sql-driver/mysql

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "root"
	password = ""
	ip = "localhost"
	port = "3306"
	dbName = "test"
)

func main() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	//DB, _ = sql.Open("mysql", path)
	db, err := sql.Open("mysql", path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	var result sql.Result
	result, err = db.Exec("insert into t_test(age, name) values(?,?)", 13, "joe")
	if err != nil {
		fmt.Println(err)
		return
	}

	lastId, _ := result.LastInsertId()
	fmt.Println("新插入记录的ID为", lastId)

	var row *sql.Row
	row = db.QueryRow("select * from t_test")
	var name string
	var id, age int
	err = row.Scan(&id, &age, &name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(id, "\t", name, "\t", age)

	result, err = db.Exec("insert into t_test(age, name) values(?,?)", 24, "black")

	var rows *sql.Rows
	rows, err = db.Query("select * from t_test")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var name string
		var id, age int
		rows.Scan(&id, &age, &name)
		fmt.Println(id, "\t", name, "\t", age)
	}
	rows.Close()

	db.Exec("truncate table t_test")
}
