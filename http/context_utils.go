package http

import (
	"context"
)

var (
	RequestIDCtxKey = "request_id"
	TraceIDCtxKey   = "trace_id"
	TenantIDCtxKey  = "tenant_id"
	HostCtxKey      = "host"
	RequestIPCtxKey = "request_ip"
	LangCtxKey      = "lang"
)

var (
	RequestIDHeader = "X-Request-Id"
	TraceIDHeader   = "X-Trace-Id"
	TenantIDHeader  = "X-Tenant-Id"
	LangHeader      = "Accept-Language"
)

func GetCtxValueStr(c context.Context, key string) string {
	value, ok := c.Value(key).(string)
	if !ok {
		return ""
	}
	return value
}
