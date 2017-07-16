package test

import (
	"path/filepath"
	"runtime"

	_ "rkl.io/latex-renderer/daemon/routers"

	"rkl.io/latex-renderer/config"
	"github.com/astaxie/beego"
	"log"
)

func init() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}

	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}
