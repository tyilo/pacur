package packer

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/Tyilo/pacur/constants"
	"github.com/Tyilo/pacur/debian"
	"github.com/Tyilo/pacur/pack"
	"github.com/Tyilo/pacur/pacman"
	"github.com/Tyilo/pacur/redhat"
)

type Packer interface {
	Prep() error
	Build() error
}

func GetPacker(pac *pack.Pack, distro, release string) (
	pcker Packer, err error) {

	switch constants.DistroPack[distro] {
	case "pacman":
		pcker = &pacman.Pacman{
			Pack: pac,
		}
	case "debian":
		pcker = &debian.Debian{
			Pack: pac,
		}
	case "redhat":
		pcker = &redhat.Redhat{
			Pack: pac,
		}
	default:
		system := distro
		if release != "" {
			system += "-" + release
		}

		err = &UnknownSystem{
			errors.Newf("packer: Unknown system %s", system),
		}
		return
	}

	return
}
