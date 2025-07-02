package service

import (
	"fmt"
	"os"
	"sweet-common/utils"
	"sweet-src/main/golang/model"
)

func generatorService(basePath string, tb model.Tables) {
	basePath = basePath + "/service"
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		panic(err)
	}
	var name = getGoModelName(tb.Name)
	str := `package service`
	str += line
	str += line
	str += getServiceImport(tb)
	str += getServiceStruct(tb)
	str += getServicePageData(tb)
	str += getServiceGetById(tb)
	str += getServiceSave(tb)
	str += getServiceUpdate(tb)
	str += getServiceDeleteById(tb)
	filePath := basePath + "/" + utils.SnakeToPascal(name) + "Service.go"
	os.WriteFile(filePath, []byte(str), 0644)

}

func getServiceUpdate(tb model.Tables) string {
	var name = getGoModelName(tb.Name)
	str := `func (that *` + name + `Service) Update(pojo models.` + name + `) utils.R {`
	str += line
	str += `    qw := mapper.BuilderQueryWrapper(&models.` + name + `{})`
	str += line
	str += `    qw.Eq(true,"` + getTableKey(tb) + `",pojo.` + utils.SnakeToPascal(getTableKey(tb)) + `)`
	str += line
	for _, filed := range tb.Fileds {
		if filed.Key {
			continue
		}
		if getFiledType(filed.Type) == "int32" || getFiledType(filed.Type) == "int64" {
			str += `    qw.Set(pojo.` + utils.SnakeToPascal(filed.Name) + ` > 0,"` + filed.Name + `",pojo.` + utils.SnakeToPascal(filed.Name) + `)`
		} else if getFiledType(filed.Type) == "string" {
			str += `    qw.Set(utils.IsNotEmpty(pojo.` + utils.SnakeToPascal(filed.Name) + `),"` + filed.Name + `",pojo.` + utils.SnakeToPascal(filed.Name) + `)`
		} else {
			continue
		}
		str += line
	}
	str += `    count := mapper.Update(qw)`
	str += line
	str += `    if count == 0 {`
	str += line
	str += `        return utils.Fail(500, "Update failed")`
	str += line
	str += `    }`
	str += line
	str += `    return utils.Success("")`
	str += line
	str += `}`
	str += line
	return str
}

func getServiceSave(tb model.Tables) string {
	var name = getGoModelName(tb.Name)
	str := `
func (that *` + name + `Service) Insert(pojo models.` + name + `) utils.R {
    count := mapper.InsertCustom(&pojo, true)
	if count == 0 {
		return utils.Fail(500, "Insert failed")
	}
	return utils.Success("")
}
`
	str += line
	return str
}

func getServiceGetById(tb model.Tables) string {
	var name = getGoModelName(tb.Name)
	str := `
func (that *` + name + `Service) GetById(id ` + getServiceIdType(tb) + `) utils.R {
    var list []models.` + name + `
    mapper.SelectById(&list, id)
	if len(list) == 0 {
		return utils.Fail(500, "Data does not exist")
	}
	return utils.Success(list[0])
}
`
	str += line
	return str
}
func getServiceDeleteById(tb model.Tables) string {
	var name = getGoModelName(tb.Name)
	str := `
func (that *` + name + `Service) DeleteById(id ` + getServiceIdType(tb) + `) utils.R {
    count := mapper.DeleteById(&models.` + name + `{}, id)
	if count == 0 {
		return utils.Fail(500, "Delete failed")
	}
	return utils.Success("")
}
`
	str += line
	return str
}

func getServicePageData(tb model.Tables) string {
	var name = getGoModelName(tb.Name)
	str := `
func (that *` + name + `Service) PageData(tmp vo.` + name + `VO) utils.R {
    var list []models.` + name + `
    qw := mapper.BuilderQueryWrapper(&list)
    page := mapper.BuilderPageUtils(tmp.Current, tmp.PageSize, qw)
	pageData := mapper.Page(page)
	return utils.Success(pageData)
}
`
	str += line
	return str
}

func getServiceStruct(tb model.Tables) string {
	str := `
type ` + getGoModelName(tb.Name) + `Service struct {
}`
	str += line
	return str
}

func getServiceImport(tb model.Tables) string {
	commonName := utils.ValueString("${sweet.moduleName.common}")
	srcName := utils.ValueString("${sweet.moduleName.src}")

	str := `import (`
	str += line
	str += `    "github.com/PurpleScorpion/go-sweet-orm/v2/mapper"`
	str += line
	str += fmt.Sprintf(`    "%s/vo"`, commonName)
	str += line
	str += fmt.Sprintf(`    "%s/utils"`, commonName)
	str += line
	str += fmt.Sprintf(`    "%s/main/golang/models"`, srcName)
	str += line
	str += `)`
	str += line
	str += line
	return str
}

func getServiceIdType(tb model.Tables) string {
	for _, filed := range tb.Fileds {
		if filed.Key {
			return getFiledType(filed.Type)
		}
	}
	panic("未查询到主键: " + tb.Name)
}

func getTableKey(tb model.Tables) string {
	for _, filed := range tb.Fileds {
		if filed.Key {
			return filed.Name
		}
	}
	panic("未查询到主键: " + tb.Name)
}
