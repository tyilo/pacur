package main

import (
	"fmt"
	"github.com/pacur/pacur/parse"
	"github.com/pacur/pacur/builder"
	"github.com/pacur/pacur/debian"
)

func main() {
	pac, err := parse.File("/pkg/PKGBUILD")
	if err != nil {
		panic(err)
	}

	err = pac.Compile()
	if err != nil {
		panic(err)
	}

	builder := builder.Builder{
		Pack: pac,
	}
	err = builder.Build()
	if err != nil {
		panic(err)
	}

	deb := debian.Debian{
		Pack: pac,
	}
	err = deb.Build()
	if err != nil {
		panic(err)
	}
}