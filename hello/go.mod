module Go/hello

go 1.25.1

require (
	github.com/yuvbro/go-greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.29.0 // indirect
	rsc.io/sampler v1.99.99 // indirect
)

replace github.com/yuvbro/go-greetings => ../greetings
