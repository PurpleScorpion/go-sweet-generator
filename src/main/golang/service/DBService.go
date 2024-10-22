package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"sweet-common/utils"
	"sweet-src/main/golang/model"
)

func connMySQL() *sql.DB {
	username := utils.ValueString("${sweet.db.username}")
	dbName := utils.ValueString("${sweet.db.dbName}")
	host := utils.ValueString("${sweet.db.host}")
	port := utils.ValueInt("${sweet.db.port}")
	password := utils.ValueString("${sweet.db.password}")

	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbName)
	db, err := sql.Open("mysql", str)
	if err != nil {
		panic(err)
	}
	return db
}

func getRowByTable(db *sql.DB, tableName string) model.Tables {
	var table model.Tables
	rows, err := db.Query("DESC " + tableName)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	table.Name = tableName

	for rows.Next() {
		var tf model.TableFiled
		var field, dataType, null, key, extra string
		var defaultValue interface{}
		err := rows.Scan(&field, &dataType, &null, &key, &defaultValue, &extra)
		if err != nil {
			log.Fatal(err)
		}
		tf.Name = field
		tf.Type = dataType
		tf.Key = false
		if key == "PRI" {
			tf.Key = true
		}

		table.Fileds = append(table.Fileds, tf)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return table
}

func getFiledType(ft string) string {
	arr := utils.ValueStringArr("${sweet.filedType}")

	for _, v := range arr {
		ss := strings.Split(v, "=")
		if strings.HasPrefix(ft, ss[0]) {
			return ss[1]
		}
	}
	panic("没有找到对应的类型: " + ft)
}
