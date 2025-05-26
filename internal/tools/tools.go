package tools

import (
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
)

type ToolFunc func() (tool *protocol.Tool, toolHandler server.ToolHandlerFunc)

func GetToolFuncList() []ToolFunc {
	return []ToolFunc{
		SingleIpParser,
		MultipleIpParser,
	}
}
