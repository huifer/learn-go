package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type column struct {
	ColumnName string
	Type       string
	Nullable   string
}
type mysqlConfig struct {
	User   string
	Pwd    string
	DbName string
	Host   string
	Port   int
}

// mysql 和 go 语言的类型对应表
var go_mysql_typemap = map[string]string{
	"int":                "int",
	"integer":            "int",
	"tinyint":            "int",
	"smallint":           "int",
	"mediumint":          "int",
	"bigint":             "int",
	"int unsigned":       "int",
	"integer unsigned":   "int",
	"tinyint unsigned":   "int",
	"smallint unsigned":  "int",
	"mediumint unsigned": "int",
	"bigint unsigned":    "int",
	"bit":                "int",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time",
	"datetime":           "time.Time",
	"timestamp":          "time.Time",
	"time":               "time.Time",
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

var query_sql = `SELECT COLUMN_NAME,DATA_TYPE,IS_NULLABLE
		FROM information_schema.COLUMNS 
		WHERE table_schema = DATABASE()`

func main() {
	fmt.Println("测试输出")
	config := mysqlConfig{
		User:   "shandsmod3",
		Pwd:    "shandsmod3",
		DbName: "shandsmod3",
		Host:   "10.10.0.124",
		Port:   3306,
	}
	conn := config.User + ":" + config.Pwd + "@tcp(10.10.0.124:3306)/" + config.DbName
	fmt.Println(conn)
	if conn != "" {
		db, err := sql.Open("mysql", conn)
		if err != nil {
			fmt.Println("[ERROR] Could not connect to database: ", err)
		}
		err, columns := getColumns("hs_log", db)
		model("hs_log", columns)
	}

}

// 获取所有列
func getColumns(tablename string, db *sql.DB) (errr error, columns []column) {
	if tablename != "" {
		sql := query_sql + "and TABLE_NAME = '" + tablename + "'"

		rows, err := db.Query(sql)
		if err != nil {
			fmt.Println("执行SQL错误", err)
		}
		defer rows.Close()

		for rows.Next() {
			col := column{}

			err := rows.Scan(&col.ColumnName, &col.Type, &col.Nullable)
			if err != nil {
				fmt.Println("sql结果转换对象失败")
				return err, nil
			}
			col.Type = go_mysql_typemap[col.Type]

			columns = append(columns, col)
		}
	}
	return errr, columns
}

// 对象输出
func model(table_name string, columns []column) {
	depth := 1
	fmt.Print("type " + table_name + " struct {\n")
	for _, v := range columns {
		fmt.Print(tab(depth) + v.ColumnName + " " + v.Type)
		fmt.Print("\n")
	}

	fmt.Print(tab(depth-1) + "}\n")

}

// 输出 table
func tab(depth int) string {
	return strings.Repeat("\t", depth)
}
