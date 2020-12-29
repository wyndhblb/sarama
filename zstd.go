package sarama

import (
	"github.com/DataDog/zstd"
)


func zstdDecompress(dst, src []byte) ([]byte, error) {
	return zstd.Decompress(dst, src)
}

func zstdCompress(dst, src []byte) ([]byte, error) {
	return zstd.Compress(dst, src)
}
