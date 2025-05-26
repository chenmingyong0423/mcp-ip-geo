package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/chenmingyong0423/gkit/slice"
	"github.com/chenmingyong0423/mcp-ip-geo/internal/domain"

	httpchain "github.com/chenmingyong0423/go-http-chain"
)

func NewIpApiService() *IpApiService {
	return &IpApiService{
		host: "http://ip-api.com",
		client: httpchain.NewWithClient(&http.Client{
			Timeout: time.Second * 10,
		}),
	}
}

type IIpApiService interface {
	GetLocation(ctx context.Context, ip string) (*domain.IpApiResponse, error)
	BatchGetLocation(ctx context.Context, ips []string) ([]domain.IpApiResponse, error)
}

var _ IIpApiService = (*IpApiService)(nil)

type IpApiService struct {
	host   string
	client *httpchain.Client
}

func (s *IpApiService) BatchGetLocation(ctx context.Context, ips []string) ([]domain.IpApiResponse, error) {
	const batchSize = 100
	numberOfBatches := (len(ips) + batchSize - 1) / batchSize     // 计算批次数量
	results := make(chan []domain.IpApiResponse, numberOfBatches) // 创建带有足够缓冲的通道
	defer func() {
		close(results)
	}()
	var eg errgroup.Group
	// 将 IPs 切分为每个批次大小为 100
	for start := 0; start < len(ips); start += batchSize {
		end := start + batchSize
		if end > len(ips) {
			end = len(ips)
		}
		// 准备当前批次的 IPs
		batchIps := ips[start:end]
		eg.Go(func() error {
			batchBody := slice.Map(batchIps, func(idx int, ip string) domain.IpApiRequestBody {
				return domain.IpApiRequestBody{
					Query: ip,
					Lang:  "zh-CN",
				}
			})

			// 发起 API 调用
			batchResult := make([]domain.IpApiResponse, 0, len(batchIps))
			err := s.client.Post(s.host+"/batch").SetBody(batchBody).SetBodyEncodeFunc(func(body any) (io.Reader, error) {
				marshal, err := json.Marshal(body)
				if err != nil {
					return nil, err
				}
				return io.NopCloser(bytes.NewReader(marshal)), nil
			}).DoAndParse(ctx, &batchResult)
			if err != nil {
				return err
			}

			results <- batchResult
			return nil
		})
	}

	err := eg.Wait()
	if err != nil {
		return nil, err
	}
	var finalResult = make([]domain.IpApiResponse, 0, len(ips))
	for res := range results {
		finalResult = append(finalResult, res...)
		numberOfBatches--
		if numberOfBatches == 0 {
			break
		}
	}
	return finalResult, nil
}

func (s *IpApiService) GetLocation(ctx context.Context, ip string) (*domain.IpApiResponse, error) {
	var resp domain.IpApiResponse
	err := s.client.Get(fmt.Sprintf("%s/json/%s", s.host, ip)).DoAndParse(ctx, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
