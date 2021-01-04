package sarama

import (
	"github.com/klauspost/compress/zstd"
)

var (
	zstdDec, _ = zstd.NewReader(nil)
	zstdEnc, _ = zstd.NewWriter(nil, zstd.WithZeroFrames(true))
)

func zstdDecompress(dst, src []byte) ([]byte, error) {
	d, _ :=  zstd.NewReader(nil)
	return d.DecodeAll(src, dst)
}

func zstdCompress(dst, src []byte) ([]byte, error) {
	return zstdEnc.EncodeAll(src, dst), nil
}
