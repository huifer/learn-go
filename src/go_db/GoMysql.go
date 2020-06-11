package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

// mysql 配置
type mysql_conf struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

func conn(conf mysql_conf) *sql.DB {
	url := conf.Username + ":" + conf.Password + "@tcp(" + conf.Host + ":" + strconv.Itoa(conf.Port) + ")/" + conf.DbName + "?charset=utf8mb4&multiStatements=true"
	fmt.Println(url)
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println("[ERROR] Could not connect to database: ", err)

	}
	return db
}

func insert(db *sql.DB, sql string) (int64, error) {
	stmt, err := db.Prepare(sql)
	checkErr(err)
	exec, err := stmt.Exec("张三", 18)
	checkErr(err)
	return exec.LastInsertId()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type t_user struct {
	Id   int
	Name string
	Age  int
}

func main() {
	conf := mysql_conf{
		Host:     "47.98.225.144",
		Port:     3306,
		Username: "huifer",
		Password: "a12345",
		DbName:   "scrum",
	}

	//sql := "INSERT INTO `scrum`.`t_user`(`name`, `age`) VALUES (?, ?)"
	//id, err := insert(conn(conf), sql)
	//checkErr(err)
	//sql = "select id , name ,age from t_user where id = ?"
	//
	//user := selectSql(conn(conf), sql, id)
	//fmt.Println(user)
	sql := "UPDATE `scrum`.`t_user` SET `age` = ? WHERE `id` = ?"
	upId, err := updateSql(conn(conf), sql, 3, 12)
	checkErr(err)

	fmt.Println("id = ", upId)
}

func updateSql(db *sql.DB, sql string, id int64, age int) (int64, error) {
	prepare, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	exec, err := prepare.Exec(age, id)
	if err != nil {
		panic(err)
	}
	return exec.RowsAffected()
}

func selectSql(db *sql.DB, s string, id int64) t_user {
	stmt, err := db.Prepare(s)
	checkErr(err)
	exec, err := stmt.Query(id)
	checkErr(err)
	user := t_user{}

	for exec.Next() {
		exec.Scan(&user.Id, &user.Name, &user.Age)
	}
	return user

}
