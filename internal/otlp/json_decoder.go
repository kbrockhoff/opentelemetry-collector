// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package otlp

import (
	"bytes"

	"github.com/gogo/protobuf/jsonpb"

	otlpcollectorlogs "go.opentelemetry.io/collector/internal/data/protogen/collector/logs/v1"
	otlpcollectormetrics "go.opentelemetry.io/collector/internal/data/protogen/collector/metrics/v1"
	otlpcollectortrace "go.opentelemetry.io/collector/internal/data/protogen/collector/trace/v1"
	"go.opentelemetry.io/collector/internal/model"
)

type decoder struct {
	delegate jsonpb.Unmarshaler
}

func newDecoder() *decoder {
	return &decoder{delegate: jsonpb.Unmarshaler{}}
}

// NewJSONTracesDecoder returns a serializer.TracesDecoder to decode from OTLP json bytes.
func NewJSONTracesDecoder() model.TracesDecoder {
	return newDecoder()
}

// NewJSONMetricsDecoder returns a serializer.MetricsDecoder to decode from OTLP json bytes.
func NewJSONMetricsDecoder() model.MetricsDecoder {
	return newDecoder()
}

// NewJSONLogsDecoder returns a serializer.LogsDecoder to decode from OTLP json bytes.
func NewJSONLogsDecoder() model.LogsDecoder {
	return newDecoder()
}

func (d *decoder) DecodeLogs(buf []byte) (interface{}, error) {
	ld := &otlpcollectorlogs.ExportLogsServiceRequest{}
	if err := d.delegate.Unmarshal(bytes.NewReader(buf), ld); err != nil {
		return nil, err
	}
	return ld, nil
}

func (d *decoder) DecodeMetrics(buf []byte) (interface{}, error) {
	md := &otlpcollectormetrics.ExportMetricsServiceRequest{}
	if err := d.delegate.Unmarshal(bytes.NewReader(buf), md); err != nil {
		return nil, err
	}
	return md, nil
}

func (d *decoder) DecodeTraces(buf []byte) (interface{}, error) {
	td := &otlpcollectortrace.ExportTraceServiceRequest{}
	if err := d.delegate.Unmarshal(bytes.NewReader(buf), td); err != nil {
		return nil, err
	}
	return td, nil
}
