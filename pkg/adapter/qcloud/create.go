// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package qcloud

import (
	"context"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"github.com/ysicing/k3s-autoscaler/pkg/provider"
	"k8s.io/klog/v2"
)

func (q *qcloud) Create(ctx context.Context, opts provider.InstanceCreateOpts) (*provider.Instance, error ) {
	return nil, nil
}

func (q *qcloud) setupKeypair(ctx context.Context) error {
	return nil
}

func (q *qcloud) create(ctx context.Context) (*provider.Instance, error) {
	client := q.getClient()
	request := cvm.NewRunInstancesRequest()
	request.InstanceChargeType = common.StringPtr(q.paidType.Value())
	request.InstanceType = common.StringPtr(q.instanceType.Value())
	request.ImageId = common.StringPtr(q.imageID)
	request.SystemDisk = &cvm.SystemDisk {
		DiskType: common.StringPtr(q.systemDiskType.Value()),
		DiskSize: common.Int64Ptr(40),
	}
	request.VirtualPrivateCloud = &cvm.VirtualPrivateCloud {
		VpcId: common.StringPtr(q.VpcID),
		SubnetId: common.StringPtr(q.SubnetID),
	}
	request.InternetAccessible = &cvm.InternetAccessible {
		InternetChargeType: common.StringPtr("TRAFFIC_POSTPAID_BY_HOUR"),
		InternetMaxBandwidthOut: common.Int64Ptr(100),
		PublicIpAssigned: common.BoolPtr(true),
	}
	request.InstanceCount = common.Int64Ptr(1)
	request.InstanceName = common.StringPtr(q.name)
	request.LoginSettings = &cvm.LoginSettings {
		KeyIds: common.StringPtrs(q.KeyID),
	}
	request.EnhancedService = &cvm.EnhancedService {
		SecurityService: &cvm.RunSecurityServiceEnabled {
			Enabled: common.BoolPtr(true),
		},
		MonitorService: &cvm.RunMonitorServiceEnabled {
			Enabled: common.BoolPtr(true),
		},
		AutomationService: &cvm.RunAutomationServiceEnabled {
			Enabled: common.BoolPtr(true),
		},
	}
	request.TagSpecification = []*cvm.TagSpecification {
		&cvm.TagSpecification {
			ResourceType: common.StringPtr("instance"),
			Tags: []*cvm.Tag {
				&cvm.Tag {
					Key: common.StringPtr("k3s-cm"),
					Value: common.StringPtr("autoscaler"),
				},
			},
		},
	}
	request.InstanceMarketOptions = &cvm.InstanceMarketOptionsRequest {
		MarketType: common.StringPtr("spot"),
		SpotOptions: &cvm.SpotMarketOptions {
			MaxPrice: common.StringPtr("0.5"),
			SpotInstanceType: common.StringPtr("one-time"),
		},
	}
	request.UserData = common.StringPtr(q.UserData)
	request.DryRun = common.BoolPtr(true)
	response, err := client.RunInstances(request)
	if err != nil {
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			klog.Error("qcloud api error has returned: %s", err)
		}
		return nil, err
	}
	resp := response.Response.InstanceIdSet
	p := provider.Instance{}
	p.Provider = provider.PrviderQcloud
	p.ID = *resp[0]
	p.Region = q.region
	return &p, nil
}