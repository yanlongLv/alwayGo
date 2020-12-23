package logs

import "fmt"

var nop Logger = new(nopLogger)

//Option ...
type Option func(*options)

type options struct {
	level   Level
	verbose Verbose
}

//AllowLevel ..
func AllowLevel(l Level) Option {
	return func(o *options) {
		o.level = l
	}
}

// AllowVerbose ..
func AllowVerbose(v Verbose) Option {
	return func(o *options) {
		o.verbose = v
	}
}

// Helper ..
type Helper struct {
	Logger
	opts options
}

//NewHelper ..
func NewHelper(name string, l Logger, opts ...Option) *Helper {
	options := options{}
	for _, o := range opts {
		o(&options)
	}
	return &Helper{Logger: With(l, "module", name), opts: options}
}

//V ...
func (h *Helper) V(v Verbose) Logger {
	if h.opts.verbose.Enabled(v) {
		return nop
	}
	return With(h, VerboseKey, v)
}

// Debug ...
func (h *Helper) Debug(a ...interface{}) {
	if h.opts.level.Enabled(LevelDebug) {
		Debug(h).Print("log", fmt.Sprint(a...))
	}
}

// Debugf ...
func (h *Helper) Debugf(format string, a ...interface{}) {
	if h.opts.level.Enabled(LevelDebug) {
		Debug(h).Print("log", fmt.Sprintf(format, a...))
	}
}

// Debugw ...
func (h *Helper) Debugw(kvpair ...interface{}) {
	if h.opts.level.Enabled(LevelDebug) {
		Debug(h).Print(kvpair)
	}
}

// Info ...
func (h *Helper) Info(a ...interface{}) {
	if h.opts.level.Enabled(LevelInfo) {
		Debug(h).Print("log", fmt.Sprint(a...))
	}
}

// Infof ...
func (h *Helper) Infof(format string, a ...interface{}) {
	if h.opts.level.Enabled(LevelInfo) {
		Debug(h).Print("log", fmt.Sprintf(format, a...))
	}
}

// Infow ...
func (h *Helper) Infow(kvpair ...interface{}) {
	if h.opts.level.Enabled(LevelInfo) {
		Debug(h).Print(kvpair...)
	}
}

// Warn ...
func (h *Helper) Warn(a ...interface{}) {
	if h.opts.level.Enabled(LevelWarn) {
		Debug(h).Print("log", fmt.Sprint(a...))
	}
}

// Warnf ...
func (h *Helper) Warnf(format string, a ...interface{}) {
	if h.opts.level.Enabled(LevelWarn) {
		Debug(h).Print("log", fmt.Sprintf(format, a...))
	}
}

// Warnw ...
func (h *Helper) Warnw(kvpair ...interface{}) {
	if h.opts.level.Enabled(LevelError) {
		Debug(h).Print(kvpair...)
	}
}

// Error ...
func (h *Helper) Error(a ...interface{}) {
	if h.opts.level.Enabled(LevelWarn) {
		Debug(h).Print("log", fmt.Sprint(a...))
	}
}

// Errorf ...
func (h *Helper) Errorf(format string, a ...interface{}) {
	if h.opts.level.Enabled(LevelWarn) {
		Debug(h).Print("log", fmt.Sprintf(format, a...))
	}
}

// Errorw ...
func (h *Helper) Errorw(kvpair ...interface{}) {
	if h.opts.level.Enabled(LevelError) {
		Debug(h).Print(kvpair...)
	}
}
