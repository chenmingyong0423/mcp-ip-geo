package server

import (
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
	"github.com/chenmingyong0423/mcp-ip-geo/internal/tools"
)

func Run(address string) error {
	var err error

	var ts transport.ServerTransport
	if address == "" {
		ts = transport.NewStdioServerTransport()
	} else {
		ts = transport.NewStreamableHTTPServerTransport(address)
	}

	s, err := server.NewServer(ts)
	if err != nil {
		return err
	}

	toolFuncList := tools.GetToolFuncList()
	for _, tool := range toolFuncList {
		s.RegisterTool(tool())
	}

	return s.Run()
}
