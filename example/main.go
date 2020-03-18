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
