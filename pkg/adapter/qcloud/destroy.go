// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package qcloud

import (
	"context"
	"github.com/ysicing/k3s-autoscaler/pkg/provider"
)

func (q *qcloud) Destroy(ctx context.Context, instance *provider.Instance) error  {
	return nil
}
