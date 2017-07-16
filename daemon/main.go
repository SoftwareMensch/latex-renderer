package main

import (
	_ "rkl.io/latex-renderer/daemon/routers"

	"rkl.io/latex-renderer/config"
	"github.com/astaxie/beego"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
