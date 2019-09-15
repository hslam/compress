go test -v -run="none" -bench=. -benchtime=1s -benchmem
```
goos: darwin
goarch: amd64
pkg: hslam.com/mgit/Mort/compress
BenchmarkCompressBestSpeedFlate-4            	   10000	    193690 ns/op	 1205608 B/op	      20 allocs/op
BenchmarkCompressBestSpeedZlib-4             	   10000	    198822 ns/op	 1206249 B/op	      24 allocs/op
BenchmarkCompressBestSpeedGzip-4             	   10000	    208990 ns/op	 1206306 B/op	      22 allocs/op
BenchmarkCompressBestCompressionFlate-4      	   10000	    154420 ns/op	  814688 B/op	      19 allocs/op
BenchmarkCompressBestCompressionZlib-4       	   10000	    158295 ns/op	  815333 B/op	      23 allocs/op
BenchmarkCompressBestCompressionGzip-4       	   10000	    149546 ns/op	  815392 B/op	      21 allocs/op
BenchmarkCompressDefaultCompressionFlate-4   	   10000	    154663 ns/op	  814688 B/op	      19 allocs/op
BenchmarkCompressDefaultCompressionZlib-4    	   10000	    153532 ns/op	  815332 B/op	      23 allocs/op
BenchmarkCompressDefaultCompressionGzip-4    	   10000	    151166 ns/op	  815392 B/op	      21 allocs/op
BenchmarkAllBestSpeedFlate-4                 	   10000	    216400 ns/op	 1248233 B/op	      28 allocs/op
BenchmarkAllBestSpeedZlib-4                  	   10000	    225127 ns/op	 1248958 B/op	      34 allocs/op
BenchmarkAllBestSpeedGzip-4                  	   10000	    220994 ns/op	 1249634 B/op	      31 allocs/op
BenchmarkAllBestCompressionFlate-4           	   10000	    171303 ns/op	  857313 B/op	      27 allocs/op
BenchmarkAllBestCompressionZlib-4            	   10000	    175180 ns/op	  858042 B/op	      33 allocs/op
BenchmarkAllBestCompressionGzip-4            	   10000	    174148 ns/op	  858720 B/op	      30 allocs/op
BenchmarkAllDefaultCompressionFlate-4        	   10000	    174826 ns/op	  857313 B/op	      27 allocs/op
BenchmarkAllDefaultCompressionZlib-4         	   10000	    176618 ns/op	  858042 B/op	      33 allocs/op
BenchmarkAllDefaultCompressionGzip-4         	   10000	    173646 ns/op	  858720 B/op	      30 allocs/op
PASS
ok  	hslam.com/mgit/Mort/compress	33.659s
```
example
```
DhPNExsTOYfCZrzQ2wPcHDHqtSJAS7gMLeqejz6ZvLf9QIZxoYGy8ooaJfS12mbUTeg0bCiwpdukKbYPP1LgxNulu7qQJMtr1WQ0boTtnJiRXsN6r0W7q7f3dh5lRZVVhwlVOupxD4D7i0YWKvsJGkSIATQtZGsPcOBKspKN6vu3ugxaO8Jv2jYtm3MB2RwokaXlRjr37whZyr8Tam3rM9OGinGfSKCwUeTzLhg2dz83slIsCdO3lqzkx9iBWyaBtDebfYN465CAndAJN1JqFzhRfaXcw9KNPBUHn3CGioQTTrWzrIWiHTi70zPDSiFRjmLQdJKKgCpVgtk20DHFcFme7dUlV4l5RFWHUoNea9FUN1cFvnyZVYpCnZz3uPgpa1dE3DkSodQMAYJZOpeB29EnyYj4UKvDehxqVAwgDyMS82PXbbi4H6UPiwOX9IomXGvL8jUt441ENrfXObR6kEjSstjnQ6pEX9VchH66cMlFE4qgRUkvKFIycfw5F208
flat 512->410->512
zlib 512->412->512
gzip 512->420->512
```