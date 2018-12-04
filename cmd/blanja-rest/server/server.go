package server

import (
	"fmt"
	"github.com/maulidihsan/interop-commerce/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(fmt.Sprintf(":%s", config.GetString("blanja.rest.port")))
}
