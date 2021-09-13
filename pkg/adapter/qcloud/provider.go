// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package qcloud

import "github.com/ysicing/k3s-autoscaler/pkg/provider"

type qcloud struct {
	region string
	systemDiskType DiskType // 系统盘类型
	systemDiskSize int // 系统盘大小
	paidType string // 付费模式
	instanceType string // 机型
}

func New(opts ...Option) provider.Provider {
	p := new(qcloud)
	for _, opt := range opts {
		opt(p)
	}
	if p.systemDiskSize == 0 {
		p.systemDiskSize = 40
	}
	return p
}