package stdlog

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
)

var _ log.Logger = (*stdLogger)(nil)

type options struct {
	prefix string
	flag   int
	skip   int
	out    io.Writer
}

//Option is std logger option
type Option func(*options)

//Prefix with logger prefix
func Prefix(prefix string) Option {
	return func(o *options) {
		o.prefix = prefix
	}
}

//Skip with logger writer
func Skip(skip int) Option {
	return func(o *options) {
		o.skip = skip
	}
}

//Writer with logger writer
func Writer(out io.Writer) Option {
	return func(o *options) {
		o.out = out
	}
}

type stdLogger struct {
	opts options
	log  *log.Logger
	pool *sync.Pool
}

//NewLogger ..
func NewLogger(opts ...Option) *stdLogger {
	options := options{
		flag: log.LstdFlags,
		skip: 2,
		out:  os.Stdout,
	}
	for _, o := range opts {

		o(&options)
	}
	return &stdLogger{
		opts: options,
		log:  log.New(options.out, options.prefix, options.flag),
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}
