package serial

import (
	"context"
	"github.com/wuyuanyi135/mvcamctrl/serial/command"
)

type SerialCommand struct {
	Command         command.CommandMeta
	Arg             []byte
	ResponseChannel chan []byte
	Ctx             context.Context
}
