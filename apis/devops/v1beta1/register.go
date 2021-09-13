// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package v1beta1

import "k8s.io/apimachinery/pkg/runtime/schema"

// SchemeGroupVersion is group version used to register these objects.
var SchemeGroupVersion = GroupVersion

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
