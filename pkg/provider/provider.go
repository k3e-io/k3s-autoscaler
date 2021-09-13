// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package provider

import "context"

// ProviderType specifies the cloud provider 服务商类型
type ProviderType string

const (
	PrviderQcloud = ProviderType("qcloud")
	PrviderAliyun = ProviderType("aliyun")
)

type Provider interface {
	// Create creates a new instance.
	Create(context.Context, InstanceCreateOpts) (*Instance, error)
	// Destroy destroys an existing instance.
	Destroy(context.Context, *Instance) error
}

type Instance struct {
	Provider ProviderType
	ID string
	Name string
	Address string
	Region string
}

type InstanceCreateOpts struct {
	Name string
}
