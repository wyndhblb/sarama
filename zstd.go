package sarama

import (
	"bytes"
	"github.com/klauspost/compress/zstd"
	"sync"
)

var (
	zstdWriter, _ = zstd.NewWriter(nil)
	zstdWriterPool = sync.Pool{
		New: func() interface{} {
			p, _ := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedFastest))
			return p
		},
	}
	zstdBytePool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(nil)
		},
	}
	zstdReader, _ = zstd.NewReader(nil)
	zstdReaderPool = sync.Pool{
		New: func() interface{} {
			p, _ := zstd.NewReader(nil, zstd.WithDecoderLowmem(false), zstd.WithDecoderConcurrency(2))
			return p
		},
	}
)

func zstdDecompress(dst, src []byte) ([]byte, error) {

	return zstdReader.DecodeAll(src, dst)

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
	bs := zstdWriter.EncodeAll(src, dst)
	return bs, nil

	writer := zstdWriterPool.Get().(*zstd.Encoder)
	defer zstdWriterPool.Put(writer)
	bs := writer.EncodeAll(src, dst)
	return bs, nil
}
