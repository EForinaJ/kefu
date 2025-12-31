package main

import (
	_ "kefu-server/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"kefu-server/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
