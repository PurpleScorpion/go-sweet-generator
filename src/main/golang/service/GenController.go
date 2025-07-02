package service

import (
	"fmt"
	"os"
	"sweet-common/utils"
	"sweet-src/main/golang/model"
)

func generatorController(basePath string, tb model.Tables) {
	basePath = basePath + "/controllers"
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		panic(err)
	}

	name := getGoModelName(tb.Name)

	str := `package controllers`
	str += line
	str += line
	str += getControllerImport(tb)
	str += getControllerStruct(tb)
	str += getControllerAutowired(tb)
	str += getControllerPageData(tb)
	str += getControllerGetById(tb)
	str += getControllerSave(tb)
	str += getControllerUpdate(tb)
	str += getControllerDeleteById(tb)
	filePath := basePath + "/" + utils.SnakeToPascal(name) + "Controller.go"
	os.WriteFile(filePath, []byte(str), 0644)
}

func getControllerUpdate(tb model.Tables) string {
	name := getGoModelName(tb.Name)
	str := `
func (that *` + name + `Controller) Update() {
	data := that.Ctx.Input.RequestBody
	var pojo models.` + name + `
	json.Unmarshal(data, &pojo)
	r := ` + utils.Lcfirst(name) + `Service.Update(pojo)
	that.Result(r)
}
`
	str += line
	return str
}

func getControllerSave(tb model.Tables) string {
	name := getGoModelName(tb.Name)
	str := `
func (that *` + name + `Controller) Insert() {
	data := that.Ctx.Input.RequestBody
	var pojo models.` + name + `
	json.Unmarshal(data, &pojo)
	r := ` + utils.Lcfirst(name) + `Service.Insert(pojo)
	that.Result(r)
}
`
	str += line
	return str
}

func getControllerDeleteById(tb model.Tables) string {
	name := getGoModelName(tb.Name)
	str := `
func (that *` + name + `Controller) DeleteById() {
	` + getControllerIdType(tb) + `
	r := ` + utils.Lcfirst(name) + `Service.DeleteById(id)
	that.Result(r)
}
`
	str += line
	return str
}
func getControllerGetById(tb model.Tables) string {
	name := getGoModelName(tb.Name)
	str := `
func (that *` + name + `Controller) GetById() {
	` + getControllerIdType(tb) + `
	r := ` + utils.Lcfirst(name) + `Service.GetById(id)
	that.Result(r)
}
`
	str += line
	return str
}

func getControllerPageData(tb model.Tables) string {
	name := getGoModelName(tb.Name)

	str := `
func (that *` + name + `Controller) PageData() {
	data := that.Ctx.Input.RequestBody
	var pageVO vo.` + name + `VO
	json.Unmarshal(data, &pageVO)
	r := ` + utils.Lcfirst(name) + `Service.PageData(pageVO)
	that.Result(r)
}
`
	str += line
	return str
}

func getControllerAutowired(tb model.Tables) string {
	str := `var ` + utils.Lcfirst(getGoModelName(tb.Name)) + "Service service." + getGoModelName(tb.Name) + "Service"
	str += line
	return str
}

func getControllerStruct(tb model.Tables) string {
	str := `
type ` + getGoModelName(tb.Name) + `Controller struct {
	BaseController
}`
	str += line
	return str
}

func getControllerImport(tb model.Tables) string {
	commonName := utils.ValueString("${sweet.moduleName.common}")
	srcName := utils.ValueString("${sweet.moduleName.src}")

	str := `import (`
	str += line
	str += `    "encoding/json"`
	str += line
	str += fmt.Sprintf(`    "%s/vo"`, commonName)
	str += line
	str += fmt.Sprintf(`    "%s/main/golang/service"`, srcName)
	str += line
	str += fmt.Sprintf(`    "%s/main/golang/models"`, srcName)
	str += line
	str += `)`
	str += line
	str += line
	return str
}

func getControllerIdType(tb model.Tables) string {
	for _, filed := range tb.Fileds {
		if filed.Key {
			filedType := getFiledType(filed.Type)
			if filedType == "string" {
				return `id := that.GetString(":id")`
			} else if filedType == "int32" {
				return `id, _ := that.GetInt32(":id")`
			} else if filedType == "int64" {
				return `id, _ := that.GetInt64(":id")`
			} else {
				panic("不正确的ID类型: " + filed.Name)
			}
		}
	}
	panic("未查询到主键: " + tb.Name)
}
