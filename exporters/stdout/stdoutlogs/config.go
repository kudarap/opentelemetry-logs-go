/*
Copyright Agoda Services Co.,Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package stdoutlogs

import (
	"io"
	"os"
)

var (
	defaultWriter      = os.Stdout
	defaultPrettyPrint = false
)

// config contains options for the STDOUT exporter.
type config struct {
	// Writer is the destination.  If not set, os.Stdout is used.
	Writer io.Writer

	// PrettyPrint will encode the output into readable JSON. Default is
	// false.
	PrettyPrint bool
}

// newConfig creates a validated Config configured with options.
func newConfig(options ...Option) (config, error) {
	cfg := config{
		Writer:      defaultWriter,
		PrettyPrint: defaultPrettyPrint,
	}
	for _, opt := range options {
		cfg = opt.apply(cfg)
	}
	return cfg, nil
}

// Option sets the value of an option for a Config.
type Option interface {
	apply(config) config
}

// WithWriter sets the export stream destination.
func WithWriter(w io.Writer) Option {
	return writerOption{w}
}

type writerOption struct {
	W io.Writer
}

func (o writerOption) apply(cfg config) config {
	cfg.Writer = o.W
	return cfg
}

// WithPrettyPrint sets the export stream format to use JSON.
func WithPrettyPrint() Option {
	return prettyPrintOption(true)
}

type prettyPrintOption bool

func (o prettyPrintOption) apply(cfg config) config {
	cfg.PrettyPrint = bool(o)
	return cfg
}
