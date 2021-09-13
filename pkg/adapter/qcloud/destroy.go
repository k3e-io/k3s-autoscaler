// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package qcloud

import (
	"context"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"github.com/ysicing/k3s-autoscaler/pkg/provider"
)

func (q *qcloud) Destroy(ctx context.Context, instance *provider.Instance) error  {
	client := q.getClient()
	request := cvm.NewTerminateInstancesRequest()
	request.InstanceIds = common.StringPtrs([]string{instance.ID})
	_, err := client.TerminateInstances(request)
	if err != nil {
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
		}
		return err
	}
	return nil
}
