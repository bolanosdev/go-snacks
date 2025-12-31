package sentry

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
)

type SentryConfig struct {
	DSN string
}

type SentryObs struct {
	hub *sentry.Hub
	cfg SentryConfig
}

func NewSentryObs(cfg SentryConfig) (*SentryObs, error) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:            cfg.DSN,
		Debug:          true,
		SendDefaultPII: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup sentry")
	}

	sentryHub := sentry.CurrentHub()
	if sentryHub == nil {
		return nil, errors.New("failed to get sentry hub")
	}

	return &SentryObs{
		hub: sentryHub,
		cfg: cfg,
	}, nil
}

func (s *SentryObs) Flush() {
	sentry.Flush(2 * time.Second)
}

func (s *SentryObs) CaptureError(err error, status_code int) *sentry.EventID {
	var event_id *sentry.EventID

	s.hub.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.LevelError)
		scope.SetExtra("status_code", status_code)

		type metadataError interface {
			GetMetadata() map[string]interface{}
		}

		var mdErr metadataError
		if errors.As(err, &mdErr) {
			metadata := mdErr.GetMetadata()
			for key, value := range metadata {
				scope.SetExtra(key, value)
			}
		}

		event_id = s.hub.CaptureException(err)
	})

	return event_id
}
