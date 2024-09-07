package main

import (
	"integriWatch/system"
)

func main() {
	if err := system.Init(); err != nil {
		panic(err)
	}
}
