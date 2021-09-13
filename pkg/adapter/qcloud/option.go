// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package qcloud

// DiskType 磁盘类型
type DiskType string

func (dt DiskType) Value() string {
	return string(dt)
}

const (
	SystemDiskLocal = DiskType("LOCAL_BASIC") // 本地硬盘
	SystemDiskLocalSSD = DiskType("LOCAL_SSD") // 本地ssd
	SystemDiskBasic = DiskType("CLOUD_BASIC") // 普通云盘
	SystemDiskSSD = DiskType("CLOUD_SSD") // ssd云盘
	SystemDiskPremium = DiskType("CLOUD_PREMIUM") // 高性能云盘
	SystemDiskDefault = SystemDiskPremium // 默认高性能云盘
)

// PaidType 计费类型
type PaidType string

const (
	PaidTypeSpot = PaidType("SPOTPAID") // 竞价
	PaidTypePre = PaidType("PREPAID") // 包月
 	PaidTypeHOUR = PaidType("POSTPAID_BY_HOUR") // 按量
	PaidTypeDefault = PaidTypeSpot
)

func (p PaidType) Value() string {
	return string(p)
}

type InstanceType string

const (
	CustomInstanceSA2 = InstanceType("SA2.SMALL1") // 0.04/h
	DefaultInstance = CustomInstanceSA2
)

func (in InstanceType) Value() string  {
	return string(in)
}

type Option func(qcloud2 *qcloud)

func WithSystemDiskType(sdtype DiskType) Option {
	return func(q *qcloud) {
		q.systemDiskType = sdtype
	}
}

func WithSystemDiskSize(sdsize int) Option  {
	return func(q *qcloud) {
		q.systemDiskSize = sdsize
	}
}

func WithPaidType(paid PaidType) Option  {
	return func(q *qcloud) {
		q.paidType = paid
	}
}

func WithInstanceType(intype InstanceType) Option  {
	return func(q *qcloud) {
		q.instanceType = intype
	}
}

func WithImageID() Option {
	return func(q *qcloud) {
		q.imageID = "img-h1yvvfw1"
	}
}
