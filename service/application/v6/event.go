// Package application code generated by oapi sdk gen
package application

import (
	"context"
)

/**
消息处理器定义
**/
type ApplicationCreatedEventHandler struct {
	handler func(context.Context, *ApplicationCreatedEvent) error
}

func NewApplicationCreatedEventHandler(handler func(context.Context, *ApplicationCreatedEvent) error) *ApplicationCreatedEventHandler {
	h := &ApplicationCreatedEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ApplicationCreatedEventHandler) Event() interface{} {
	return &ApplicationCreatedEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ApplicationCreatedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ApplicationCreatedEvent))
}

/**
消息处理器定义
**/
type ApplicationAppVersionAuditEventHandler struct {
	handler func(context.Context, *ApplicationAppVersionAuditEvent) error
}

func NewApplicationAppVersionAuditEventHandler(handler func(context.Context, *ApplicationAppVersionAuditEvent) error) *ApplicationAppVersionAuditEventHandler {
	h := &ApplicationAppVersionAuditEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ApplicationAppVersionAuditEventHandler) Event() interface{} {
	return &ApplicationAppVersionAuditEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ApplicationAppVersionAuditEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ApplicationAppVersionAuditEvent))
}

/**
消息处理器定义
**/
type ApplicationAppVersionPublishApplyEventHandler struct {
	handler func(context.Context, *ApplicationAppVersionPublishApplyEvent) error
}

func NewApplicationAppVersionPublishApplyEventHandler(handler func(context.Context, *ApplicationAppVersionPublishApplyEvent) error) *ApplicationAppVersionPublishApplyEventHandler {
	h := &ApplicationAppVersionPublishApplyEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ApplicationAppVersionPublishApplyEventHandler) Event() interface{} {
	return &ApplicationAppVersionPublishApplyEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ApplicationAppVersionPublishApplyEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ApplicationAppVersionPublishApplyEvent))
}

/**
消息处理器定义
**/
type ApplicationAppVersionPublishRevokeEventHandler struct {
	handler func(context.Context, *ApplicationAppVersionPublishRevokeEvent) error
}

func NewApplicationAppVersionPublishRevokeEventHandler(handler func(context.Context, *ApplicationAppVersionPublishRevokeEvent) error) *ApplicationAppVersionPublishRevokeEventHandler {
	h := &ApplicationAppVersionPublishRevokeEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ApplicationAppVersionPublishRevokeEventHandler) Event() interface{} {
	return &ApplicationAppVersionPublishRevokeEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ApplicationAppVersionPublishRevokeEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ApplicationAppVersionPublishRevokeEvent))
}

/**
消息处理器定义
**/
type ApplicationFeedbackCreatedEventHandler struct {
	handler func(context.Context, *ApplicationFeedbackCreatedEvent) error
}

func NewApplicationFeedbackCreatedEventHandler(handler func(context.Context, *ApplicationFeedbackCreatedEvent) error) *ApplicationFeedbackCreatedEventHandler {
	h := &ApplicationFeedbackCreatedEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ApplicationFeedbackCreatedEventHandler) Event() interface{} {
	return &ApplicationFeedbackCreatedEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ApplicationFeedbackCreatedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ApplicationFeedbackCreatedEvent))
}

/**
消息处理器定义
**/
type ApplicationFeedbackUpdatedEventHandler struct {
	handler func(context.Context, *ApplicationFeedbackUpdatedEvent) error
}

func NewApplicationFeedbackUpdatedEventHandler(handler func(context.Context, *ApplicationFeedbackUpdatedEvent) error) *ApplicationFeedbackUpdatedEventHandler {
	h := &ApplicationFeedbackUpdatedEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ApplicationFeedbackUpdatedEventHandler) Event() interface{} {
	return &ApplicationFeedbackUpdatedEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ApplicationFeedbackUpdatedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ApplicationFeedbackUpdatedEvent))
}

/**
消息处理器定义
**/
type ApplicationVisibilityAddedEventHandler struct {
	handler func(context.Context, *ApplicationVisibilityAddedEvent) error
}

func NewApplicationVisibilityAddedEventHandler(handler func(context.Context, *ApplicationVisibilityAddedEvent) error) *ApplicationVisibilityAddedEventHandler {
	h := &ApplicationVisibilityAddedEventHandler{handler: handler}
	return h
}

/**
返回事件的消息体的实例，用于反序列化用
**/
func (h *ApplicationVisibilityAddedEventHandler) Event() interface{} {
	return &ApplicationVisibilityAddedEvent{}
}

/**
回调开发者注册的handle
r**/
func (h *ApplicationVisibilityAddedEventHandler) Handle(ctx context.Context, event interface{}) error {
	return h.handler(ctx, event.(*ApplicationVisibilityAddedEvent))
}