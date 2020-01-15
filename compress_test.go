// Copyright (c) 2019 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package compress

import (
	"bytes"
	"testing"
)

func TestCompressor(t *testing.T) {
	readyBytes := []byte("DhPNExsTOYfCZrzQ2wPcHDHqtSJAS7gMLeqejz6ZvLf9QIZxoYGy8ooaJfS12mbUTeg0bCiwpdukKbYPP1LgxNulu7qQJMtr1WQ0boTtnJiRXsN6r0W7q7f3dh5lRZVVhwlVOupxD4D7i0YWKvsJGkSIATQtZGsPcOBKspKN6vu3ugxaO8Jv2jYtm3MB2RwokaXlRjr37whZyr8Tam3rM9OGinGfSKCwUeTzLhg2dz83slIsCdO3lqzkx9iBWyaBtDebfYN465CAndAJN1JqFzhRfaXcw9KNPBUHn3CGioQTTrWzrIWiHTi70zPDSiFRjmLQdJKKgCpVgtk20DHFcFme7dUlV4l5RFWHUoNea9FUN1cFvnyZVYpCnZz3uPgpa1dE3DkSodQMAYJZOpeB29EnyYj4UKvDehxqVAwgDyMS82PXbbi4H6UPiwOX9IomXGvL8jUt441ENrfXObR6kEjSstjnQ6pEX9VchH66cMlFE4qgRUkvKFIycfw5F208")
	Levels := []Level{BestCompression, BestSpeed, DefaultCompression, Level2, Level3, Level4, Level5, Level6, Level7, Level8, NoCompression, HuffmanOnly}
	CompressTypes := []Type{Flate, Zlib, Gzip}
	for _, level := range Levels {
		for _, compressType := range CompressTypes {
			compressor := &Compressor{Type: compressType, Level: level}
			compressed := compressor.Compress(readyBytes)
			out := compressor.Uncompress(compressed)
			if !bytes.Equal(readyBytes, out) {
				t.Errorf("error %v !Equal %v", readyBytes, out)
			}
		}
	}
	{
		compressor := &Compressor{Type: Gzip, Level: DefaultCompression}
		compressed := compressor.Compress(readyBytes)
		out := compressor.Uncompress(compressed)
		if !bytes.Equal(readyBytes, out) {
			t.Errorf("error %v !Equal %v", readyBytes, out)
		}
	}
	{
		compressor := &Compressor{Level: DefaultCompression}
		compressed := compressor.Compress(readyBytes)
		out := compressor.Uncompress(compressed)
		if !bytes.Equal(readyBytes, out) {
			t.Errorf("error %v !Equal %v", readyBytes, out)
		}
	}
}

func BenchmarkFlate(b *testing.B) {
	readyBytes := []byte("DhPNExsTOYfCZrzQ2wPcHDHqtSJAS7gMLeqejz6ZvLf9QIZxoYGy8ooaJfS12mbUTeg0bCiwpdukKbYPP1LgxNulu7qQJMtr1WQ0boTtnJiRXsN6r0W7q7f3dh5lRZVVhwlVOupxD4D7i0YWKvsJGkSIATQtZGsPcOBKspKN6vu3ugxaO8Jv2jYtm3MB2RwokaXlRjr37whZyr8Tam3rM9OGinGfSKCwUeTzLhg2dz83slIsCdO3lqzkx9iBWyaBtDebfYN465CAndAJN1JqFzhRfaXcw9KNPBUHn3CGioQTTrWzrIWiHTi70zPDSiFRjmLQdJKKgCpVgtk20DHFcFme7dUlV4l5RFWHUoNea9FUN1cFvnyZVYpCnZz3uPgpa1dE3DkSodQMAYJZOpeB29EnyYj4UKvDehxqVAwgDyMS82PXbbi4H6UPiwOX9IomXGvL8jUt441ENrfXObR6kEjSstjnQ6pEX9VchH66cMlFE4qgRUkvKFIycfw5F208")
	compressor := &Compressor{Type: Flate, Level: DefaultCompression}
	for i := 0; i < b.N; i++ {
		compressed := compressor.Compress(readyBytes)
		compressor.Uncompress(compressed)
	}
}

func BenchmarkZlib(b *testing.B) {
	readyBytes := []byte("DhPNExsTOYfCZrzQ2wPcHDHqtSJAS7gMLeqejz6ZvLf9QIZxoYGy8ooaJfS12mbUTeg0bCiwpdukKbYPP1LgxNulu7qQJMtr1WQ0boTtnJiRXsN6r0W7q7f3dh5lRZVVhwlVOupxD4D7i0YWKvsJGkSIATQtZGsPcOBKspKN6vu3ugxaO8Jv2jYtm3MB2RwokaXlRjr37whZyr8Tam3rM9OGinGfSKCwUeTzLhg2dz83slIsCdO3lqzkx9iBWyaBtDebfYN465CAndAJN1JqFzhRfaXcw9KNPBUHn3CGioQTTrWzrIWiHTi70zPDSiFRjmLQdJKKgCpVgtk20DHFcFme7dUlV4l5RFWHUoNea9FUN1cFvnyZVYpCnZz3uPgpa1dE3DkSodQMAYJZOpeB29EnyYj4UKvDehxqVAwgDyMS82PXbbi4H6UPiwOX9IomXGvL8jUt441ENrfXObR6kEjSstjnQ6pEX9VchH66cMlFE4qgRUkvKFIycfw5F208")
	compressor := &Compressor{Type: Zlib, Level: DefaultCompression}
	for i := 0; i < b.N; i++ {
		compressed := compressor.Compress(readyBytes)
		compressor.Uncompress(compressed)
	}
}

func BenchmarkGzip(b *testing.B) {
	readyBytes := []byte("DhPNExsTOYfCZrzQ2wPcHDHqtSJAS7gMLeqejz6ZvLf9QIZxoYGy8ooaJfS12mbUTeg0bCiwpdukKbYPP1LgxNulu7qQJMtr1WQ0boTtnJiRXsN6r0W7q7f3dh5lRZVVhwlVOupxD4D7i0YWKvsJGkSIATQtZGsPcOBKspKN6vu3ugxaO8Jv2jYtm3MB2RwokaXlRjr37whZyr8Tam3rM9OGinGfSKCwUeTzLhg2dz83slIsCdO3lqzkx9iBWyaBtDebfYN465CAndAJN1JqFzhRfaXcw9KNPBUHn3CGioQTTrWzrIWiHTi70zPDSiFRjmLQdJKKgCpVgtk20DHFcFme7dUlV4l5RFWHUoNea9FUN1cFvnyZVYpCnZz3uPgpa1dE3DkSodQMAYJZOpeB29EnyYj4UKvDehxqVAwgDyMS82PXbbi4H6UPiwOX9IomXGvL8jUt441ENrfXObR6kEjSstjnQ6pEX9VchH66cMlFE4qgRUkvKFIycfw5F208")
	compressor := &Compressor{Type: Gzip, Level: DefaultCompression}
	for i := 0; i < b.N; i++ {
		compressed := compressor.Compress(readyBytes)
		compressor.Uncompress(compressed)
	}
}
