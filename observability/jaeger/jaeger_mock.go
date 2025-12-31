package jaeger

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type MockTracer struct {
	name string
	cfg  JaegerConfig
}

func NewMockTracer() MockTracer {
	return MockTracer{}
}

func (m MockTracer) WithConfig(cfg JaegerConfig) JaegerInterface {
	m.cfg = cfg
	return m
}

func (m MockTracer) Initialize() (JaegerInterface, error) {
	return m, nil
}

func (m MockTracer) Trace(c context.Context, name string) (context.Context, trace.Span) {
	return c, trace.SpanFromContext(c)
}

func (m MockTracer) TraceFunc(c context.Context) context.Context {
	return c
}

func (m MockTracer) TraceDB(c context.Context, query string, args interface{}) context.Context {
	return c
}
