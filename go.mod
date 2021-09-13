module github.com/ysicing/k3s-autoscaler

go 1.16

require (
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/spf13/pflag v1.0.5
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.249 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm v1.0.249 // indirect
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	k8s.io/klog/v2 v2.4.0
	sigs.k8s.io/controller-runtime v0.8.3
)
