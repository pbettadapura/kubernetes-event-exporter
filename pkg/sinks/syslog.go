package sinks

import (
	"context"
	"encoding/json"
	"github.com/resmoio/kubernetes-event-exporter/pkg/kube"
	"github.com/RackSec/srslog"
)

type SyslogConfig struct {
	Network string `yaml:"network"`
	Address string `yaml:"address"`
	Tag     string `yaml:"tag"`
}

type SyslogSink struct {
	sw *srslog.Writer
}

func NewSyslogSink(config *SyslogConfig) (Sink, error) {
	w, err := srslog.Dial(config.Network, config.Address, srslog.LOG_LOCAL0, config.Tag)
	if err != nil {
		return nil, err
	}
	return &SyslogSink{sw: w}, nil
}

func (w *SyslogSink) Close() {
	w.sw.Close()
}

func (w *SyslogSink) Send(ctx context.Context, ev *kube.EnhancedEvent) error {

	if b, err := json.Marshal(ev); err == nil {
		_, writeErr := w.sw.Write(b)

		if writeErr != nil {
			return writeErr
		}
	} else {
		return err
	}
	return nil
}
