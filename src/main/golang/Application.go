package serviceMain

import (
	"os"
	"sweet-common/utils"
	"sweet-src/main/golang/service"
)

func Main() {
	os.RemoveAll("generatorFile")
	utils.InitYml()
	service.Generator()
}
