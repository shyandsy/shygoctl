package main

import (
	"github.com/shyandsy/shygoctl/cmd"
	"github.com/zeromicro/go-zero/core/load"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	logx.MustSetup(logx.LogConf{Mode: "console", Encoding: "plain"})
	load.Disable()
	cmd.Execute()
}
