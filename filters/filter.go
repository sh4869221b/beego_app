package filters

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/context"
)

// LogManager is log output function
var LogManager = func(ctx *context.Context) {
	fmt.Println("IP::" + ctx.Request.RemoteAddr + ", Url::" + ctx.Request.RequestURI + ", Time::" + time.Now().Format(time.RFC3339Nano))
}
