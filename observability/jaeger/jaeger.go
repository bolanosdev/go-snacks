package jaeger

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type JaegerConfig struct {
	Name              string
	Hostname          string
	SensitiveKeywords []string
}

// JaegerInterface defines the interface for tracing operations
type JaegerInterface interface {
	WithConfig(cfg JaegerConfig) JaegerInterface
	Initialize() (JaegerInterface, error)
	Trace(c context.Context, name string) (context.Context, trace.Span)
	TraceFunc(c context.Context) context.Context
	TraceDB(c context.Context, query string, args interface{}) context.Context
}

type JaegerObs struct {
	cfg JaegerConfig
	ctx context.Context
	tp  trace.TracerProvider
}

func NewJaegerObs(ctx context.Context) JaegerObs {
	tp := otel.GetTracerProvider()
	return JaegerObs{
		cfg: JaegerConfig{
			Name:              "tracer",
			Hostname:          "",
			SensitiveKeywords: []string{},
		},
		ctx: ctx,
		tp:  tp,
	}
}

func (t JaegerObs) WithConfig(cfg JaegerConfig) JaegerInterface {
	t.cfg = cfg
	return t
}

func (t JaegerObs) Initialize() (JaegerInterface, error) {
	// dont register trace provider if JAEGER information isnt provided through app.yaml
	if t.cfg.Hostname == "" {
		return t, errors.New("missing jaeger dial hostname")
	}

	res, err := resource.New(t.ctx, resource.WithAttributes(
		semconv.ServiceName(t.cfg.Name),
	))

	if err != nil {
		return t, errors.Wrap(err, "failed to create resource for jaeger")
	}

	ctx, cancel := context.WithTimeout(t.ctx, time.Second)
	defer cancel()

	conn, err := grpc.NewClient(t.cfg.Hostname,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return t, errors.Wrap(err, "failed to create grpc connection for jaeger")
	}

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn), otlptracegrpc.WithTimeout(1000*time.Millisecond))
	if err != nil {
		return t, errors.Wrap(err, "failed to create exporter for jaeger")
	}

	processor := sdktrace.NewSimpleSpanProcessor(exporter)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(processor),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return t, nil
}

func (t JaegerObs) TraceFunc(ctx context.Context) context.Context {
	pc, _, _, _ := runtime.Caller(1)

	funcName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(funcName, ".")
	spanName := parts[len(parts)-1]
	if len(parts) >= 2 {
		receiver := parts[len(parts)-2]
		receiver = strings.TrimPrefix(receiver, "(*")
		receiver = strings.TrimSuffix(receiver, ")")
		spanName = receiver + "." + spanName
	}

	tracedCtx, span := t.Trace(ctx, spanName)
	defer span.End()

	return tracedCtx
}

func (t JaegerObs) TraceDB(ctx context.Context, query string, args interface{}) context.Context {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(funcName, ".")
	spanName := parts[len(parts)-1]
	if len(parts) >= 2 {
		receiver := parts[len(parts)-2]
		receiver = strings.TrimPrefix(receiver, "(*")
		receiver = strings.TrimSuffix(receiver, ")")
		spanName = receiver + "." + spanName
	}

	tracedCtx, span := t.Trace(ctx, spanName)
	defer span.End()

	span.SetAttributes(
		attribute.String("db.statement", query),
	)

	if args != nil {
		argsStr := fmt.Sprintf("%+v", args)
		maskedArgs := t.MaskSensitiveData(argsStr)
		span.SetAttributes(
			attribute.String("db.args", maskedArgs),
		)
	}

	return tracedCtx
}

func (t JaegerObs) Trace(c context.Context, name string) (context.Context, trace.Span) {
	ctx, span := t.tp.Tracer(t.cfg.Name).Start(c, name)

	return ctx, span
}

func (t JaegerObs) MaskSensitiveData(argsStr string) string {
	masked := argsStr

	for _, keyword := range t.cfg.SensitiveKeywords {
		if strings.Contains(strings.ToLower(masked), strings.ToLower(keyword)) {
			lowerMasked := strings.ToLower(masked)
			idx := strings.Index(lowerMasked, strings.ToLower(keyword))

			if idx != -1 {
				valueStart := idx + len(keyword)
				for valueStart < len(masked) && (masked[valueStart] == ':' || masked[valueStart] == ' ') {
					valueStart++
				}

				valueEnd := valueStart
				for valueEnd < len(masked) && masked[valueEnd] != ' ' && masked[valueEnd] != ',' && masked[valueEnd] != '}' && masked[valueEnd] != ']' {
					valueEnd++
				}

				if valueEnd > valueStart {
					masked = masked[:valueStart] + "***" + masked[valueEnd:]
				}
			}
		}
	}

	return masked
}
