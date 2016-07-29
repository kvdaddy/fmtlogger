package fmtlogger

import (
	"io"
	"github.com/go-logfmt/logfmt"
	"bytes"
	"sync"
)

type FmtLogger interface {
	Log(keyvals ...interface{}) error
}

type fmtEncoder struct {
	*logfmt.Encoder
	buf bytes.Buffer
}

func (e *fmtEncoder) Reset() {
	e.Encoder.Reset()
	e.buf.Reset()
}

var fmtEncoderPool = sync.Pool{
	New:func() interface{}{
		var enc fmtEncoder
		enc.Encoder = logfmt.NewEncoder(&enc.buf)
		return &enc
	},
}

func NewFmtLogger(w io.Writer) FmtLogger {
	return &fmtLogger{w: w}
}

type fmtLogger struct {
	w io.Writer
}

func (l fmtLogger) Log(keyvals ...interface{}) error {
	enc := fmtEncoderPool.Get().(*fmtEncoder)
	enc.Reset()
	defer fmtEncoderPool.Put(enc)

	if err := enc.EncodeKeyvals(keyvals...); err != nil {
		return err
	}

	if err := enc.EndRecord(); err != nil {
		return err
	}

	if _, err := l.w.Write(enc.buf.Bytes()); err != nil {
		return err
	}
	return nil
}











