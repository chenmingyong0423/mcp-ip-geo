package tools

import (
	"context"
	"encoding/json"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/chenmingyong0423/mcp-ip-geo/internal/service"
)

var multipleIpParserTool *protocol.Tool

type multipleIpRequest struct {
	Ips []string `json:"ips"`
}

func init() {
	var err error
	multipleIpParserTool, err = protocol.NewTool("multiple-ip-details", "a tool that provides IPs geolocation information", multipleIpRequest{})
	if err != nil {
		panic(err)
	}
}

func MultipleIpParser() (*protocol.Tool, server.ToolHandlerFunc) {
	ipApiService := service.NewIpApiService()

	return multipleIpParserTool, func(ctx context.Context, toolRequest *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
		var req multipleIpRequest
		if err := protocol.VerifyAndUnmarshal(toolRequest.RawArguments, &req); err != nil {
			return nil, err
		}
		resp, err := ipApiService.BatchGetLocation(context.Background(), req.Ips)
		if err != nil {
			return nil, err
		}

		marshal, err := json.Marshal(resp)
		if err != nil {
			return nil, err
		}

		return &protocol.CallToolResult{
			Content: []protocol.Content{
				&protocol.TextContent{
					Type: "text",
					Text: string(marshal),
				},
			},
		}, nil
	}
}
