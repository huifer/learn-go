# GO 使用 mysql
> 作者: [HuiFer](https://github.com/huifer)
>
> 仓库地址: https://github.com/huifer/learn-go

## 依赖,及安装
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- 使用下列命令进行安装
```shell script
go get -u github.com/go-sql-driver/mysql
```


## 使用

### 创建连接对象
- 一个数据库链接需要具备如下几个要素
    1. 数据库ip
    1. 数据库端口
    1. 数据库账号
    1. 数据库密码
    1. 数据库名称
- 在知道上述条件后根据 `user_name:user_password@tcp(db_ip:db_port)/db_name` 替换其中的关键字即可. 
- 根据上述要素我们可以创建一个结构体

```go
type mysql_conf struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}
```


- 创建数据库链接
```go
func conn(conf mysql_conf) *sql.DB {
	url := conf.Username + ":" + conf.Password + "@tcp(" + conf.Host + ":" + strconv.Itoa(conf.Port) + ")/" + conf.DbName + "?charset=utf8mb4&multiStatements=true"
	fmt.Println(url)
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println("[ERROR] Could not connect to database: ", err)

	}
	return db
}
```

### 插入数据

```go
	sql := "INSERT INTO `scrum`.`t_user`(`name`, `age`) VALUES (?, ?)"
	id, err := insert(conn(conf), sql)

func insert(db *sql.DB, sql string) (int64, error) {
	stmt, err := db.Prepare(sql)
	checkErr(err)
	exec, err := stmt.Exec("张三", 18)
	checkErr(err)
    // 最后插入的id
	return exec.LastInsertId()
}
```


### 查询数据
- 查询后的结果需要放到一个结构体中进行后续的使用,因此我们需要先定义一个`t_user`的结构体为后续进行使用

```go

	sql = "select id , name ,age from t_user where id = ?"

	user := selectSql(conn(conf), sql, id)

type t_user struct {
	Id   int
	Name string
	Age  int
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
```


### 修改数据


```go
	sql := "UPDATE `scrum`.`t_user` SET `age` = ? WHERE `id` = ?"
	upId, err := updateSql(conn(conf), sql, 3, 12)

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
```


### 表对象转换结构体
- 在查询的时候我们定义过一个结构体`t_user`这是我们手动创建的. 在实际场景中我们可能有很多个需要创建,一个个手写有点麻烦因此开发一个转换工具十分有必要


```go
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
		User:   "user",
		Pwd:    "pwd",
		DbName: "db",
		Host:   "127.0.0.1",
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

```



