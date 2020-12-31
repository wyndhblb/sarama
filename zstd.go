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
	zstdReaderPool = sync.Pool{
		New: func() interface{} {
			p, _ := zstd.NewReader(nil)
			return p
		},
	}
)

func zstdDecompress(dst, src []byte) ([]byte, error) {
	reader := zstdReaderPool.Get().(*zstd.Decoder)
	defer zstdReaderPool.Put(reader)
	return reader.DecodeAll(src, dst)
}

func zstdCompress(dst, src []byte) ([]byte, error) {
	writer := zstdWriterPool.Get().(*zstd.Encoder)
	defer zstdWriterPool.Put(writer)
	bs := writer.EncodeAll(src, dst)
	return bs, nil
}
