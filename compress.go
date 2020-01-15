// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package compress implements compressing and uncompressing.
package compress

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"io"
)

// Level defines the level for compression.
// Higher levels typically run slower but compress more.
type Level int

const (
	//NoCompression does not attempt any compression.
	NoCompression Level = 0
	//BestSpeed defines the level of best speed.
	BestSpeed Level = 1
	//Level2 defines the level 2.
	Level2 Level = 2
	//Level3 defines the level 3.
	Level3 Level = 3
	//Level4 defines the level 4.
	Level4 Level = 4
	//Level5 defines the level 5.
	Level5 Level = 5
	//Level6 defines the level 6.
	Level6 Level = 6
	//Level7 defines the level 7.
	Level7 Level = 7
	//Level8 defines the level 8.
	Level8 Level = 8
	//BestCompression defines the level of best compression.
	BestCompression Level = 9
	//DefaultCompression uses the default compression level.
	DefaultCompression Level = -1
	// HuffmanOnly will use Huffman compression only, giving
	// a very fast compression for all types of input,
	// but sacrificing considerable compression efficiency.
	HuffmanOnly Level = -2
)

// Type defines the type for compression.
type Type string

const (
	//Flate defines the compressor type of flate .
	Flate Type = "flate"
	//Zlib defines the compressor type of zlib .
	Zlib Type = "zlib"
	//Gzip defines the compressor type of gzip .
	Gzip Type = "gzip"
)

// Writer defines the writer interface for compressing data.
type Writer interface {
	Write(data []byte) (n int, err error)
	Flush() error
	Close() error
}

// Reader defines the reader interface for uncompressing data.
type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

// Compressor defines the struct of compressor.
type Compressor struct {
	Type  Type
	Level Level
}

// Compress returns the compressed data.
func (c *Compressor) Compress(data []byte) []byte {
	buf := bytes.NewBuffer(nil)
	var writer Writer
	switch c.Type {
	case Flate:
		writer, _ = flate.NewWriter(buf, int(c.Level))
	case Zlib:
		writer, _ = zlib.NewWriterLevel(buf, int(c.Level))
	case Gzip:
		writer, _ = gzip.NewWriterLevel(buf, int(c.Level))
	default:
		return data
	}
	defer writer.Close()
	writer.Write(data)
	writer.Flush()
	return buf.Bytes()
}

// Uncompress returns the uncompressed data.
func (c *Compressor) Uncompress(data []byte) []byte {
	var reader Reader
	switch c.Type {
	case Flate:
		reader = flate.NewReader(bytes.NewBuffer(data))
	case Zlib:
		reader, _ = zlib.NewReader(bytes.NewBuffer(data))
	case Gzip:
		reader, _ = gzip.NewReader(bytes.NewBuffer(data))
	default:
		return data
	}
	defer reader.Close()
	var out bytes.Buffer
	io.Copy(&out, reader)
	return out.Bytes()
}
