package service

import (
	"fmt"
	"github.com/PurpleScorpion/go-sweet-keqing/keqing"
	"os"
	"sweet-common/utils"
	"sweet-src/main/golang/model"
)

func generatorRouters(arr []model.Tables) {
	basePath := "generatorFile/" + utils.ValueString("${sweet.mainPath}")
	basePath = basePath + "/routers"
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		panic(err)
	}
	str := `package routers`
	str += line
	str += line

	str += getRouterImport()
	str += getRouterInit(arr)

	filePath := basePath + "/router.go"
	os.WriteFile(filePath, []byte(str), 0644)
}

func getRouterInit(arr []model.Tables) string {
	str := "func InitRouters() {"
	str += line
	for i := 0; i < len(arr); i++ {
		tb := arr[i]
		name := getGoModelName(tb.Name)
		str += `    beego.Router("/` + keqing.Camel2WebURL(name) + `/page-data", &controllers.` + name + `Controller{}, "post:PageData")`
		str += line
		str += `    beego.Router("/` + keqing.Camel2WebURL(name) + `/id/:id", &controllers.` + name + `Controller{}, "get:GetById")`
		str += line
		str += `    beego.Router("/` + keqing.Camel2WebURL(name) + `/delete/:id", &controllers.` + name + `Controller{}, "delete:DeleteById")`
		str += line
		str += `    beego.Router("/` + keqing.Camel2WebURL(name) + `/insert", &controllers.` + name + `Controller{}, "post:Insert")`
		str += line
		str += `    beego.Router("/` + keqing.Camel2WebURL(name) + `/update", &controllers.` + name + `Controller{}, "put:Update")`
		str += line
		str += line
	}
	str += `}`
	return str
}

func getRouterImport() string {
	srcName := utils.ValueString("${sweet.moduleName.src}")
	str := `import (`
	str += line
	str += `    beego "github.com/beego/beego/v2/server/web"`
	str += line
	str += fmt.Sprintf(`    "%s/main/golang/controllers"`, srcName)
	str += line
	str += `)`
	str += line
	str += line
	return str
}
