package serviceMain

import (
	"sweet-common/utils"
	"sweet-src/main/golang/service"
)

func Main() {
	utils.InitYml()
	service.Generator()
}
