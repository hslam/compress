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