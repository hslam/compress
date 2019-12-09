# compress
A compress library written in golang.

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
```
package main

import (
	"github.com/hslam/compress"
	"fmt"
)

func main() {
	readyBytes:=[]byte("DhPNExsTOYfCZrzQ2wPcHDHqtSJAS7gMLeqejz6ZvLf9QIZxoYGy8ooaJfS12mbUTeg0bCiwpdukKbYPP1LgxNulu7qQJMtr1WQ0boTtnJiRXsN6r0W7q7f3dh5lRZVVhwlVOupxD4D7i0YWKvsJGkSIATQtZGsPcOBKspKN6vu3ugxaO8Jv2jYtm3MB2RwokaXlRjr37whZyr8Tam3rM9OGinGfSKCwUeTzLhg2dz83slIsCdO3lqzkx9iBWyaBtDebfYN465CAndAJN1JqFzhRfaXcw9KNPBUHn3CGioQTTrWzrIWiHTi70zPDSiFRjmLQdJKKgCpVgtk20DHFcFme7dUlV4l5RFWHUoNea9FUN1cFvnyZVYpCnZz3uPgpa1dE3DkSodQMAYJZOpeB29EnyYj4UKvDehxqVAwgDyMS82PXbbi4H6UPiwOX9IomXGvL8jUt441ENrfXObR6kEjSstjnQ6pEX9VchH66cMlFE4qgRUkvKFIycfw5F208")
	var (
		compressor compress.Compressor
		compressBytes []byte
		uncompressBytes []byte
	)
	//FlateCompressor
	compressor=compress.FlateCompressor{Level:compress.BestCompression}
	compressBytes,_=compressor.Compress(readyBytes)
	uncompressBytes,_=compressor.Uncompress(compressBytes)
	fmt.Printf("flat %d->%d->%d\n",len(readyBytes),len(compressBytes),len(uncompressBytes))

	//ZlibCompressor
	compressor=compress.ZlibCompressor{Level:compress.BestCompression}
	compressBytes,_=compressor.Compress(readyBytes)
	uncompressBytes,_=compressor.Uncompress(compressBytes)
	fmt.Printf("zlib %d->%d->%d\n",len(readyBytes),len(compressBytes),len(uncompressBytes))
	//GzipCompressor
	compressor=compress.GzipCompressor{Level:compress.BestCompression}
	compressBytes,_=compressor.Compress(readyBytes)
	uncompressBytes,_=compressor.Uncompress(compressBytes)
	fmt.Printf("gzip %d->%d->%d\n",len(readyBytes),len(compressBytes),len(uncompressBytes))
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
BenchmarkCompressBestSpeedFlate-4            	   10000	    190878 ns/op	 1205606 B/op	      20 allocs/op
BenchmarkCompressBestSpeedZlib-4             	   10000	    195048 ns/op	 1206249 B/op	      24 allocs/op
BenchmarkCompressBestSpeedGzip-4             	   10000	    194059 ns/op	 1206306 B/op	      22 allocs/op
BenchmarkCompressBestCompressionFlate-4      	   10000	    154021 ns/op	  814688 B/op	      19 allocs/op
BenchmarkCompressBestCompressionZlib-4       	   10000	    156747 ns/op	  815332 B/op	      23 allocs/op
BenchmarkCompressBestCompressionGzip-4       	   10000	    150377 ns/op	  815392 B/op	      21 allocs/op
BenchmarkCompressDefaultCompressionFlate-4   	   10000	    154451 ns/op	  814688 B/op	      19 allocs/op
BenchmarkCompressDefaultCompressionZlib-4    	   10000	    158017 ns/op	  815333 B/op	      23 allocs/op
BenchmarkCompressDefaultCompressionGzip-4    	   10000	    149367 ns/op	  815392 B/op	      21 allocs/op
BenchmarkAllBestSpeedFlate-4                 	   10000	    211545 ns/op	 1248232 B/op	      28 allocs/op
BenchmarkAllBestSpeedZlib-4                  	   10000	    216422 ns/op	 1248958 B/op	      34 allocs/op
BenchmarkAllBestSpeedGzip-4                  	   10000	    213844 ns/op	 1249635 B/op	      31 allocs/op
BenchmarkAllBestCompressionFlate-4           	   10000	    171548 ns/op	  857314 B/op	      27 allocs/op
BenchmarkAllBestCompressionZlib-4            	   10000	    174537 ns/op	  858043 B/op	      33 allocs/op
BenchmarkAllBestCompressionGzip-4            	   10000	    172721 ns/op	  858720 B/op	      30 allocs/op
BenchmarkAllDefaultCompressionFlate-4        	   10000	    173061 ns/op	  857313 B/op	      27 allocs/op
BenchmarkAllDefaultCompressionZlib-4         	   10000	    184280 ns/op	  858042 B/op	      33 allocs/op
BenchmarkAllDefaultCompressionGzip-4         	   10000	    173342 ns/op	  858720 B/op	      30 allocs/op
PASS
ok  	github.com/hslam/compress	33.265s
```

### Licence
This package is licenced under a MIT licence (Copyright (c) 2019 Meng Huang)

### Authors
compress was written by Meng Huang.
