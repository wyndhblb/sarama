package sarama

import (
	"bytes"
	"github.com/klauspost/compress/zstd"
	"sync"
)

var (
	zstdWriterPool = sync.Pool{
		New: func() interface{} {
			p, _ := zstd.NewWriter(nil)
			return p
		},
	}
	zstdBytePool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(nil)
		},
	}
	zstdReaderPool = sync.Pool{
		New: func() interface{} {
			p, _ := zstd.NewReader(nil)
			return p
		},
	}
)

func zstdDecompress(dst, src []byte) ([]byte, error) {
	reader := zstdReaderPool.Get().(*zstd.Decoder)
	if dst == nil{
		bs := zstdBytePool.Get().(*bytes.Buffer)
		bs.Reset()
		dst = bs.Bytes()
		zstdBytePool.Put(bs)
	}
	defer zstdReaderPool.Put(reader)
	return reader.DecodeAll(src, dst)
}

func zstdCompress(dst, src []byte) ([]byte, error) {
	writer := zstdWriterPool.Get().(*zstd.Encoder)
	defer zstdWriterPool.Put(writer)
	bs := writer.EncodeAll(src, dst)
	return bs, nil
}
