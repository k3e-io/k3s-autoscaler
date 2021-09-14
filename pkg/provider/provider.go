// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package provider

import (
	"context"
)

// ProviderType specifies the hosting provider.
type ProviderType string

// Value converts the value to a sql string.
func (s ProviderType) Value() string {
	return string(s)
}

// Provider type enumeration.
const (
	ProviderAliyun = ProviderType("aliyun")
	ProviderQcloud = ProviderType("qcloud")
)

// A Provider represents a hosting provider, such as
type Provider interface {
	// Create creates a new server.
	Create(context.Context, InstanceCreateOpts) (*Instance, error)
	// Destroy destroys an existing server.
	Destroy(context.Context, *Instance) error
	// Show existing server
	Show(ctx context.Context) []Instance
}

// An Instance represents a server instance
type Instance struct {
	Provider           ProviderType
	ID                 string
	Name               string
	Region             string
	InstanceType       string // 实例规格
	ImageID            string // 实例镜像
	InstanceChargeType string // 付费模式
}

// InstanceCreateOpts define soptional instructions for
// creating server instances.
type InstanceCreateOpts struct {
	Name string
}
