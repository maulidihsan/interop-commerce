package server

import (
	"fmt"
	"github.com/maulidihsan/flashdeal-webservice/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(fmt.Sprintf(":%s", config.GetString("blanja.rest.port")))
}