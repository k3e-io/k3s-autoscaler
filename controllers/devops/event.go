// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package devops

import (
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

type eventHandler struct {
	reader client.Reader
}

func (e *eventHandler) Create(evt event.CreateEvent, q workqueue.RateLimitingInterface) {
}

func (e *eventHandler) Update(evt event.UpdateEvent, q workqueue.RateLimitingInterface) {
}

func (e *eventHandler) Delete(evt event.DeleteEvent, q workqueue.RateLimitingInterface) {
}

func (e *eventHandler) Generic(evt event.GenericEvent, q workqueue.RateLimitingInterface) {
}
