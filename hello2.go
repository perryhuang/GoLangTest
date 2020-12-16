package hello2

import (
	"rsc.io/quote/v3"
)

func Hello() string {
	//	return quote.HelloV3()
	return "Perry Test V1"
	//	return "Hello, world."
	//	return "Hello, world."
}

func Proverb() string {
	return quote.Concurrency()
}
