package main

import (
	"github.com/gogf/gf/frame/g"
	_ "san616qi/app/timer"
	_ "san616qi/boot"
	_ "san616qi/router"
)

func main() {
	g.Server().Run()
}
