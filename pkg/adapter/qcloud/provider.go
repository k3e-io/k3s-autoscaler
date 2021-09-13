// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package qcloud

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/ysicing/k3s-autoscaler/pkg/provider"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type qcloud struct {
	region string
	name string
	systemDiskType DiskType // 系统盘类型
	systemDiskSize int // 系统盘大小
	paidType PaidType // 付费模式
	instanceType InstanceType // 机型
	imageID string // 镜像ID
	VpcID string
	SubnetID string
	KeyID []string
	UserData string
	ApiKey string
	ApiSecret string
}

func (q *qcloud) getClient() *cvm.Client {
	credential := common.NewCredential(
		q.ApiKey,
		q.ApiSecret,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	client, _ := cvm.NewClient(credential, fmt.Sprintf("ap-%v", q.region), cpf)
	return client
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