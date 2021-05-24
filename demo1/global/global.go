package global

import (
	"demo1/config"
	"demo1/model"
)

var (
	G_CONFIG   config.Config
	G_JSONPATH string
	G_JSONDATA map[string]model.Army
)
