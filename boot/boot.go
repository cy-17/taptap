package boot

import (
	_ "san616qi/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
)

// 用于应用初始化。
func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
}
