# compress
[![GoDoc](https://godoc.org/github.com/hslam/compress?status.svg)](https://godoc.org/github.com/hslam/compress)
[![Build Status](https://github.com/hslam/compress/workflows/build/badge.svg)](https://github.com/hslam/compress/actions)
[![codecov](https://codecov.io/gh/hslam/compress/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/compress)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/compress)](https://goreportcard.com/report/github.com/hslam/compress)
[![GitHub release](https://img.shields.io/github/release/hslam/compress.svg)](https://github.com/hslam/compress/releases/latest)
[![LICENSE](https://img.shields.io/github/license/hslam/compress.svg?style=flat-square)](https://github.com/hslam/compress/blob/master/LICENSE)

Package compress implements compressing and uncompressing in golang.

## Feature
* flate
* zlib
* gzip

## Get started

### Install
```
go get github.com/hslam/compress
```
### Import
```
import "github.com/hslam/compress"
```
### Usage
#### Example
```go
package main

import (
	"fmt"
	"github.com/hslam/compress"
)

var readyBytes = []byte("DhPNExsTOYfCZrzQ2wPcHDHqtSJAS7gMLeqejz6ZvLf9QIZxoYGy8ooaJfS12mbUTeg0bCiwpdukKbYPP1LgxNulu7qQJMtr1WQ0boTtnJiRXsN6r0W7q7f3dh5lRZVVhwlVOupxD4D7i0YWKvsJGkSIATQtZGsPcOBKspKN6vu3ugxaO8Jv2jYtm3MB2RwokaXlRjr37whZyr8Tam3rM9OGinGfSKCwUeTzLhg2dz83slIsCdO3lqzkx9iBWyaBtDebfYN465CAndAJN1JqFzhRfaXcw9KNPBUHn3CGioQTTrWzrIWiHTi70zPDSiFRjmLQdJKKgCpVgtk20DHFcFme7dUlV4l5RFWHUoNea9FUN1cFvnyZVYpCnZz3uPgpa1dE3DkSodQMAYJZOpeB29EnyYj4UKvDehxqVAwgDyMS82PXbbi4H6UPiwOX9IomXGvL8jUt441ENrfXObR6kEjSstjnQ6pEX9VchH66cMlFE4qgRUkvKFIycfw5F208")

func main() {
	Flate()
	Zlib()
	Gzip()
}

//Flate Example
func Flate() {
	compressor := &compress.Compressor{Type: compress.Flate, Level: compress.DefaultCompression}
	compressBytes := compressor.Compress(readyBytes)
	uncompressBytes := compressor.Uncompress(compressBytes)
	fmt.Printf("flat %d->%d->%d\n", len(readyBytes), len(compressBytes), len(uncompressBytes))
}

//Zlib Example
func Zlib() {
	compressor := &compress.Compressor{Type: compress.Zlib, Level: compress.DefaultCompression}
	compressBytes := compressor.Compress(readyBytes)
	uncompressBytes := compressor.Uncompress(compressBytes)
	fmt.Printf("zlib %d->%d->%d\n", len(readyBytes), len(compressBytes), len(uncompressBytes))
}

//Gzip Example
func Gzip() {
	compressor := &compress.Compressor{Type: compress.Gzip, Level: compress.DefaultCompression}
	compressBytes := compressor.Compress(readyBytes)
	uncompressBytes := compressor.Uncompress(compressBytes)
	fmt.Printf("gzip %d->%d->%d\n", len(readyBytes), len(compressBytes), len(uncompressBytes))
}
```

### Output
```
flat 512->410->512
zlib 512->412->512
gzip 512->420->512
```

### Benchmark
go test -v -run="none" -bench=. -benchtime=1s -benchmem
```
goos: darwin
goarch: amd64
pkg: github.com/hslam/compress
BenchmarkFlate-4   	    6006	    194793 ns/op	  857312 B/op	      27 allocs/op
BenchmarkZlib-4    	    6276	    199555 ns/op	  858042 B/op	      33 allocs/op
BenchmarkGzip-4    	    6139	    197678 ns/op	  858720 B/op	      30 allocs/op
PASS
ok  	github.com/hslam/compress	4.902s
```

### License
This package is licensed under a MIT license (Copyright (c) 2019 Meng Huang)

### Authors
compress was written by Meng Huang.
