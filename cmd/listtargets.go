package cmd

import (
	"fmt"
	"github.com/Tyilo/pacur/constants"
)

func ListTargets() (_ error) {
	for _, release := range constants.Releases {
		fmt.Println(release)
	}

	return
}
