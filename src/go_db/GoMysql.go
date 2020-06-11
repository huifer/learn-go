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
	url := conf.Username + ":" + conf.Password + "@tcp(" + conf.Host + ":" + strconv.Itoa(conf.Port) + ")/" + conf.DbName
	fmt.Println(url)
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println("[ERROR] Could not connect to database: ", err)

	}
	return db
}

func main() {
	conf := mysql_conf{
		Host:     "10.10.0.124",
		Port:     3306,
		Username: "shandsmod3",
		Password: "shandsmod3",
		DbName:   "shandsmod3",
	}

	db := conn(conf)
	query, err := db.Query("select * from hs_log")
	if err != nil {
		for query.Next() {
			columns, err := query.Columns()
			if err == nil {
				print(columns)
			}
		}
	}
}
