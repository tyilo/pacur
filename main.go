package main

import (
	"github.com/Tyilo/pacur/cmd"
)

func main() {
	err := cmd.Parse()
	if err != nil {
		panic(err)
	}
}
