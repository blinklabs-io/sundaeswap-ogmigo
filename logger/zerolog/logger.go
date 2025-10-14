package logger

import (
	ogmigo "github.com/SundaeSwap-finance/ogmigo/v6"
	"github.com/rs/zerolog"
)

type Logger struct {
	target zerolog.Logger
	kvs    []ogmigo.KeyValue
}

func Wrap(logger zerolog.Logger) ogmigo.Logger {
	return Logger{
		target: logger,
	}
}

func (l Logger) log(event *zerolog.Event, message string, kvs ...ogmigo.KeyValue) {
	for _, kv := range kvs {
		event = event.Str(kv.Key, kv.Value)
	}
	for _, kv := range kvs {
		event = event.Str(kv.Key, kv.Value)
	}
	event.Msg(message)
}

func (l Logger) Debug(message string, kvs ...ogmigo.KeyValue) {
	l.log(l.target.Debug(), message, kvs...)
}

func (l Logger) Info(message string, kvs ...ogmigo.KeyValue) {
	l.log(l.target.Info(), message, kvs...)
}

func (l Logger) With(kvs ...ogmigo.KeyValue) ogmigo.Logger {
	return Logger{
		target: l.target,
		kvs:    append(l.kvs, kvs...),
	}
}
