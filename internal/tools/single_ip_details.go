package tools

import (
	"context"
	"encoding/json"
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/chenmingyong0423/mcp-ip-geo/internal/service"
)

var singleIpParserTool *protocol.Tool

type ipRequest struct {
	Ip string `json:"ip"`
}

func init() {
	var err error
	singleIpParserTool, err = protocol.NewTool("ip-details", "a tool that provides IP geolocation information", ipRequest{})
	if err != nil {
		panic(err)
	}
}

func SingleIpParser() (*protocol.Tool, server.ToolHandlerFunc) {
	ipApiService := service.NewIpApiService()

	return singleIpParserTool, func(ctx context.Context, toolRequest *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
		var req ipRequest
		if err := protocol.VerifyAndUnmarshal(toolRequest.RawArguments, &req); err != nil {
			return nil, err
		}
		resp, err := ipApiService.GetLocation(context.Background(), req.Ip)
		if err != nil {
			return nil, err
		}

		marshal, err := json.Marshal(resp)
		if err != nil {
			return nil, err
		}

		return &protocol.CallToolResult{
			Content: []protocol.Content{
				protocol.TextContent{
					Type: "text",
					Text: string(marshal),
				},
			},
		}, nil
	}
}
