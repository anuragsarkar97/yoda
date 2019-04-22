package main

import (
	"fmt"
	"yoda/pkg/Streams"
)

func main()  {
	fmt.Println("yolo")
	a := Streams.GenerateFileStreams("/Users/oyo/go/src/yoda/pkg/main.go", 100000000)
	fmt.Println(a)
}