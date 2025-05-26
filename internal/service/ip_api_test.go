package service

import (
	"context"
	"testing"
)

func TestIpApiService_GetLocation(t *testing.T) {
	s := NewIpApiService()
	resp, err := s.GetLocation(context.Background(), "1.1.1.1")
	if err != nil {
		t.Fatalf("GetLocation error: %v", err)
	}
	if resp.Status != "success" {
		t.Fatalf("GetLocation status error: %v", resp.Status)
	}
	t.Log(resp)
}

func TestIpApiService_BatchGetLocation(t *testing.T) {
	s := NewIpApiService()
	ips := []string{"1.1.1.1", "208.80.152.201"}
	resp, err := s.BatchGetLocation(context.Background(), ips)
	if err != nil {
		t.Fatalf("BatchGetLocation error: %v", err)
	}

	if len(resp) != len(ips) {
		t.Fatalf("BatchGetLocation expected %d results, got %d", len(ips), len(resp))
	}

	for _, r := range resp {
		if r.Status != "success" {
			t.Fatalf("BatchGetLocation status error: %v", r.Status)
		}
		t.Log(r)
	}
}
