package main

import (
	"github.com/zeromicro/go-zero/core/load"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/tools/goctl/cmd"
)

func main() {
	logx.MustSetup(logx.LogConf{Mode: "console", Encoding: "plain"})
	load.Disable()
	cmd.Execute()
}
