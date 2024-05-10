package main

import (
	_ "hello_world_demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"hello_world_demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
