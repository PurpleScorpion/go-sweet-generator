package service

import (
	"os"
	"strings"
	"sweet-common/utils"
	"sweet-src/main/golang/model"
)

var line = "\n"

func Generator() {
	db := connMySQL()
	str := utils.ValueStringArr("${sweet.tableName}")
	var arr []model.Tables
	for i := 0; i < len(str); i++ {
		tb := getRowByTable(db, str[i])
		saveFile(tb)
		arr = append(arr, tb)
	}
	generatorRouters(arr)

	defer db.Close()
}

func saveFile(tb model.Tables) {
	basePath := "generatorFile/" + utils.ValueString("${sweet.mainPath}")
	generatorModel(basePath, tb)
	generatorVO(tb)
	generatorController(basePath, tb)
	generatorService(basePath, tb)
}

func generatorVO(tb model.Tables) {
	basePath := "generatorFile/common/vo"
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		panic(err)
	}

	str := `package vo`
	str += getModelImports(tb)
	str += "type " + getGoModelName(tb.Name) + "VO" + " struct{"
	str += line
	str += "    DefaultPageVO"
	str += line
	for _, filed := range tb.Fileds {
		str += "    " + utils.SnakeToPascal(filed.Name) + " " + getFiledType(filed.Type) + " `json:\"" + utils.SnakeToLowerCamel(filed.Name) + "\"`"
		str += line
	}

	str += `}`

	filePath := basePath + "/" + utils.SnakeToPascal(tb.Name) + "VO" + ".go"
	os.WriteFile(filePath, []byte(str), 0644)
}

func generatorModel(basePath string, tb model.Tables) {
	basePath = basePath + "/models"
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		panic(err)
	}

	str := `package models`
	str += getModelImports(tb)
	str += "type " + getGoModelName(tb.Name) + " struct{"
	str += line
	for _, filed := range tb.Fileds {
		if filed.Key {
			str += "    " + utils.SnakeToPascal(filed.Name) + " " + getFiledType(filed.Type) + " `json:\"" + utils.SnakeToLowerCamel(filed.Name) + "\" tableId:\"" + filed.Name + "\"`"
		} else {
			str += "    " + utils.SnakeToPascal(filed.Name) + " " + getFiledType(filed.Type) + " `json:\"" + utils.SnakeToLowerCamel(filed.Name) + "\"`"
		}
		str += line
	}

	str += `}`
	str += line
	str += line
	str += line
	str += line
	str += `func (` + utils.SnakeToPascal(tb.Name) + `) TableName() string {`
	str += line
	str += `    return "` + tb.Name + `"`
	str += line
	str += `}`

	filePath := basePath + "/" + utils.SnakeToPascal(tb.Name) + ".go"
	os.WriteFile(filePath, []byte(str), 0644)
}

func getGoModelName(name string) string {
	tablePrefix := utils.ValueString("${sweet.tablePrefix}")
	if utils.IsEmpty(tablePrefix) {
		return utils.SnakeToPascal(name)
	}
	newName := strings.Replace(name, tablePrefix, "", 1)
	return utils.SnakeToPascal(newName)
}

func getModelImports(tb model.Tables) string {
	str := line
	str += line
	str += line
	for _, filed := range tb.Fileds {
		filedType := getFiledType(filed.Type)
		if filedType == "time" {
			str += "import \"time\""
			str += line
			str += line
			str += line
			return str
		}
	}
	return str
}
