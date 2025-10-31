package util

import (
	"bytes"
	"net/url"
	"strings"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func NewBuffer() *bytes.Buffer {
	buf := bufferPool.Get().(*bytes.Buffer)
	return buf
}

func ReleaseBuffer(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}

var stringsBuilder = sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

func NewStringsBuilder() *strings.Builder {
	b := stringsBuilder.Get().(*strings.Builder)
	return b
}

func ReleaseStringsBuilder(b *strings.Builder) {
	b.Reset()
	stringsBuilder.Put(b)
}

var urlValuesPool = sync.Pool{
	New: func() any {
		return make(url.Values)
	},
}

func NewURLValues() url.Values {
	vals := urlValuesPool.Get().(url.Values)
	return vals
}

func ReleaseURLValues(vals url.Values) {
	for k := range vals {
		vals.Del(k)
	}
	urlValuesPool.Put(vals)
}
