package strategy

type CompressionStrategy interface {
	Compress(data []byte) []byte
}

type ZipStrategy struct{}

func (z ZipStrategy) Compress(data []byte) []byte { return []byte("zip") }

type GzipStrategy struct{}

func (g GzipStrategy) Compress(data []byte) []byte { return []byte("gzip") }
