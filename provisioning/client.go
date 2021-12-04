package provisioning

import (
	aarch64 "github.com/natesales/aarch64-client-go"
)

type A64Client struct {
	aarch64.Client
}

func NewA64Client(APIKey string) *A64Client {
	return &A64Client{aarch64.NewClient(APIKey)}
}
